// generated from spec version: 1.0
package wancommonifconfig

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
	XMLName                                  xml.Name `xml:"GetCommonLinkPropertiesResponse"`
	NewWANAccessType                         string   `xml:"NewWANAccessType"`
	NewLayer1UpstreamMaxBitRate              uint32   `xml:"NewLayer1UpstreamMaxBitRate"`
	NewLayer1DownstreamMaxBitRate            uint32   `xml:"NewLayer1DownstreamMaxBitRate"`
	NewPhysicalLinkStatus                    string   `xml:"NewPhysicalLinkStatus"`
	NewX_AVM_DE_DownstreamCurrentUtilization string   `xml:"NewX_AVM-DE_DownstreamCurrentUtilization"`
	NewX_AVM_DE_UpstreamCurrentUtilization   string   `xml:"NewX_AVM-DE_UpstreamCurrentUtilization"`
	NewX_AVM_DE_DownstreamCurrentMaxSpeed    uint32   `xml:"NewX_AVM-DE_DownstreamCurrentMaxSpeed"`
	NewX_AVM_DE_UpstreamCurrentMaxSpeed      uint32   `xml:"NewX_AVM-DE_UpstreamCurrentMaxSpeed"`
}

func (client *ServiceClient) GetCommonLinkProperties(out *GetCommonLinkPropertiesResponse) error {
	in := &GetCommonLinkPropertiesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetCommonLinkPropertiesRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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

type X_AVM_DE_SetWANAccessTypeRequest struct {
	XMLName       xml.Name `xml:"u:X_AVM-DE_SetWANAccessTypeRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewAccessType string   `xml:"NewAccessType"`
}

type X_AVM_DE_SetWANAccessTypeResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetWANAccessTypeResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetWANAccessType(in *X_AVM_DE_SetWANAccessTypeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetWANAccessTypeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetWANAccessTypeRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetWANAccessTypeResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetWANAccessTypeResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetWANAccessTypeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetWANAccessType", soapRequest, soapResponse)
}

type X_AVM_DE_GetActiveProviderRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetActiveProviderRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetActiveProviderResponse struct {
	XMLName              xml.Name `xml:"X_AVM-DE_GetActiveProviderResponse"`
	NewX_AVM_DE_Provider string   `xml:"NewX_AVM-DE_Provider"`
}

func (client *ServiceClient) X_AVM_DE_GetActiveProvider(out *X_AVM_DE_GetActiveProviderResponse) error {
	in := &X_AVM_DE_GetActiveProviderRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetActiveProviderRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetActiveProviderRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetActiveProviderResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetActiveProviderResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetActiveProvider", soapRequest, soapResponse)
}

type X_AVM_DE_GetOnlineMonitorRequest struct {
	XMLName           xml.Name `xml:"u:X_AVM-DE_GetOnlineMonitorRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewSyncGroupIndex uint32   `xml:"NewSyncGroupIndex"`
}

type X_AVM_DE_GetOnlineMonitorResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetOnlineMonitorResponse"`
	NewTotalNumberSyncGroups uint32   `xml:"NewTotalNumberSyncGroups"`
	NewSyncGroupName         string   `xml:"NewSyncGroupName"`
	NewSyncGroupMode         string   `xml:"NewSyncGroupMode"`
	Newmax_ds                uint32   `xml:"Newmax_ds"`
	Newmax_us                uint32   `xml:"Newmax_us"`
	Newds_current_bps        string   `xml:"Newds_current_bps"`
	Newmc_current_bps        string   `xml:"Newmc_current_bps"`
	Newus_current_bps        string   `xml:"Newus_current_bps"`
	Newprio_realtime_bps     string   `xml:"Newprio_realtime_bps"`
	Newprio_high_bps         string   `xml:"Newprio_high_bps"`
	Newprio_default_bps      string   `xml:"Newprio_default_bps"`
	Newprio_low_bps          string   `xml:"Newprio_low_bps"`
}

func (client *ServiceClient) X_AVM_DE_GetOnlineMonitor(in *X_AVM_DE_GetOnlineMonitorRequest, out *X_AVM_DE_GetOnlineMonitorResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetOnlineMonitorRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetOnlineMonitorRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetOnlineMonitorResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetOnlineMonitorResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetOnlineMonitor", soapRequest, soapResponse)
}
