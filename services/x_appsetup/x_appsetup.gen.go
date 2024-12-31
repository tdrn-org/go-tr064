// generated from spec version: 1.0
package x_appsetup

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
	XMLName                           xml.Name `xml:"GetInfoResponse"`
	NewMinCharsAppId                  uint16   `xml:"NewMinCharsAppId"`
	NewMaxCharsAppId                  uint16   `xml:"NewMaxCharsAppId"`
	NewAllowedCharsAppId              string   `xml:"NewAllowedCharsAppId"`
	NewMinCharsAppDisplayName         uint16   `xml:"NewMinCharsAppDisplayName"`
	NewMaxCharsAppDisplayName         uint16   `xml:"NewMaxCharsAppDisplayName"`
	NewMinCharsAppUsername            uint16   `xml:"NewMinCharsAppUsername"`
	NewMaxCharsAppUsername            uint16   `xml:"NewMaxCharsAppUsername"`
	NewAllowedCharsAppUsername        string   `xml:"NewAllowedCharsAppUsername"`
	NewMinCharsAppPassword            uint16   `xml:"NewMinCharsAppPassword"`
	NewMaxCharsAppPassword            uint16   `xml:"NewMaxCharsAppPassword"`
	NewAllowedCharsAppPassword        string   `xml:"NewAllowedCharsAppPassword"`
	NewMinCharsIPSecIdentifier        uint16   `xml:"NewMinCharsIPSecIdentifier"`
	NewMaxCharsIPSecIdentifier        uint16   `xml:"NewMaxCharsIPSecIdentifier"`
	NewAllowedCharsIPSecIdentifier    string   `xml:"NewAllowedCharsIPSecIdentifier"`
	NewAllowedCharsCryptAlgos         string   `xml:"NewAllowedCharsCryptAlgos"`
	NewAllowedCharsAppAVMAddress      string   `xml:"NewAllowedCharsAppAVMAddress"`
	NewMinCharsFilter                 uint16   `xml:"NewMinCharsFilter"`
	NewMaxCharsFilter                 uint16   `xml:"NewMaxCharsFilter"`
	NewAllowedCharsFilter             string   `xml:"NewAllowedCharsFilter"`
	NewMinCharsIPSecPreSharedKey      uint16   `xml:"NewMinCharsIPSecPreSharedKey"`
	NewMaxCharsIPSecPreSharedKey      uint16   `xml:"NewMaxCharsIPSecPreSharedKey"`
	NewAllowedCharsIPSecPreSharedKey  string   `xml:"NewAllowedCharsIPSecPreSharedKey"`
	NewMinCharsIPSecXauthUsername     uint16   `xml:"NewMinCharsIPSecXauthUsername"`
	NewMaxCharsIPSecXauthUsername     uint16   `xml:"NewMaxCharsIPSecXauthUsername"`
	NewAllowedCharsIPSecXauthUsername string   `xml:"NewAllowedCharsIPSecXauthUsername"`
	NewMinCharsIPSecXauthPassword     uint16   `xml:"NewMinCharsIPSecXauthPassword"`
	NewMaxCharsIPSecXauthPassword     uint16   `xml:"NewMaxCharsIPSecXauthPassword"`
	NewAllowedCharsIPSecXauthPassword string   `xml:"NewAllowedCharsIPSecXauthPassword"`
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

type GetConfigRequest struct {
	XMLName      xml.Name `xml:"u:GetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetConfigResponse struct {
	XMLName               xml.Name `xml:"GetConfigResponse"`
	NewConfigRight        string   `xml:"NewConfigRight"`
	NewAppRight           string   `xml:"NewAppRight"`
	NewNasRight           string   `xml:"NewNasRight"`
	NewPhoneRight         string   `xml:"NewPhoneRight"`
	NewDialRight          string   `xml:"NewDialRight"`
	NewHomeautoRight      string   `xml:"NewHomeautoRight"`
	NewInternetRights     bool     `xml:"NewInternetRights"`
	NewAccessFromInternet bool     `xml:"NewAccessFromInternet"`
}

func (client *ServiceClient) GetConfig(out *GetConfigResponse) error {
	in := &GetConfigRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetConfigRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetConfigRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetConfigResponse]{
		Body: &tr064.SOAPResponseBody[GetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetConfig", soapRequest, soapResponse)
}

type GetAppMessageFilterRequest struct {
	XMLName      xml.Name `xml:"u:GetAppMessageFilterRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewAppId     string   `xml:"NewAppId"`
}

