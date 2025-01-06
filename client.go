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

type ServiceDescriptor interface {
	Spec() ServiceSpec
	Type() string
	ShortType() string
	Id() string
	Url() string
}

type StaticServiceDescriptor struct {
	ServiceSpec ServiceSpec
	ServiceType string
	ServiceId   string
	ServiceUrl  string
}

func (service *StaticServiceDescriptor) Spec() ServiceSpec {
	return service.ServiceSpec
}

func (service *StaticServiceDescriptor) Type() string {
	return service.ServiceType
}

func (service *StaticServiceDescriptor) ShortType() string {
	return serviceShortType(service.ServiceType)
}

func (service *StaticServiceDescriptor) Id() string {
	return service.ServiceId
}

func (service *StaticServiceDescriptor) Url() string {
	return service.ServiceUrl
}

// NewClient instantiates a new TR-064 client for accessing the given URL.
//
// If the given URL contains a userinfo, the contained username and password are automatically used for authentication.
func NewClient(deviceUrl *url.URL, specs ...ServiceSpec) *Client {
	anonymousDeviceUrl := *deviceUrl
	anonymousDeviceUrl.User = nil
	username := deviceUrl.User.Username()
	password, _ := deviceUrl.User.Password()
	client := &Client{
		DeviceUrl: &anonymousDeviceUrl,
		Username:  username,
		Password:  password,
		Specs:     specs,
		mutex:     &sync.Mutex{},
	}
	client.cachedHttpClient = sync.OnceValue(client.httpClient)
	return client
}

// Client provides the necessary parameters to access a TR-064 capable server and perform service discovery.
//
// To access an actual service, this client as well as the desired service descriptor is combined into a
// service specific service client:
//
//	client := tr064.NewClient(deviceUrl)
//	services, _ := client.ServiceByName(deviceinfo.ServiceName)
//	serviceClient := deviceinfo.ServiceClient {
//		TR064Client: client,
//		Service:     services[0],
//	}
//	info := &deviceinfo.GetInfoResponse{}
//	_ = serviceClient.GetInfo(info)
//
// The service client is then used to access the individual service functions.
type Client struct {
	// Url defines the URL to access the TR-064 server.
	DeviceUrl *url.URL
	// Username is set to the login to use for accessing restricted services.
	Username string
	// Password is set to the password to use for accessing restricted services.
	Password string
	// Specs lists the TR-064 specs describing the TR-064 server's services. If empty [DefaultServiceSpec] is assumed.
	Specs []ServiceSpec
	// Timeout sets the timeout for HTTP(S) communication.
	Timeout time.Duration
	// InsecureSkipVerify disables certificate validation while access the TR-064 server. Use with care.
	InsecureSkipVerify bool
	// Debug enables debug logging while accessing the TR-064 server.
	Debug                 bool
	mutex                 *sync.Mutex
	cachedServices        []ServiceDescriptor
	cachedAuthentications map[string]string
	cachedHttpClient      func() *http.Client
}

// Services fetches and parses the TR-064 server's service specifications and returns the available services.
func (client *Client) Services() ([]ServiceDescriptor, error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.cachedServices == nil {
		specs := client.Specs
		if len(specs) == 0 {
			specs = []ServiceSpec{DefaultServiceSpec}
		}
		services := make([]ServiceDescriptor, 0)
		httpClient := client.cachedHttpClient()
		for _, spec := range specs {
			tr64desc, err := fetchServiceSpec(httpClient, client.DeviceUrl, spec)
			if err != nil {
				return nil, err
			}
			collector := &serviceCollector{serviceMap: make(map[string]ServiceDescriptor)}
			err = tr64desc.walk(httpClient, client.DeviceUrl, collector.collectService)
			if err != nil {
				return nil, err
			}
			services = append(services, slices.Collect(maps.Values(collector.serviceMap))...)
		}
		slices.SortFunc(services, func(a ServiceDescriptor, b ServiceDescriptor) int { return strings.Compare(a.Type(), b.Type()) })
		client.cachedServices = services
	}
	return client.cachedServices, nil
}

type serviceCollector struct {
	serviceMap map[string]ServiceDescriptor
}

func (collector *serviceCollector) collectService(service *serviceDoc, scpd *scpdDoc) error {
	collector.serviceMap[service.ServiceId] = service
	return nil
}

// ServicesByType fetches and parses the TR-064 server's service specifications
// like [Services], but returns only the services matching the given spec and service type.
func (client *Client) ServicesByType(spec ServiceSpec, serviceType string) ([]ServiceDescriptor, error) {
	all, err := client.Services()
	if err != nil {
		return nil, err
	}
	services := make([]ServiceDescriptor, 0)
	for _, service := range all {
		if service.Spec() == spec && (service.Type() == serviceType || service.ShortType() == serviceType) {
			services = append(services, service)
		}
	}
	return services, nil
}

// Get performs a simple GET toward the TR-064 server using the given path.
func (client *Client) Get(ref string) (*http.Response, error) {
	refUrl, err := url.Parse(ref)
	if err != nil {
		return nil, fmt.Errorf("failed to parse reference: '%s' (cause: %w)", ref, err)
	}
	targetUrl := client.DeviceUrl.ResolveReference(refUrl)
	return client.cachedHttpClient().Get(targetUrl.String())
}

func NewSOAPRequest[T any](in *T) *SOAPRequest[T] {
	return &SOAPRequest[T]{
		XMLNameSpace:     XMLNameSpace,
		XMLEncodingStyle: XMLEncodingStyle,
		Body: &SOAPRequestBody[T]{
			In: in,
		},
	}
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

func NewSOAPResponse[T any](out *T) *SOAPResponse[T] {
	return &SOAPResponse[T]{
		Body: &SOAPResponseBody[T]{
			Out: out,
		},
	}
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
	endpoint := client.DeviceUrl.ResolveReference(controlUrl).String()
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
	ha1 := client.md5Hash(fmt.Sprintf("%s:%s:%s", client.Username, digestRealm, client.Password))
	ha2 := client.md5Hash(fmt.Sprintf("%s:%s", http.MethodPost, serviceType))
	nonce := challengeValues["nonce"]
	qop := challengeValues["qop"]
	cnonce := client.newCNonce()
	nc := "1"
	response := client.md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, nonce, nc, cnonce, qop, ha2))
	authentication := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
		client.Username, digestRealm, nonce, serviceType, cnonce, nc, qop, response)
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
	response, err := client.cachedHttpClient().Do(request)
	if err != nil {
		return response, fmt.Errorf("failed to post request (cause: %w)", err)
	}
	if client.Debug {
		log.Println("Status: ", response.Status)
	}
	return response, nil
}

func (client *Client) httpClient() *http.Client {
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: client.InsecureSkipVerify,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsClientConfig,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   client.Timeout,
	}
}
