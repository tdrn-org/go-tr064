// generated from spec version: 1.0
package x_upnp

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
	XMLName            xml.Name `xml:"GetInfoResponse"`
	NewEnable          bool     `xml:"NewEnable"`
	NewUPnPMediaServer bool     `xml:"NewUPnPMediaServer"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
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

type SetConfigRequest struct {
	XMLName            xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewEnable          bool     `xml:"NewEnable"`
	NewUPnPMediaServer bool     `xml:"NewUPnPMediaServer"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetConfigRequest]{
			In: in,
		},
	}
	out := &SetConfigResponse{}
	soapResponse := &tr064.SOAPResponse[SetConfigResponse]{
		Body: &tr064.SOAPResponseBody[SetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", soapRequest, soapResponse)
}
