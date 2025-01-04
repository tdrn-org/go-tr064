// generated from spec version: 1.0
package x_wanmobileconn

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
	NewEnabled         bool     `xml:"NewEnabled"`
	NewStatus          string   `xml:"NewStatus"`
	NewPINFailureCount uint16   `xml:"NewPINFailureCount"`
	NewPUKFailureCount uint16   `xml:"NewPUKFailureCount"`
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

type GetInfoExRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoExRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoExResponse struct {
	XMLName                    xml.Name `xml:"GetInfoExResponse"`
	NewSerialNumber            string   `xml:"NewSerialNumber"`
	NewEnableVoIPPDN           bool     `xml:"NewEnableVoIPPDN"`
	NewPPPUsername             string   `xml:"NewPPPUsername"`
	NewPPPUsernameVoIP         string   `xml:"NewPPPUsernameVoIP"`
	NewPPPAuthProtocol         string   `xml:"NewPPPAuthProtocol"`
	NewPPPAuthProtocolVoIP     string   `xml:"NewPPPAuthProtocolVoIP"`
	NewSoftwareVersion         string   `xml:"NewSoftwareVersion"`
	NewUptime                  uint32   `xml:"NewUptime"`
	NewPDN1_MTU                uint32   `xml:"NewPDN1_MTU"`
	NewPDN2_MTU                uint32   `xml:"NewPDN2_MTU"`
	NewIMSI                    string   `xml:"NewIMSI"`
	NewAPN_VoIP                string   `xml:"NewAPN_VoIP"`
	NewAPN                     string   `xml:"NewAPN"`
	NewRoaming                 bool     `xml:"NewRoaming"`
	NewCurrentAccessTechnology string   `xml:"NewCurrentAccessTechnology"`
	NewSignalRSRP0             int32    `xml:"NewSignalRSRP0"`
	NewSignalRSRP1             int32    `xml:"NewSignalRSRP1"`
	NewCellList                string   `xml:"NewCellList"`
}

func (client *ServiceClient) GetInfoEx(out *GetInfoExResponse) error {
	in := &GetInfoExRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoExRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetInfoExRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoExResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoExResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfoEx", soapRequest, soapResponse)
}

type SetPINRequest struct {
	XMLName      xml.Name `xml:"u:SetPINRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewPIN       string   `xml:"NewPIN"`
}

type SetPINResponse struct {
	XMLName xml.Name `xml:"SetPINResponse"`
}

func (client *ServiceClient) SetPIN(in *SetPINRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetPINRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetPINRequest]{
			In: in,
		},
	}
	out := &SetPINResponse{}
	soapResponse := &tr064.SOAPResponse[SetPINResponse]{
		Body: &tr064.SOAPResponseBody[SetPINResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetPIN", soapRequest, soapResponse)
}

type SetPUKRequest struct {
	XMLName      xml.Name `xml:"u:SetPUKRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewPIN       string   `xml:"NewPIN"`
	NewPUK       string   `xml:"NewPUK"`
}

type SetPUKResponse struct {
	XMLName xml.Name `xml:"SetPUKResponse"`
}

