// generated from spec version: 1.0
package x_homeplug

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetNumberOfDeviceEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDeviceEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDeviceEntriesResponse struct {
	XMLName            xml.Name `xml:"GetNumberOfDeviceEntriesResponse"`
	NewNumberOfEntries uint16   `xml:"NewNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfDeviceEntries(out *GetNumberOfDeviceEntriesResponse) error {
	in := &GetNumberOfDeviceEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfDeviceEntriesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetNumberOfDeviceEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfDeviceEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfDeviceEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDeviceEntries", soapRequest, soapResponse)
}

type GetGenericDeviceEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericDeviceEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericDeviceEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericDeviceEntryResponse"`
	NewMACAddress       string   `xml:"NewMACAddress"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
}

func (client *ServiceClient) GetGenericDeviceEntry(in *GetGenericDeviceEntryRequest, out *GetGenericDeviceEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericDeviceEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetGenericDeviceEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericDeviceEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericDeviceEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericDeviceEntry", soapRequest, soapResponse)
}

type GetSpecificDeviceEntryRequest struct {
	XMLName       xml.Name `xml:"u:GetSpecificDeviceEntryRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type GetSpecificDeviceEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificDeviceEntryResponse"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
}

func (client *ServiceClient) GetSpecificDeviceEntry(in *GetSpecificDeviceEntryRequest, out *GetSpecificDeviceEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificDeviceEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetSpecificDeviceEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificDeviceEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificDeviceEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificDeviceEntry", soapRequest, soapResponse)
}

type DeviceDoUpdateRequest struct {
	XMLName       xml.Name `xml:"u:DeviceDoUpdateRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type DeviceDoUpdateResponse struct {
	XMLName xml.Name `xml:"DeviceDoUpdateResponse"`
}

func (client *ServiceClient) DeviceDoUpdate(in *DeviceDoUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeviceDoUpdateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeviceDoUpdateRequest]{
			In: in,
		},
	}
	out := &DeviceDoUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[DeviceDoUpdateResponse]{
		Body: &tr064.SOAPResponseBody[DeviceDoUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeviceDoUpdate", soapRequest, soapResponse)
}