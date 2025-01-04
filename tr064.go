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

//go:generate go run cmd/build/build.go generate

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const XMLNameSpace = "http://schemas.xmlsoap.org/soap/envelope/"

const XMLEncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"

type TR064Spec string

const (
	DefaultSpec = "tr64desc"
	IGDSpec     = "igddesc"
)

func (spec TR064Spec) Name() string {
	return string(spec)
}

func (spec TR064Spec) Path() string {
	return "/" + string(spec) + ".xml"
}

func unmarshalDocument(url *url.URL, v any) error {
	response, err := http.Get(url.String())
	if err != nil {
		return fmt.Errorf("failed to access URL '%s' (cause: %w)", url, err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get URL '%s' (status: %s)", url, response.Status)
	}
	document := response.Body
	defer document.Close()
	documentBytes, err := io.ReadAll(document)
	if err != nil {
		return fmt.Errorf("failed to read URL '%s' (cause: %w)", url, err)
	}
	err = xml.Unmarshal(documentBytes, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal URL '%s' (cause: %w)", url, err)
	}
	return nil
}

type tr64descDoc struct {
	SpecVersion   specVersionDoc   `xml:"specVersion"`
	SystemVersion systemVersionDoc `xml:"systemVersion"`
	Device        deviceDoc        `xml:"device"`
}

type WalkServiceFunc func(*serviceDoc, *scpdDoc) error

func (doc *tr64descDoc) walk(baseUrl *url.URL, f WalkServiceFunc) error {
	return doc.Device.walk(baseUrl, f)
}

type specVersionDoc struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

func (doc *specVersionDoc) signature() string {
	return fmt.Sprintf("spec version: %d.%d", doc.Major, doc.Minor)
}

type systemVersionDoc struct {
	HW          int    `xml:"HW"`
	Major       int    `xml:"Major"`
	Minor       int    `xml:"Minor"`
	Patch       int    `xml:"Patch"`
	Buildnumber int    `xml:"Buildnumber"`
	Display     string `xml:"Display"`
}

type deviceListDoc struct {
	Devices []deviceDoc `xml:"device"`
}

type deviceDoc struct {
	DeviceType       string         `xml:"deviceType"`
	FriendlyName     string         `xml:"friendlyName"`
	Manufacturer     string         `xml:"manufacturer"`
	ManufacturerURL  string         `xml:"manufacturerURL"`
	ModelDescription string         `xml:"modelDescription"`
	ModelName        string         `xml:"modelName"`
	ModelNumber      string         `xml:"modelNumber"`
	ModelURL         string         `xml:"modelURL"`
	UDN              string         `xml:"UDN"`
	SerialNumber     string         `xml:"serialNumber"`
	OriginUDN        string         `xml:"originUDN"`
	ServiceList      serviceListDoc `xml:"serviceList"`
	DeviceList       deviceListDoc  `xml:"deviceList"`
	PresentationURL  string         `xml:"presentationURL"`
}

func (doc *deviceDoc) walk(baseUrl *url.URL, f WalkServiceFunc) error {
	for _, service := range doc.ServiceList.Services {
		if strings.HasPrefix(service.ServiceType, "urn:schemas-any-com:service:Any:") {
			continue
		}
		scpd, err := service.scpd(baseUrl)
		if err != nil {
			return err
		}
		err = f(&service, scpd)
		if err != nil {
			return err
		}
	}
	for _, device := range doc.DeviceList.Devices {
		err := device.walk(baseUrl, f)
		if err != nil {
			return err
		}
	}
	return nil
}

type serviceListDoc struct {
	Services []serviceDoc `xml:"service"`
}

type serviceDoc struct {
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	ControlURL  string   `xml:"controlURL"`
	EventSubURL string   `xml:"eventSubURL"`
	SCPDURL     string   `xml:"SCPDURL"`
	cachedSCPD  *scpdDoc `xml:"-"`
}

func (doc *serviceDoc) scpd(baseUrl *url.URL) (*scpdDoc, error) {
	if doc.cachedSCPD == nil {
		scpdUrl := baseUrl.JoinPath(doc.SCPDURL)
		scpd := &scpdDoc{}
		err := unmarshalDocument(scpdUrl, scpd)
		if err != nil {
			return nil, err
		}
		doc.cachedSCPD = scpd
	}
	return doc.cachedSCPD, nil
}

var serviceSCPDPattern = regexp.MustCompile(`^/(.+)SCPD.xml$`)

func (service *serviceDoc) scpdName() string {
	match := serviceSCPDPattern.FindStringSubmatch(service.SCPDURL)
	if match == nil {
		log.Fatal(fmt.Errorf("unexpected SCPD URL '%s'", service.SCPDURL))
	}
	return match[1]
}

var serviceTypeNamePattern = regexp.MustCompile(`^urn\:(.+)\:service\:(.+):\d+$`)

func (service *serviceDoc) Name() string {
	match := serviceTypeNamePattern.FindStringSubmatch(service.ServiceType)
	if match == nil {
		log.Fatal(fmt.Errorf("unexpected service type '%s'", service.ServiceType))
	}
	mangledServiceName := mangleName(match[2])
	return mangledServiceName
}

func (service *serviceDoc) Type() string {
	return service.ServiceType
}

func (service *serviceDoc) Id() string {
	return service.ServiceId
}

func (service *serviceDoc) Url() string {
	return service.ControlURL
}

type scpdDoc struct {
	SpecVersion       specVersionDoc       `xml:"specVersion"`
	ActionList        actionListDoc        `xml:"actionList"`
	ServiceStateTable serviceStateTableDoc `xml:"serviceStateTable"`
}

func (doc *scpdDoc) lookupVariable(name string) *stateVariableDoc {
	for _, variable := range doc.ServiceStateTable.StateVariables {
		if variable.Name == name {
			return &variable
		}
	}
	return nil
}

type actionListDoc struct {
	Actions []actionDoc `xml:"action"`
}

type actionDoc struct {
	Name         string          `xml:"name"`
	ArgumentList argumentListDoc `xml:"argumentList"`
}

type argumentListDoc struct {
	Arguments []argumentDoc `xml:"argument"`
}

func (doc *argumentListDoc) hasIn() bool {
	for _, argument := range doc.Arguments {
		if argument.Direction == "in" {
			return true
		}
	}
	return false
}

func (doc *argumentListDoc) hasOut() bool {
	for _, argument := range doc.Arguments {
		if argument.Direction == "out" {
			return true
		}
	}
	return false
}

type argumentDoc struct {
	Name                 string `xml:"name"`
	Direction            string `xml:"direction"`
	RelatedStateVariable string `xml:"relatedStateVariable"`
}

type serviceStateTableDoc struct {
	StateVariables []stateVariableDoc `xml:"stateVariable"`
}

type stateVariableDoc struct {
	Name          string              `xml:"name"`
	DataType      string              `xml:"dataType"`
	DefaultValue  string              `xml:"defaultValue"`
	AllowedValues allowedValueListDoc `xml:"allowedValueList"`
}

type allowedValueListDoc struct {
	AllowedValues []string `xml:"allowedValue"`
}
