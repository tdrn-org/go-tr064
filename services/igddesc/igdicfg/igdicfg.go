// generated from spec version: 1.0
package igdicfg

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetCommonLinkPropertiesRequest struct {
	XMLName      xml.Name `xml:"u:GetCommonLinkPropertiesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetCommonLinkPropertiesResponse struct {
	XMLName                       xml.Name `xml:"GetCommonLinkPropertiesResponse"`
	NewWANAccessType              string   `xml:"NewWANAccessType"`
	NewLayer1UpstreamMaxBitRate   uint32   `xml:"NewLayer1UpstreamMaxBitRate"`
	NewLayer1DownstreamMaxBitRate uint32   `xml:"NewLayer1DownstreamMaxBitRate"`
	NewPhysicalLinkStatus         string   `xml:"NewPhysicalLinkStatus"`
}

func (client *ServiceClient) GetCommonLinkProperties(out *GetCommonLinkPropertiesResponse) error {
	in := &GetCommonLinkPropertiesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetCommonLinkPropertiesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetCommonLinkPropertiesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetCommonLinkPropertiesResponse]{
		Body: &tr064.SOAPResponseBody[GetCommonLinkPropertiesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetCommonLinkProperties", soapRequest, soapResponse)
}

type GetTotalBytesSentRequest struct {
	XMLName      xml.Name `xml:"u:GetTotalBytesSentRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetTotalBytesSentResponse struct {
	XMLName           xml.Name `xml:"GetTotalBytesSentResponse"`
	NewTotalBytesSent uint32   `xml:"NewTotalBytesSent"`
}

