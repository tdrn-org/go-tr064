// generated from spec version: 1.0
package ethifconfig

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableResponse struct {
	XMLName xml.Name `xml:"SetEnableResponse"`
}

func (client *ServiceClient) SetEnable(in *SetEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetEnableRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetEnableRequest]{
			In: in,
		},
	}
	out := &SetEnableResponse{}
	soapResponse := &tr064.SOAPResponse[SetEnableResponse]{
		Body: &tr064.SOAPResponseBody[SetEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetEnable", soapRequest, soapResponse)
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName       xml.Name `xml:"GetInfoResponse"`
	NewEnable     bool     `xml:"NewEnable"`
	NewStatus     string   `xml:"NewStatus"`
	NewMACAddress string   `xml:"NewMACAddress"`
	NewMaxBitRate string   `xml:"NewMaxBitRate"`
	NewDuplexMode string   `xml:"NewDuplexMode"`
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

type GetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsResponse struct {
	XMLName            xml.Name `xml:"GetStatisticsResponse"`
	NewBytesSent       uint32   `xml:"NewBytesSent"`
	NewBytesReceived   uint32   `xml:"NewBytesReceived"`
	NewPacketsSent     uint32   `xml:"NewPacketsSent"`
	NewPacketsReceived uint32   `xml:"NewPacketsReceived"`
}

func (client *ServiceClient) GetStatistics(out *GetStatisticsResponse) error {
	in := &GetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetStatisticsRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetStatisticsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetStatisticsResponse]{
		Body: &tr064.SOAPResponseBody[GetStatisticsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetStatistics", soapRequest, soapResponse)
}
