/*
 * Copyright 2024 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tr064

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"
)

func NewClient(url *url.URL) *Client {
	return &Client{
		Url:   url,
		mutex: &sync.Mutex{},
	}
}

type Client struct {
	Url                   *url.URL
	Timeout               time.Duration
	InsecureSkipVerify    bool
	Debug                 bool
	mutex                 *sync.Mutex
	cachedServices        []ServiceDescriptor
	cachedHttpClient      *http.Client
	cachedAuthentications map[string]string
}

type ServiceDescriptor interface {
	Name() string
	Type() string
	Id() string
	Url() string
}

type StaticServiceDescriptor struct {
	ServiceName string
	ServiceType string
	ServiceId   string
	ServiceUrl  string
}

func (service *StaticServiceDescriptor) Name() string {
	return service.ServiceName
}

func (service *StaticServiceDescriptor) Type() string {
	return service.ServiceType
}

func (service *StaticServiceDescriptor) Id() string {
	return service.ServiceId
}

func (service *StaticServiceDescriptor) Url() string {
	return service.ServiceUrl
}

func (client *Client) Services() ([]ServiceDescriptor, error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.cachedServices == nil {
		tr64desc, err := client.fetchSpec()
		if err != nil {
			return nil, err
		}
		collector := &serviceCollector{serviceMap: make(map[string]ServiceDescriptor)}
		err = tr64desc.walk(client.Url, collector.collectService)
		if err != nil {
			return nil, err
		}
		services := slices.Collect(maps.Values(collector.serviceMap))
		slices.SortFunc(services, func(a ServiceDescriptor, b ServiceDescriptor) int { return strings.Compare(a.Type(), b.Type()) })
		client.cachedServices = services
	}
	return client.cachedServices, nil
}

func (client *Client) ServicesByName(name string) ([]ServiceDescriptor, error) {
	all, err := client.Services()
	if err != nil {
		return nil, err
	}
	services := make([]ServiceDescriptor, 0)
	for _, service := range all {
		if service.Name() == name {
			services = append(services, service)
		}
	}
	return services, nil
}

func (client *Client) fetchSpec() (*tr64descDoc, error) {
	tr64descUrl := client.Url.JoinPath(rootSpec)
	tr64desc := &tr64descDoc{}
	err := unmarshalDocument(tr64descUrl, tr64desc)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch '%s' (cause: %w)", tr64descUrl, err)
	}
	return tr64desc, nil
}

type serviceCollector struct {
	serviceMap map[string]ServiceDescriptor
}

func (collector *serviceCollector) collectService(service *serviceDoc, scpd *scpdDoc) error {
	collector.serviceMap[service.ServiceId] = service
	return nil
}

type SOAPRequest[T any] struct {
	XMLName          xml.Name `xml:"s:Envelope"`
	XMLNameSpace     string   `xml:"xmlns:s,attr"`
	XMLEncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body             *SOAPRequestBody[T]
}

type SOAPRequestBody[T any] struct {
	XMLName xml.Name `xml:"s:Body"`
	In      *T
}

type SOAPResponse[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    *SOAPResponseBody[T]
}

type SOAPResponseBody[T any] struct {
	XMLName xml.Name `xml:"Body"`
	Out     *T
}

func (client *Client) InvokeService(service ServiceDescriptor, actionName, in any, out any) error {
	controlUrl, err := url.Parse(service.Url())
	if err != nil {
		return fmt.Errorf("failed to parse control URL '%s' (cause: %w)", service.Url(), err)
	}
	endpoint := client.Url.ResolveReference(controlUrl).String()
	soapAction := fmt.Sprintf("%s#%s", service.Type(), actionName)
	requestBody, err := xml.MarshalIndent(in, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal request (cause: %w)", err)
	}
	if client.Debug {
		log.Println("Request:\n", string(requestBody))
	}
	authentication := client.authentication(service.Type())
	response, err := client.postSoapActionRequest(endpoint, soapAction, requestBody, authentication)
	if err != nil {
		return err
	}
	if response.StatusCode == http.StatusUnauthorized {
		authentication, err := client.authenticate(response, service.Type())
		if err == nil {
			response, err = client.postSoapActionRequest(endpoint, soapAction, requestBody, authentication)
			if err != nil {
				return err
			}
		}
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("service call failure (status: %s)", response.Status)
	}
	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read response body (cause: %w)", err)
	}
	if client.Debug {
		log.Println("Response:\n", string(responseBody))
	}
	err = xml.Unmarshal(responseBody, out)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body (cause: %w)", err)
	}
	return nil
}

func (client *Client) authentication(serviceType string) string {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.cachedAuthentications == nil {
		return ""
	}
	return client.cachedAuthentications[serviceType]
}

func (client *Client) authenticate(challenge *http.Response, serviceType string) (string, error) {
	challengeHeader := challenge.Header["Www-Authenticate"]
	if len(challengeHeader) != 1 {
		return "", fmt.Errorf("missing or unexpected WWW-Authenticate header")
	}
	challengeValues := make(map[string]string)
	for _, challengeHeaderValue := range strings.Split(challengeHeader[0], ",") {
		splitChallengeHeaderValue := strings.Split(challengeHeaderValue, "=")
		if len(splitChallengeHeaderValue) == 2 {
			key := splitChallengeHeaderValue[0]
			value := splitChallengeHeaderValue[1]
			if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
				value = value[1 : len(value)-1]
			}
			challengeValues[key] = value
		}
	}
	digestRealm := challengeValues["Digest realm"]
	username := client.Url.User.Username()
	password, _ := client.Url.User.Password()
	ha1 := client.md5Hash(fmt.Sprintf("%s:%s:%s", username, digestRealm, password))
	ha2 := client.md5Hash(fmt.Sprintf("%s:%s", http.MethodPost, serviceType))
	nonce := challengeValues["nonce"]
	qop := challengeValues["qop"]
	cnonce := client.newCNonce()
	nc := "1"
	response := client.md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, nonce, nc, cnonce, qop, ha2))
	authentication := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
		username, digestRealm, nonce, serviceType, cnonce, nc, qop, response)
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.cachedAuthentications == nil {
		client.cachedAuthentications = make(map[string]string)
	}
	client.cachedAuthentications[serviceType] = authentication
	return authentication, nil
}

func (client *Client) md5Hash(s string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func (client *Client) newCNonce() string {
	cnonceBytes := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, cnonceBytes)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%016x", cnonceBytes)
}
func (client *Client) postSoapActionRequest(endpoint string, action string, requestBody []byte, authentication string) (*http.Response, error) {
	if client.Debug {
		log.Printf("Invoking action %s on endpoint %s ...\n", action, endpoint)
	}
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request (cause: %w)", err)
	}
	request.Header.Add("Content-Type", "text/xml")
	request.Header.Add("SoapAction", action)
	if authentication != "" {
		request.Header.Add("Authorization", authentication)
	}
	response, err := client.httpClient().Do(request)
	if err != nil {
		return response, fmt.Errorf("failed to post request (cause: %w)", err)
	}
	if client.Debug {
		log.Println("Status: ", response.Status)
	}
	return response, nil
}

func (client *Client) httpClient() *http.Client {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.cachedHttpClient == nil {
		tlsClientConfig := &tls.Config{
			InsecureSkipVerify: client.InsecureSkipVerify,
		}
		transport := &http.Transport{
			TLSClientConfig: tlsClientConfig,
		}
		client.cachedHttpClient = &http.Client{
			Transport: transport,
		}
	}
	return client.cachedHttpClient
}
