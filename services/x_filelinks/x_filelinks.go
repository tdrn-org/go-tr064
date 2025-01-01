// generated from spec version: 1.0
package x_filelinks

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetNumberOfFilelinkEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfFilelinkEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfFilelinkEntriesResponse struct {
	XMLName            xml.Name `xml:"GetNumberOfFilelinkEntriesResponse"`
	NewNumberOfEntries uint16   `xml:"NewNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfFilelinkEntries(out *GetNumberOfFilelinkEntriesResponse) error {
	in := &GetNumberOfFilelinkEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetNumberOfFilelinkEntriesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetNumberOfFilelinkEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetNumberOfFilelinkEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetNumberOfFilelinkEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfFilelinkEntries", soapRequest, soapResponse)
}

type GetGenericFilelinkEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericFilelinkEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericFilelinkEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericFilelinkEntryResponse"`
	NewID               string   `xml:"NewID"`
	NewValid            bool     `xml:"NewValid"`
	NewPath             string   `xml:"NewPath"`
	NewIsDirectory      bool     `xml:"NewIsDirectory"`
	NewUrl              string   `xml:"NewUrl"`
	NewUsername         string   `xml:"NewUsername"`
	NewAccessCountLimit uint16   `xml:"NewAccessCountLimit"`
	NewAccessCount      uint16   `xml:"NewAccessCount"`
	NewExpire           uint16   `xml:"NewExpire"`
	NewExpireDate       string   `xml:"NewExpireDate"`
}

func (client *ServiceClient) GetGenericFilelinkEntry(in *GetGenericFilelinkEntryRequest, out *GetGenericFilelinkEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericFilelinkEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetGenericFilelinkEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericFilelinkEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericFilelinkEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericFilelinkEntry", soapRequest, soapResponse)
}

type GetSpecificFilelinkEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetSpecificFilelinkEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type GetSpecificFilelinkEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificFilelinkEntryResponse"`
	NewValid            bool     `xml:"NewValid"`
	NewPath             string   `xml:"NewPath"`
	NewIsDirectory      bool     `xml:"NewIsDirectory"`
	NewUrl              string   `xml:"NewUrl"`
	NewUsername         string   `xml:"NewUsername"`
	NewAccessCountLimit uint16   `xml:"NewAccessCountLimit"`
	NewAccessCount      uint16   `xml:"NewAccessCount"`
	NewExpire           uint16   `xml:"NewExpire"`
	NewExpireDate       string   `xml:"NewExpireDate"`
}

func (client *ServiceClient) GetSpecificFilelinkEntry(in *GetSpecificFilelinkEntryRequest, out *GetSpecificFilelinkEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificFilelinkEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetSpecificFilelinkEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificFilelinkEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificFilelinkEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificFilelinkEntry", soapRequest, soapResponse)
}

type NewFilelinkEntryRequest struct {
	XMLName             xml.Name `xml:"u:NewFilelinkEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPath             string   `xml:"NewPath"`
	NewAccessCountLimit uint16   `xml:"NewAccessCountLimit"`
	NewExpire           uint16   `xml:"NewExpire"`
}

type NewFilelinkEntryResponse struct {
	XMLName xml.Name `xml:"NewFilelinkEntryResponse"`
	NewID   string   `xml:"NewID"`
}

func (client *ServiceClient) NewFilelinkEntry(in *NewFilelinkEntryRequest, out *NewFilelinkEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[NewFilelinkEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[NewFilelinkEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[NewFilelinkEntryResponse]{
		Body: &tr064.SOAPResponseBody[NewFilelinkEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "NewFilelinkEntry", soapRequest, soapResponse)
}

type SetFilelinkEntryRequest struct {
	XMLName             xml.Name `xml:"u:SetFilelinkEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewID               string   `xml:"NewID"`
	NewAccessCountLimit uint16   `xml:"NewAccessCountLimit"`
	NewExpire           uint16   `xml:"NewExpire"`
}

type SetFilelinkEntryResponse struct {
	XMLName xml.Name `xml:"SetFilelinkEntryResponse"`
}

func (client *ServiceClient) SetFilelinkEntry(in *SetFilelinkEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetFilelinkEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetFilelinkEntryRequest]{
			In: in,
		},
	}
	out := &SetFilelinkEntryResponse{}
	soapResponse := &tr064.SOAPResponse[SetFilelinkEntryResponse]{
		Body: &tr064.SOAPResponseBody[SetFilelinkEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetFilelinkEntry", soapRequest, soapResponse)
}

type DeleteFilelinkEntryRequest struct {
	XMLName      xml.Name `xml:"u:DeleteFilelinkEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type DeleteFilelinkEntryResponse struct {
	XMLName xml.Name `xml:"DeleteFilelinkEntryResponse"`
}

func (client *ServiceClient) DeleteFilelinkEntry(in *DeleteFilelinkEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[DeleteFilelinkEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[DeleteFilelinkEntryRequest]{
			In: in,
		},
	}
	out := &DeleteFilelinkEntryResponse{}
	soapResponse := &tr064.SOAPResponse[DeleteFilelinkEntryResponse]{
		Body: &tr064.SOAPResponseBody[DeleteFilelinkEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "DeleteFilelinkEntry", soapRequest, soapResponse)
}

type GetFilelinkListPathRequest struct {
	XMLName      xml.Name `xml:"u:GetFilelinkListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetFilelinkListPathResponse struct {
	XMLName             xml.Name `xml:"GetFilelinkListPathResponse"`
	NewFilelinkListPath string   `xml:"NewFilelinkListPath"`
}

func (client *ServiceClient) GetFilelinkListPath(out *GetFilelinkListPathResponse) error {
	in := &GetFilelinkListPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetFilelinkListPathRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetFilelinkListPathRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetFilelinkListPathResponse]{
		Body: &tr064.SOAPResponseBody[GetFilelinkListPathResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetFilelinkListPath", soapRequest, soapResponse)
}
