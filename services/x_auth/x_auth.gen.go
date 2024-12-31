// generated from spec version: 1.0
package x_auth

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName    xml.Name `xml:"GetInfoResponse"`
	NewEnabled bool     `xml:"NewEnabled"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type GetStateRequest struct {
	XMLName      xml.Name `xml:"u:GetStateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStateResponse struct {
	XMLName  xml.Name `xml:"GetStateResponse"`
	NewState string   `xml:"NewState"`
}

func (client *ServiceClient) GetState(out *GetStateResponse) error {
	in := &GetStateRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetStateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetStateRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetStateResponse]{
		Body: &tr064.SOAPResponseBody[GetStateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetState", soapRequest, soapResponse)
}

type SetConfigRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewAction    string   `xml:"NewAction"`
}

type SetConfigResponse struct {
	XMLName    xml.Name `xml:"SetConfigResponse"`
	NewState   string   `xml:"NewState"`
	NewToken   string   `xml:"NewToken"`
	NewMethods string   `xml:"NewMethods"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest, out *SetConfigResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetConfigRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[SetConfigResponse]{
		Body: &tr064.SOAPResponseBody[SetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", soapRequest, soapResponse)
}