func (client *ServiceClient) SetPUK(in *SetPUKRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetPUKRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetPUKRequest]{
			In: in,
		},
	}
	out := &SetPUKResponse{}
	soapResponse := &tr064.SOAPResponse[SetPUKResponse]{
		Body: &tr064.SOAPResponseBody[SetPUKResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetPUK", soapRequest, soapResponse)
}

type SetAccessTechnologyRequest struct {
	XMLName             xml.Name `xml:"u:SetAccessTechnologyRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewAccessTechnology string   `xml:"NewAccessTechnology"`
}

type SetAccessTechnologyResponse struct {
	XMLName xml.Name `xml:"SetAccessTechnologyResponse"`
}

func (client *ServiceClient) SetAccessTechnology(in *SetAccessTechnologyRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAccessTechnologyRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetAccessTechnologyRequest]{
			In: in,
		},
	}
	out := &SetAccessTechnologyResponse{}
	soapResponse := &tr064.SOAPResponse[SetAccessTechnologyResponse]{
		Body: &tr064.SOAPResponseBody[SetAccessTechnologyResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAccessTechnology", soapRequest, soapResponse)
}

type GetAccessTechnologyRequest struct {
	XMLName      xml.Name `xml:"u:GetAccessTechnologyRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAccessTechnologyResponse struct {
	XMLName                     xml.Name `xml:"GetAccessTechnologyResponse"`
	NewAccessTechnology         string   `xml:"NewAccessTechnology"`
	NewPossibleAccessTechnology string   `xml:"NewPossibleAccessTechnology"`
	NewCurrentAccessTechnology  string   `xml:"NewCurrentAccessTechnology"`
}

func (client *ServiceClient) GetAccessTechnology(out *GetAccessTechnologyResponse) error {
	in := &GetAccessTechnologyRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetAccessTechnologyRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetAccessTechnologyRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAccessTechnologyResponse]{
		Body: &tr064.SOAPResponseBody[GetAccessTechnologyResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAccessTechnology", soapRequest, soapResponse)
}

type SetPreferredAccessTechnologyRequest struct {
	XMLName                      xml.Name `xml:"u:SetPreferredAccessTechnologyRequest"`
	XMLNameSpace                 string   `xml:"xmlns:u,attr"`
	NewPreferredAccessTechnology string   `xml:"NewPreferredAccessTechnology"`
}

type SetPreferredAccessTechnologyResponse struct {
	XMLName xml.Name `xml:"SetPreferredAccessTechnologyResponse"`
}

func (client *ServiceClient) SetPreferredAccessTechnology(in *SetPreferredAccessTechnologyRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetPreferredAccessTechnologyRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetPreferredAccessTechnologyRequest]{
			In: in,
		},
	}
	out := &SetPreferredAccessTechnologyResponse{}
	soapResponse := &tr064.SOAPResponse[SetPreferredAccessTechnologyResponse]{
		Body: &tr064.SOAPResponseBody[SetPreferredAccessTechnologyResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetPreferredAccessTechnology", soapRequest, soapResponse)
}

type GetPreferredAccessTechnologyRequest struct {
	XMLName      xml.Name `xml:"u:GetPreferredAccessTechnologyRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetPreferredAccessTechnologyResponse struct {
	XMLName                              xml.Name `xml:"GetPreferredAccessTechnologyResponse"`
	NewPreferredAccessTechnology         string   `xml:"NewPreferredAccessTechnology"`
	NewPossiblePreferredAccessTechnology string   `xml:"NewPossiblePreferredAccessTechnology"`
}

func (client *ServiceClient) GetPreferredAccessTechnology(out *GetPreferredAccessTechnologyResponse) error {
	in := &GetPreferredAccessTechnologyRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetPreferredAccessTechnologyRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetPreferredAccessTechnologyRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetPreferredAccessTechnologyResponse]{
		Body: &tr064.SOAPResponseBody[GetPreferredAccessTechnologyResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetPreferredAccessTechnology", soapRequest, soapResponse)
}

type SetEnabledBandCapabilitiesRequest struct {
	XMLName                  xml.Name `xml:"u:SetEnabledBandCapabilitiesRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewBandCapabilitiesLTE   string   `xml:"NewBandCapabilitiesLTE"`
	NewBandCapabilities5GNSA string   `xml:"NewBandCapabilities5GNSA"`
	NewBandCapabilities5GSA  string   `xml:"NewBandCapabilities5GSA"`
}

type SetEnabledBandCapabilitiesResponse struct {
	XMLName xml.Name `xml:"SetEnabledBandCapabilitiesResponse"`
}

func (client *ServiceClient) SetEnabledBandCapabilities(in *SetEnabledBandCapabilitiesRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetEnabledBandCapabilitiesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetEnabledBandCapabilitiesRequest]{
			In: in,
		},
	}
	out := &SetEnabledBandCapabilitiesResponse{}
	soapResponse := &tr064.SOAPResponse[SetEnabledBandCapabilitiesResponse]{
		Body: &tr064.SOAPResponseBody[SetEnabledBandCapabilitiesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetEnabledBandCapabilities", soapRequest, soapResponse)
}

type GetEnabledBandCapabilitiesRequest struct {
	XMLName      xml.Name `xml:"u:GetEnabledBandCapabilitiesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetEnabledBandCapabilitiesResponse struct {
	XMLName                  xml.Name `xml:"GetEnabledBandCapabilitiesResponse"`
	NewBandCapabilitiesLTE   string   `xml:"NewBandCapabilitiesLTE"`
	NewBandCapabilities5GNSA string   `xml:"NewBandCapabilities5GNSA"`
	NewBandCapabilities5GSA  string   `xml:"NewBandCapabilities5GSA"`
}

func (client *ServiceClient) GetEnabledBandCapabilities(out *GetEnabledBandCapabilitiesResponse) error {
	in := &GetEnabledBandCapabilitiesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetEnabledBandCapabilitiesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetEnabledBandCapabilitiesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetEnabledBandCapabilitiesResponse]{
		Body: &tr064.SOAPResponseBody[GetEnabledBandCapabilitiesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetEnabledBandCapabilities", soapRequest, soapResponse)
}

type GetBandCapabilitiesRequest struct {
	XMLName      xml.Name `xml:"u:GetBandCapabilitiesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetBandCapabilitiesResponse struct {
	XMLName                  xml.Name `xml:"GetBandCapabilitiesResponse"`
	NewBandCapabilitiesLTE   string   `xml:"NewBandCapabilitiesLTE"`
	NewBandCapabilities5GNSA string   `xml:"NewBandCapabilities5GNSA"`
	NewBandCapabilities5GSA  string   `xml:"NewBandCapabilities5GSA"`
}

func (client *ServiceClient) GetBandCapabilities(out *GetBandCapabilitiesResponse) error {
	in := &GetBandCapabilitiesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetBandCapabilitiesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetBandCapabilitiesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetBandCapabilitiesResponse]{
		Body: &tr064.SOAPResponseBody[GetBandCapabilitiesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetBandCapabilities", soapRequest, soapResponse)
}
