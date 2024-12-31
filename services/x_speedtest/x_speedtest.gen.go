// generated from spec version: 1.0
package x_speedtest

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
	XMLName              xml.Name `xml:"GetInfoResponse"`
	NewEnableTcp         bool     `xml:"NewEnableTcp"`
	NewEnableUdp         bool     `xml:"NewEnableUdp"`
	NewEnableUdpBidirect bool     `xml:"NewEnableUdpBidirect"`
	NewWANEnableTcp      bool     `xml:"NewWANEnableTcp"`
	NewWANEnableUdp      bool     `xml:"NewWANEnableUdp"`
	NewPortTcp           uint32   `xml:"NewPortTcp"`
	NewPortUdp           uint32   `xml:"NewPortUdp"`
	NewPortUdpBidirect   uint32   `xml:"NewPortUdpBidirect"`
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
	XMLName              xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewEnableTcp         bool     `xml:"NewEnableTcp"`
	NewEnableUdp         bool     `xml:"NewEnableUdp"`
	NewEnableUdpBidirect bool     `xml:"NewEnableUdpBidirect"`
	NewWANEnableTcp      bool     `xml:"NewWANEnableTcp"`
	NewWANEnableUdp      bool     `xml:"NewWANEnableUdp"`
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

type GetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsResponse struct {
	XMLName         xml.Name `xml:"GetStatisticsResponse"`
	NewByteCount    uint32   `xml:"NewByteCount"`
	NewKbitsCurrent uint32   `xml:"NewKbitsCurrent"`
	NewKbitsAvg     uint32   `xml:"NewKbitsAvg"`
	NewPacketCount  uint32   `xml:"NewPacketCount"`
	NewPPSCurrent   uint32   `xml:"NewPPSCurrent"`
	NewPPSAvg       uint32   `xml:"NewPPSAvg"`
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

type ResetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:ResetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type ResetStatisticsResponse struct {
	XMLName xml.Name `xml:"ResetStatisticsResponse"`
}

func (client *ServiceClient) ResetStatistics() error {
	in := &ResetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[ResetStatisticsRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[ResetStatisticsRequest]{
			In: in,
		},
	}
	out := &ResetStatisticsResponse{}
	soapResponse := &tr064.SOAPResponse[ResetStatisticsResponse]{
		Body: &tr064.SOAPResponseBody[ResetStatisticsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "ResetStatistics", soapRequest, soapResponse)
}
