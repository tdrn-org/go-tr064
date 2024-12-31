// generated from spec version: 1.0
package x_dect

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetNumberOfDectEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDectEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDectEntriesResponse struct {
	XMLName            xml.Name `xml:"GetNumberOfDectEntriesResponse"`
	NewNumberOfEntries uint16   `xml:"NewNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfDectEntries(out *GetNumberOfDectEntriesResponse) error {
	in := &GetNumberOfDectEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfDectEntriesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetNumberOfDectEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfDectEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfDectEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDectEntries", soapRequest, soapResponse)
}

type GetGenericDectEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericDectEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericDectEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericDectEntryResponse"`
	NewID               string   `xml:"NewID"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
	NewUpdateInfo       string   `xml:"NewUpdateInfo"`
}

func (client *ServiceClient) GetGenericDectEntry(in *GetGenericDectEntryRequest, out *GetGenericDectEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericDectEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetGenericDectEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericDectEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericDectEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericDectEntry", soapRequest, soapResponse)
}

type GetSpecificDectEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetSpecificDectEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type GetSpecificDectEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificDectEntryResponse"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
	NewUpdateInfo       string   `xml:"NewUpdateInfo"`
}

func (client *ServiceClient) GetSpecificDectEntry(in *GetSpecificDectEntryRequest, out *GetSpecificDectEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificDectEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetSpecificDectEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificDectEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificDectEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificDectEntry", soapRequest, soapResponse)
}

type DectDoUpdateRequest struct {
	XMLName      xml.Name `xml:"u:DectDoUpdateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type DectDoUpdateResponse struct {
	XMLName xml.Name `xml:"DectDoUpdateResponse"`
}

func (client *ServiceClient) DectDoUpdate(in *DectDoUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DectDoUpdateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DectDoUpdateRequest]{
			In: in,
		},
	}
	out := &DectDoUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[DectDoUpdateResponse]{
		Body: &tr064.SOAPResponseBody[DectDoUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DectDoUpdate", soapRequest, soapResponse)
}

type GetDectListPathRequest struct {
	XMLName      xml.Name `xml:"u:GetDectListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDectListPathResponse struct {
	XMLName         xml.Name `xml:"GetDectListPathResponse"`
	NewDectListPath string   `xml:"NewDectListPath"`
}

func (client *ServiceClient) GetDectListPath(out *GetDectListPathResponse) error {
	in := &GetDectListPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDectListPathRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDectListPathRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDectListPathResponse]{
		Body: &tr064.SOAPResponseBody[GetDectListPathResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDectListPath", soapRequest, soapResponse)
}