func (client *ServiceClient) GetTotalBytesSent(out *GetTotalBytesSentResponse) error {
	in := &GetTotalBytesSentRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetTotalBytesSentRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetTotalBytesSentRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetTotalBytesSentResponse]{
		Body: &tr064.SOAPResponseBody[GetTotalBytesSentResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetTotalBytesSent", soapRequest, soapResponse)
}

type GetTotalBytesReceivedRequest struct {
	XMLName      xml.Name `xml:"u:GetTotalBytesReceivedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetTotalBytesReceivedResponse struct {
	XMLName               xml.Name `xml:"GetTotalBytesReceivedResponse"`
	NewTotalBytesReceived uint32   `xml:"NewTotalBytesReceived"`
}

func (client *ServiceClient) GetTotalBytesReceived(out *GetTotalBytesReceivedResponse) error {
	in := &GetTotalBytesReceivedRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetTotalBytesReceivedRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetTotalBytesReceivedRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetTotalBytesReceivedResponse]{
		Body: &tr064.SOAPResponseBody[GetTotalBytesReceivedResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetTotalBytesReceived", soapRequest, soapResponse)
}

type GetTotalPacketsSentRequest struct {
	XMLName      xml.Name `xml:"u:GetTotalPacketsSentRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetTotalPacketsSentResponse struct {
	XMLName             xml.Name `xml:"GetTotalPacketsSentResponse"`
	NewTotalPacketsSent uint32   `xml:"NewTotalPacketsSent"`
}

func (client *ServiceClient) GetTotalPacketsSent(out *GetTotalPacketsSentResponse) error {
	in := &GetTotalPacketsSentRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetTotalPacketsSentRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetTotalPacketsSentRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetTotalPacketsSentResponse]{
		Body: &tr064.SOAPResponseBody[GetTotalPacketsSentResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetTotalPacketsSent", soapRequest, soapResponse)
}

type GetTotalPacketsReceivedRequest struct {
	XMLName      xml.Name `xml:"u:GetTotalPacketsReceivedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetTotalPacketsReceivedResponse struct {
	XMLName                 xml.Name `xml:"GetTotalPacketsReceivedResponse"`
	NewTotalPacketsReceived uint32   `xml:"NewTotalPacketsReceived"`
}

func (client *ServiceClient) GetTotalPacketsReceived(out *GetTotalPacketsReceivedResponse) error {
	in := &GetTotalPacketsReceivedRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetTotalPacketsReceivedRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetTotalPacketsReceivedRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetTotalPacketsReceivedResponse]{
		Body: &tr064.SOAPResponseBody[GetTotalPacketsReceivedResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetTotalPacketsReceived", soapRequest, soapResponse)
}

type GetAddonInfosRequest struct {
	XMLName      xml.Name `xml:"u:GetAddonInfosRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAddonInfosResponse struct {
	XMLName                          xml.Name `xml:"GetAddonInfosResponse"`
	NewByteSendRate                  uint32   `xml:"NewByteSendRate"`
	NewByteReceiveRate               uint32   `xml:"NewByteReceiveRate"`
	NewPacketSendRate                uint32   `xml:"NewPacketSendRate"`
	NewPacketReceiveRate             uint32   `xml:"NewPacketReceiveRate"`
	NewTotalBytesSent                uint32   `xml:"NewTotalBytesSent"`
	NewTotalBytesReceived            uint32   `xml:"NewTotalBytesReceived"`
	NewAutoDisconnectTime            uint32   `xml:"NewAutoDisconnectTime"`
	NewIdleDisconnectTime            uint32   `xml:"NewIdleDisconnectTime"`
	NewDNSServer1                    string   `xml:"NewDNSServer1"`
	NewDNSServer2                    string   `xml:"NewDNSServer2"`
	NewVoipDNSServer1                string   `xml:"NewVoipDNSServer1"`
	NewVoipDNSServer2                string   `xml:"NewVoipDNSServer2"`
	NewUpnpControlEnabled            bool     `xml:"NewUpnpControlEnabled"`
	NewRoutedBridgedModeBoth         uint8    `xml:"NewRoutedBridgedModeBoth"`
	NewX_AVM_DE_TotalBytesSent64     string   `xml:"NewX_AVM_DE_TotalBytesSent64"`
	NewX_AVM_DE_TotalBytesReceived64 string   `xml:"NewX_AVM_DE_TotalBytesReceived64"`
	NewX_AVM_DE_WANAccessType        string   `xml:"NewX_AVM_DE_WANAccessType"`
}

func (client *ServiceClient) GetAddonInfos(out *GetAddonInfosResponse) error {
	in := &GetAddonInfosRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetAddonInfosRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetAddonInfosRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAddonInfosResponse]{
		Body: &tr064.SOAPResponseBody[GetAddonInfosResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAddonInfos", soapRequest, soapResponse)
}

type X_AVM_DE_GetDsliteStatusRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetDsliteStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDsliteStatusResponse struct {
	XMLName                  xml.Name `xml:"X_AVM_DE_GetDsliteStatusResponse"`
	NewX_AVM_DE_DsliteStatus bool     `xml:"NewX_AVM_DE_DsliteStatus"`
}

func (client *ServiceClient) X_AVM_DE_GetDsliteStatus(out *X_AVM_DE_GetDsliteStatusResponse) error {
	in := &X_AVM_DE_GetDsliteStatusRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetDsliteStatusRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetDsliteStatusRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetDsliteStatusResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetDsliteStatusResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetDsliteStatus", soapRequest, soapResponse)
}

type X_AVM_DE_GetIPTVInfosRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetIPTVInfosRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPTVInfosResponse struct {
	XMLName                   xml.Name `xml:"X_AVM_DE_GetIPTVInfosResponse"`
	NewX_AVM_DE_IPTV_Enabled  bool     `xml:"NewX_AVM_DE_IPTV_Enabled"`
	NewX_AVM_DE_IPTV_Provider string   `xml:"NewX_AVM_DE_IPTV_Provider"`
	NewX_AVM_DE_IPTV_URL      string   `xml:"NewX_AVM_DE_IPTV_URL"`
}

func (client *ServiceClient) X_AVM_DE_GetIPTVInfos(out *X_AVM_DE_GetIPTVInfosResponse) error {
	in := &X_AVM_DE_GetIPTVInfosRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetIPTVInfosRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetIPTVInfosRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetIPTVInfosResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetIPTVInfosResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetIPTVInfos", soapRequest, soapResponse)
}