type GetAppMessageFilterResponse struct {
	XMLName       xml.Name `xml:"GetAppMessageFilterResponse"`
	NewFilterList string   `xml:"NewFilterList"`
}

func (client *ServiceClient) GetAppMessageFilter(in *GetAppMessageFilterRequest, out *GetAppMessageFilterResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetAppMessageFilterRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetAppMessageFilterRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAppMessageFilterResponse]{
		Body: &tr064.SOAPResponseBody[GetAppMessageFilterResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAppMessageFilter", soapRequest, soapResponse)
}

type RegisterAppRequest struct {
	XMLName              xml.Name `xml:"u:RegisterAppRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewAppId             string   `xml:"NewAppId"`
	NewAppDisplayName    string   `xml:"NewAppDisplayName"`
	NewAppDeviceMAC      string   `xml:"NewAppDeviceMAC"`
	NewAppUsername       string   `xml:"NewAppUsername"`
	NewAppPassword       string   `xml:"NewAppPassword"`
	NewAppRight          string   `xml:"NewAppRight"`
	NewNasRight          string   `xml:"NewNasRight"`
	NewPhoneRight        string   `xml:"NewPhoneRight"`
	NewHomeautoRight     string   `xml:"NewHomeautoRight"`
	NewAppInternetRights bool     `xml:"NewAppInternetRights"`
}

type RegisterAppResponse struct {
	XMLName xml.Name `xml:"RegisterAppResponse"`
}

func (client *ServiceClient) RegisterApp(in *RegisterAppRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[RegisterAppRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[RegisterAppRequest]{
			In: in,
		},
	}
	out := &RegisterAppResponse{}
	soapResponse := &tr064.SOAPResponse[RegisterAppResponse]{
		Body: &tr064.SOAPResponseBody[RegisterAppResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "RegisterApp", soapRequest, soapResponse)
}

type SetAppVPNRequest struct {
	XMLName               xml.Name `xml:"u:SetAppVPNRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewAppId              string   `xml:"NewAppId"`
	NewIPSecIdentifier    string   `xml:"NewIPSecIdentifier"`
	NewIPSecPreSharedKey  string   `xml:"NewIPSecPreSharedKey"`
	NewIPSecXauthUsername string   `xml:"NewIPSecXauthUsername"`
	NewIPSecXauthPassword string   `xml:"NewIPSecXauthPassword"`
}

type SetAppVPNResponse struct {
	XMLName xml.Name `xml:"SetAppVPNResponse"`
}

func (client *ServiceClient) SetAppVPN(in *SetAppVPNRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAppVPNRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetAppVPNRequest]{
			In: in,
		},
	}
	out := &SetAppVPNResponse{}
	soapResponse := &tr064.SOAPResponse[SetAppVPNResponse]{
		Body: &tr064.SOAPResponseBody[SetAppVPNResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAppVPN", soapRequest, soapResponse)
}

type SetAppVPNwithPFSRequest struct {
	XMLName               xml.Name `xml:"u:SetAppVPNwithPFSRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewAppId              string   `xml:"NewAppId"`
	NewIPSecIdentifier    string   `xml:"NewIPSecIdentifier"`
	NewIPSecPreSharedKey  string   `xml:"NewIPSecPreSharedKey"`
	NewIPSecXauthUsername string   `xml:"NewIPSecXauthUsername"`
	NewIPSecXauthPassword string   `xml:"NewIPSecXauthPassword"`
}

type SetAppVPNwithPFSResponse struct {
	XMLName xml.Name `xml:"SetAppVPNwithPFSResponse"`
}

func (client *ServiceClient) SetAppVPNwithPFS(in *SetAppVPNwithPFSRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAppVPNwithPFSRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetAppVPNwithPFSRequest]{
			In: in,
		},
	}
	out := &SetAppVPNwithPFSResponse{}
	soapResponse := &tr064.SOAPResponse[SetAppVPNwithPFSResponse]{
		Body: &tr064.SOAPResponseBody[SetAppVPNwithPFSResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAppVPNwithPFS", soapRequest, soapResponse)
}

type SetAppMessageFilterRequest struct {
	XMLName      xml.Name `xml:"u:SetAppMessageFilterRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewAppId     string   `xml:"NewAppId"`
	NewType      string   `xml:"NewType"`
	NewFilter    string   `xml:"NewFilter"`
}

type SetAppMessageFilterResponse struct {
	XMLName xml.Name `xml:"SetAppMessageFilterResponse"`
}

func (client *ServiceClient) SetAppMessageFilter(in *SetAppMessageFilterRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAppMessageFilterRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetAppMessageFilterRequest]{
			In: in,
		},
	}
	out := &SetAppMessageFilterResponse{}
	soapResponse := &tr064.SOAPResponse[SetAppMessageFilterResponse]{
		Body: &tr064.SOAPResponseBody[SetAppMessageFilterResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAppMessageFilter", soapRequest, soapResponse)
}

type SetAppMessageReceiverRequest struct {
	XMLName               xml.Name `xml:"u:SetAppMessageReceiverRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewAppId              string   `xml:"NewAppId"`
	NewCryptAlgos         string   `xml:"NewCryptAlgos"`
	NewAppAVMAddress      string   `xml:"NewAppAVMAddress"`
	NewAppAVMPasswordHash string   `xml:"NewAppAVMPasswordHash"`
}

type SetAppMessageReceiverResponse struct {
	XMLName          xml.Name `xml:"SetAppMessageReceiverResponse"`
	EncryptionSecret string   `xml:"EncryptionSecret"`
	BoxSenderId      string   `xml:"BoxSenderId"`
}

func (client *ServiceClient) SetAppMessageReceiver(in *SetAppMessageReceiverRequest, out *SetAppMessageReceiverResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetAppMessageReceiverRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetAppMessageReceiverRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[SetAppMessageReceiverResponse]{
		Body: &tr064.SOAPResponseBody[SetAppMessageReceiverResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetAppMessageReceiver", soapRequest, soapResponse)
}

type ResetEventRequest struct {
	XMLName      xml.Name `xml:"u:ResetEventRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEventId   uint32   `xml:"NewEventId"`
}

type ResetEventResponse struct {
	XMLName xml.Name `xml:"ResetEventResponse"`
}

func (client *ServiceClient) ResetEvent(in *ResetEventRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[ResetEventRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[ResetEventRequest]{
			In: in,
		},
	}
	out := &ResetEventResponse{}
	soapResponse := &tr064.SOAPResponse[ResetEventResponse]{
		Body: &tr064.SOAPResponseBody[ResetEventResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "ResetEvent", soapRequest, soapResponse)
}

type GetAppRemoteInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetAppRemoteInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAppRemoteInfoResponse struct {
	XMLName                    xml.Name `xml:"GetAppRemoteInfoResponse"`
	NewSubnetMask              string   `xml:"NewSubnetMask"`
	NewIPAddress               string   `xml:"NewIPAddress"`
	NewExternalIPAddress       string   `xml:"NewExternalIPAddress"`
	NewExternalIPv6Address     string   `xml:"NewExternalIPv6Address"`
	NewRemoteAccessDDNSEnabled bool     `xml:"NewRemoteAccessDDNSEnabled"`
	NewRemoteAccessDDNSDomain  string   `xml:"NewRemoteAccessDDNSDomain"`
	NewMyFritzEnabled          bool     `xml:"NewMyFritzEnabled"`
	NewMyFritzDynDNSName       string   `xml:"NewMyFritzDynDNSName"`
}

func (client *ServiceClient) GetAppRemoteInfo(out *GetAppRemoteInfoResponse) error {
	in := &GetAppRemoteInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetAppRemoteInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetAppRemoteInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetAppRemoteInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetAppRemoteInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetAppRemoteInfo", soapRequest, soapResponse)
}

type GetBoxSenderIdRequest struct {
	XMLName      xml.Name `xml:"u:GetBoxSenderIdRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewAppId     string   `xml:"NewAppId"`
}

type GetBoxSenderIdResponse struct {
	XMLName        xml.Name `xml:"GetBoxSenderIdResponse"`
	NewBoxSenderId string   `xml:"NewBoxSenderId"`
}

func (client *ServiceClient) GetBoxSenderId(in *GetBoxSenderIdRequest, out *GetBoxSenderIdResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetBoxSenderIdRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetBoxSenderIdRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetBoxSenderIdResponse]{
		Body: &tr064.SOAPResponseBody[GetBoxSenderIdResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetBoxSenderId", soapRequest, soapResponse)
}
