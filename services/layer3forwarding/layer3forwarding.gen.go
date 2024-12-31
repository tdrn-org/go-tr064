// generated from spec version: 1.0
package layer3forwarding

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetDefaultConnectionServiceRequest struct {
	XMLName                     xml.Name `xml:"u:SetDefaultConnectionServiceRequest"`
	XMLNameSpace                string   `xml:"xmlns:u,attr"`
	NewDefaultConnectionService string   `xml:"NewDefaultConnectionService"`
}

type SetDefaultConnectionServiceResponse struct {
	XMLName xml.Name `xml:"SetDefaultConnectionServiceResponse"`
}

func (client *ServiceClient) SetDefaultConnectionService(in *SetDefaultConnectionServiceRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetDefaultConnectionServiceRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetDefaultConnectionServiceRequest]{
			In: in,
		},
	}
	out := &SetDefaultConnectionServiceResponse{}
	soapResponse := &tr064.SOAPResponse[SetDefaultConnectionServiceResponse]{
		Body: &tr064.SOAPResponseBody[SetDefaultConnectionServiceResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetDefaultConnectionService", soapRequest, soapResponse)
}

type GetDefaultConnectionServiceRequest struct {
	XMLName      xml.Name `xml:"u:GetDefaultConnectionServiceRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDefaultConnectionServiceResponse struct {
	XMLName                     xml.Name `xml:"GetDefaultConnectionServiceResponse"`
	NewDefaultConnectionService string   `xml:"NewDefaultConnectionService"`
}

func (client *ServiceClient) GetDefaultConnectionService(out *GetDefaultConnectionServiceResponse) error {
	in := &GetDefaultConnectionServiceRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetDefaultConnectionServiceRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetDefaultConnectionServiceRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetDefaultConnectionServiceResponse]{
		Body: &tr064.SOAPResponseBody[GetDefaultConnectionServiceResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetDefaultConnectionService", soapRequest, soapResponse)
}

type GetForwardNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetForwardNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetForwardNumberOfEntriesResponse struct {
	XMLName                   xml.Name `xml:"GetForwardNumberOfEntriesResponse"`
	NewForwardNumberOfEntries uint16   `xml:"NewForwardNumberOfEntries"`
}

func (client *ServiceClient) GetForwardNumberOfEntries(out *GetForwardNumberOfEntriesResponse) error {
	in := &GetForwardNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetForwardNumberOfEntriesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetForwardNumberOfEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetForwardNumberOfEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetForwardNumberOfEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetForwardNumberOfEntries", soapRequest, soapResponse)
}

type AddForwardingEntryRequest struct {
	XMLName             xml.Name `xml:"u:AddForwardingEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewType             string   `xml:"NewType"`
	NewDestIPAddress    string   `xml:"NewDestIPAddress"`
	NewDestSubnetMask   string   `xml:"NewDestSubnetMask"`
	NewSourceIPAddress  string   `xml:"NewSourceIPAddress"`
	NewSourceSubnetMask string   `xml:"NewSourceSubnetMask"`
	NewGatewayIPAddress string   `xml:"NewGatewayIPAddress"`
	NewInterface        string   `xml:"NewInterface"`
	NewForwardingMetric int32    `xml:"NewForwardingMetric"`
}

type AddForwardingEntryResponse struct {
	XMLName xml.Name `xml:"AddForwardingEntryResponse"`
}

func (client *ServiceClient) AddForwardingEntry(in *AddForwardingEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[AddForwardingEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[AddForwardingEntryRequest]{
			In: in,
		},
	}
	out := &AddForwardingEntryResponse{}
	soapResponse := &tr064.SOAPResponse[AddForwardingEntryResponse]{
		Body: &tr064.SOAPResponseBody[AddForwardingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "AddForwardingEntry", soapRequest, soapResponse)
}

type DeleteForwardingEntryRequest struct {
	XMLName             xml.Name `xml:"u:DeleteForwardingEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewDestIPAddress    string   `xml:"NewDestIPAddress"`
	NewDestSubnetMask   string   `xml:"NewDestSubnetMask"`
	NewSourceIPAddress  string   `xml:"NewSourceIPAddress"`
	NewSourceSubnetMask string   `xml:"NewSourceSubnetMask"`
}

type DeleteForwardingEntryResponse struct {
	XMLName xml.Name `xml:"DeleteForwardingEntryResponse"`
}

func (client *ServiceClient) DeleteForwardingEntry(in *DeleteForwardingEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeleteForwardingEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[DeleteForwardingEntryRequest]{
			In: in,
		},
	}
	out := &DeleteForwardingEntryResponse{}
	soapResponse := &tr064.SOAPResponse[DeleteForwardingEntryResponse]{
		Body: &tr064.SOAPResponseBody[DeleteForwardingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeleteForwardingEntry", soapRequest, soapResponse)
}

type GetSpecificForwardingEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetSpecificForwardingEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewDestIPAddress    string   `xml:"NewDestIPAddress"`
	NewDestSubnetMask   string   `xml:"NewDestSubnetMask"`
	NewSourceIPAddress  string   `xml:"NewSourceIPAddress"`
	NewSourceSubnetMask string   `xml:"NewSourceSubnetMask"`
}

type GetSpecificForwardingEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificForwardingEntryResponse"`
	NewGatewayIPAddress string   `xml:"NewGatewayIPAddress"`
	NewEnable           bool     `xml:"NewEnable"`
	NewStatus           string   `xml:"NewStatus"`
	NewType             string   `xml:"NewType"`
	NewInterface        string   `xml:"NewInterface"`
	NewForwardingMetric int32    `xml:"NewForwardingMetric"`
}

func (client *ServiceClient) GetSpecificForwardingEntry(in *GetSpecificForwardingEntryRequest, out *GetSpecificForwardingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificForwardingEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetSpecificForwardingEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificForwardingEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificForwardingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificForwardingEntry", soapRequest, soapResponse)
}

type GetGenericForwardingEntryRequest struct {
	XMLName            xml.Name `xml:"u:GetGenericForwardingEntryRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewForwardingIndex uint16   `xml:"NewForwardingIndex"`
}

type GetGenericForwardingEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericForwardingEntryResponse"`
	NewEnable           bool     `xml:"NewEnable"`
	NewStatus           string   `xml:"NewStatus"`
	NewType             string   `xml:"NewType"`
	NewDestIPAddress    string   `xml:"NewDestIPAddress"`
	NewDestSubnetMask   string   `xml:"NewDestSubnetMask"`
	NewSourceIPAddress  string   `xml:"NewSourceIPAddress"`
	NewSourceSubnetMask string   `xml:"NewSourceSubnetMask"`
	NewGatewayIPAddress string   `xml:"NewGatewayIPAddress"`
	NewInterface        string   `xml:"NewInterface"`
	NewForwardingMetric int32    `xml:"NewForwardingMetric"`
}

func (client *ServiceClient) GetGenericForwardingEntry(in *GetGenericForwardingEntryRequest, out *GetGenericForwardingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericForwardingEntryRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetGenericForwardingEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericForwardingEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericForwardingEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericForwardingEntry", soapRequest, soapResponse)
}

type SetForwardingEntryEnableRequest struct {
	XMLName             xml.Name `xml:"u:SetForwardingEntryEnableRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewDestIPAddress    string   `xml:"NewDestIPAddress"`
	NewDestSubnetMask   string   `xml:"NewDestSubnetMask"`
	NewSourceIPAddress  string   `xml:"NewSourceIPAddress"`
	NewSourceSubnetMask string   `xml:"NewSourceSubnetMask"`
	NewEnable           bool     `xml:"NewEnable"`
}

type SetForwardingEntryEnableResponse struct {
	XMLName xml.Name `xml:"SetForwardingEntryEnableResponse"`
}

func (client *ServiceClient) SetForwardingEntryEnable(in *SetForwardingEntryEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetForwardingEntryEnableRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetForwardingEntryEnableRequest]{
			In: in,
		},
	}
	out := &SetForwardingEntryEnableResponse{}
	soapResponse := &tr064.SOAPResponse[SetForwardingEntryEnableResponse]{
		Body: &tr064.SOAPResponseBody[SetForwardingEntryEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetForwardingEntryEnable", soapRequest, soapResponse)
}
