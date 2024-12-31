// generated from spec version: 1.0
package x_voip

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetInfoExRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoExRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoExResponse struct {
	XMLName                                xml.Name `xml:"GetInfoExResponse"`
	NewVoIPNumberMinChars                  uint16   `xml:"NewVoIPNumberMinChars"`
	NewVoIPNumberMaxChars                  uint16   `xml:"NewVoIPNumberMaxChars"`
	NewVoIPNumberAllowedChars              string   `xml:"NewVoIPNumberAllowedChars"`
	NewVoIPUsernameMinChars                uint16   `xml:"NewVoIPUsernameMinChars"`
	NewVoIPUsernameMaxChars                uint16   `xml:"NewVoIPUsernameMaxChars"`
	NewVoIPUsernameAllowedChars            string   `xml:"NewVoIPUsernameAllowedChars"`
	NewVoIPPasswordMinChars                uint16   `xml:"NewVoIPPasswordMinChars"`
	NewVoIPPasswordMaxChars                uint16   `xml:"NewVoIPPasswordMaxChars"`
	NewVoIPPasswordAllowedChars            string   `xml:"NewVoIPPasswordAllowedChars"`
	NewVoIPRegistrarMinChars               uint16   `xml:"NewVoIPRegistrarMinChars"`
	NewVoIPRegistrarMaxChars               uint16   `xml:"NewVoIPRegistrarMaxChars"`
	NewVoIPRegistrarAllowedChars           string   `xml:"NewVoIPRegistrarAllowedChars"`
	NewVoIPSTUNServerMinChars              uint16   `xml:"NewVoIPSTUNServerMinChars"`
	NewVoIPSTUNServerMaxChars              uint16   `xml:"NewVoIPSTUNServerMaxChars"`
	NewVoIPSTUNServerAllowedChars          string   `xml:"NewVoIPSTUNServerAllowedChars"`
	NewX_AVM_DE_ClientUsernameMinChars     uint16   `xml:"NewX_AVM-DE_ClientUsernameMinChars"`
	NewX_AVM_DE_ClientUsernameMaxChars     uint16   `xml:"NewX_AVM-DE_ClientUsernameMaxChars"`
	NewX_AVM_DE_ClientUsernameAllowedChars string   `xml:"NewX_AVM-DE_ClientUsernameAllowedChars"`
	NewX_AVM_DE_ClientPasswordMinChars     uint16   `xml:"NewX_AVM-DE_ClientPasswordMinChars"`
	NewX_AVM_DE_ClientPasswordMaxChars     uint16   `xml:"NewX_AVM-DE_ClientPasswordMaxChars"`
	NewX_AVM_DE_ClientPasswordAllowedChars string   `xml:"NewX_AVM-DE_ClientPasswordAllowedChars"`
}

func (client *ServiceClient) GetInfoEx(out *GetInfoExResponse) error {
	in := &GetInfoExRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoExRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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

type X_AVM_DE_AddVoIPAccountRequest struct {
	XMLName              xml.Name `xml:"u:X_AVM-DE_AddVoIPAccountRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex  uint16   `xml:"NewVoIPAccountIndex"`
	NewVoIPRegistrar     string   `xml:"NewVoIPRegistrar"`
	NewVoIPNumber        string   `xml:"NewVoIPNumber"`
	NewVoIPUsername      string   `xml:"NewVoIPUsername"`
	NewVoIPPassword      string   `xml:"NewVoIPPassword"`
	NewVoIPOutboundProxy string   `xml:"NewVoIPOutboundProxy"`
	NewVoIPSTUNServer    string   `xml:"NewVoIPSTUNServer"`
}

type X_AVM_DE_AddVoIPAccountResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_AddVoIPAccountResponse"`
}

func (client *ServiceClient) X_AVM_DE_AddVoIPAccount(in *X_AVM_DE_AddVoIPAccountRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_AddVoIPAccountRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_AddVoIPAccountRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_AddVoIPAccountResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_AddVoIPAccountResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_AddVoIPAccountResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_AddVoIPAccount", soapRequest, soapResponse)
}

type X_AVM_DE_GetVoIPAccountRequest struct {
	XMLName             xml.Name `xml:"u:X_AVM-DE_GetVoIPAccountRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex uint16   `xml:"NewVoIPAccountIndex"`
}

type X_AVM_DE_GetVoIPAccountResponse struct {
	XMLName                xml.Name `xml:"X_AVM-DE_GetVoIPAccountResponse"`
	NewVoIPRegistrar       string   `xml:"NewVoIPRegistrar"`
	NewVoIPNumber          string   `xml:"NewVoIPNumber"`
	NewVoIPUsername        string   `xml:"NewVoIPUsername"`
	NewVoIPPassword        string   `xml:"NewVoIPPassword"`
	NewVoIPOutboundProxy   string   `xml:"NewVoIPOutboundProxy"`
	NewVoIPSTUNServer      string   `xml:"NewVoIPSTUNServer"`
	NewX_AVM_DE_VoIPStatus string   `xml:"NewX_AVM-DE_VoIPStatus"`
}

func (client *ServiceClient) X_AVM_DE_GetVoIPAccount(in *X_AVM_DE_GetVoIPAccountRequest, out *X_AVM_DE_GetVoIPAccountResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetVoIPAccountRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetVoIPAccountRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetVoIPAccountResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetVoIPAccountResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPAccount", soapRequest, soapResponse)
}

type X_AVM_DE_DelVoIPAccountRequest struct {
	XMLName             xml.Name `xml:"u:X_AVM-DE_DelVoIPAccountRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex uint16   `xml:"NewVoIPAccountIndex"`
}

type X_AVM_DE_DelVoIPAccountResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DelVoIPAccountResponse"`
}

func (client *ServiceClient) X_AVM_DE_DelVoIPAccount(in *X_AVM_DE_DelVoIPAccountRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DelVoIPAccountRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DelVoIPAccountRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DelVoIPAccountResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DelVoIPAccountResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DelVoIPAccountResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DelVoIPAccount", soapRequest, soapResponse)
}

type X_AVM_DE_GetVoIPAccountsRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetVoIPAccountsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetVoIPAccountsResponse struct {
	XMLName                     xml.Name `xml:"X_AVM-DE_GetVoIPAccountsResponse"`
	NewX_AVM_DE_VoIPAccountList string   `xml:"NewX_AVM-DE_VoIPAccountList"`
}

func (client *ServiceClient) X_AVM_DE_GetVoIPAccounts(out *X_AVM_DE_GetVoIPAccountsResponse) error {
	in := &X_AVM_DE_GetVoIPAccountsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetVoIPAccountsRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetVoIPAccountsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetVoIPAccountsResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetVoIPAccountsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPAccounts", soapRequest, soapResponse)
}

type X_AVM_DE_GetVoIPStatusRequest struct {
	XMLName             xml.Name `xml:"u:X_AVM-DE_GetVoIPStatusRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex uint16   `xml:"NewVoIPAccountIndex"`
}

type X_AVM_DE_GetVoIPStatusResponse struct {
	XMLName                xml.Name `xml:"X_AVM-DE_GetVoIPStatusResponse"`
	NewX_AVM_DE_VoIPStatus string   `xml:"NewX_AVM-DE_VoIPStatus"`
}

func (client *ServiceClient) X_AVM_DE_GetVoIPStatus(in *X_AVM_DE_GetVoIPStatusRequest, out *X_AVM_DE_GetVoIPStatusResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetVoIPStatusRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetVoIPStatusRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetVoIPStatusResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetVoIPStatusResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPStatus", soapRequest, soapResponse)
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName         xml.Name `xml:"GetInfoResponse"`
	NewFaxT38Enable bool     `xml:"NewFaxT38Enable"`
	NewVoiceCoding  string   `xml:"NewVoiceCoding"`
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

type SetConfigRequest struct {
	XMLName         xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewFaxT38Enable bool     `xml:"NewFaxT38Enable"`
	NewVoiceCoding  string   `xml:"NewVoiceCoding"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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

type GetMaxVoIPNumbersRequest struct {
	XMLName      xml.Name `xml:"u:GetMaxVoIPNumbersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetMaxVoIPNumbersResponse struct {
	XMLName           xml.Name `xml:"GetMaxVoIPNumbersResponse"`
	NewMaxVoIPNumbers uint16   `xml:"NewMaxVoIPNumbers"`
}

func (client *ServiceClient) GetMaxVoIPNumbers(out *GetMaxVoIPNumbersResponse) error {
	in := &GetMaxVoIPNumbersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetMaxVoIPNumbersRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetMaxVoIPNumbersRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetMaxVoIPNumbersResponse]{
		Body: &tr064.SOAPResponseBody[GetMaxVoIPNumbersResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetMaxVoIPNumbers", soapRequest, soapResponse)
}

type GetExistingVoIPNumbersRequest struct {
	XMLName      xml.Name `xml:"u:GetExistingVoIPNumbersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetExistingVoIPNumbersResponse struct {
	XMLName                xml.Name `xml:"GetExistingVoIPNumbersResponse"`
	NewExistingVoIPNumbers uint16   `xml:"NewExistingVoIPNumbers"`
}

func (client *ServiceClient) GetExistingVoIPNumbers(out *GetExistingVoIPNumbersResponse) error {
	in := &GetExistingVoIPNumbersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetExistingVoIPNumbersRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetExistingVoIPNumbersRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetExistingVoIPNumbersResponse]{
		Body: &tr064.SOAPResponseBody[GetExistingVoIPNumbersResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetExistingVoIPNumbers", soapRequest, soapResponse)
}

type X_AVM_DE_GetNumberOfClientsRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetNumberOfClientsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetNumberOfClientsResponse struct {
	XMLName                     xml.Name `xml:"X_AVM-DE_GetNumberOfClientsResponse"`
	NewX_AVM_DE_NumberOfClients uint16   `xml:"NewX_AVM-DE_NumberOfClients"`
}

func (client *ServiceClient) X_AVM_DE_GetNumberOfClients(out *X_AVM_DE_GetNumberOfClientsResponse) error {
	in := &X_AVM_DE_GetNumberOfClientsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetNumberOfClientsRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetNumberOfClientsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetNumberOfClientsResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetNumberOfClientsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfClients", soapRequest, soapResponse)
}

type X_AVM_DE_GetClientRequest struct {
	XMLName                 xml.Name `xml:"u:X_AVM-DE_GetClientRequest"`
	XMLNameSpace            string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex uint16   `xml:"NewX_AVM-DE_ClientIndex"`
}

type X_AVM_DE_GetClientResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetClientResponse"`
	NewX_AVM_DE_ClientUsername      string   `xml:"NewX_AVM-DE_ClientUsername"`
	NewX_AVM_DE_ClientRegistrar     string   `xml:"NewX_AVM-DE_ClientRegistrar"`
	NewX_AVM_DE_ClientRegistrarPort uint16   `xml:"NewX_AVM-DE_ClientRegistrarPort"`
	NewX_AVM_DE_PhoneName           string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_OutGoingNumber      string   `xml:"NewX_AVM-DE_OutGoingNumber"`
}

func (client *ServiceClient) X_AVM_DE_GetClient(in *X_AVM_DE_GetClientRequest, out *X_AVM_DE_GetClientResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetClientRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetClientRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetClientResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetClientResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient", soapRequest, soapResponse)
}

type X_AVM_DE_GetClient2Request struct {
	XMLName                 xml.Name `xml:"u:X_AVM-DE_GetClient2Request"`
	XMLNameSpace            string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex uint16   `xml:"NewX_AVM-DE_ClientIndex"`
}

type X_AVM_DE_GetClient2Response struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetClient2Response"`
	NewX_AVM_DE_ClientUsername      string   `xml:"NewX_AVM-DE_ClientUsername"`
	NewX_AVM_DE_ClientRegistrar     string   `xml:"NewX_AVM-DE_ClientRegistrar"`
	NewX_AVM_DE_ClientRegistrarPort uint16   `xml:"NewX_AVM-DE_ClientRegistrarPort"`
	NewX_AVM_DE_PhoneName           string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_ClientId            string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_OutGoingNumber      string   `xml:"NewX_AVM-DE_OutGoingNumber"`
	NewX_AVM_DE_InternalNumber      string   `xml:"NewX_AVM-DE_InternalNumber"`
}

func (client *ServiceClient) X_AVM_DE_GetClient2(in *X_AVM_DE_GetClient2Request, out *X_AVM_DE_GetClient2Response) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetClient2Request]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetClient2Request]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetClient2Response]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetClient2Response]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient2", soapRequest, soapResponse)
}

type X_AVM_DE_SetClientRequest struct {
	XMLName                    xml.Name `xml:"u:X_AVM-DE_SetClientRequest"`
	XMLNameSpace               string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex    uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_ClientPassword string   `xml:"NewX_AVM-DE_ClientPassword"`
	NewX_AVM_DE_PhoneName      string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_OutGoingNumber string   `xml:"NewX_AVM-DE_OutGoingNumber"`
}

type X_AVM_DE_SetClientResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetClientResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetClient(in *X_AVM_DE_SetClientRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetClientRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetClientRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetClientResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetClientResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetClientResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient", soapRequest, soapResponse)
}

type X_AVM_DE_SetClient2Request struct {
	XMLName                    xml.Name `xml:"u:X_AVM-DE_SetClient2Request"`
	XMLNameSpace               string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex    uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_ClientPassword string   `xml:"NewX_AVM-DE_ClientPassword"`
	NewX_AVM_DE_ClientId       string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_PhoneName      string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_OutGoingNumber string   `xml:"NewX_AVM-DE_OutGoingNumber"`
}

type X_AVM_DE_SetClient2Response struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetClient2Response"`
}

func (client *ServiceClient) X_AVM_DE_SetClient2(in *X_AVM_DE_SetClient2Request) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetClient2Request]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetClient2Request]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetClient2Response{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetClient2Response]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetClient2Response]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient2", soapRequest, soapResponse)
}

type X_AVM_DE_GetClient3Request struct {
	XMLName                 xml.Name `xml:"u:X_AVM-DE_GetClient3Request"`
	XMLNameSpace            string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex uint16   `xml:"NewX_AVM-DE_ClientIndex"`
}

type X_AVM_DE_GetClient3Response struct {
	XMLName                             xml.Name `xml:"X_AVM-DE_GetClient3Response"`
	NewX_AVM_DE_ClientUsername          string   `xml:"NewX_AVM-DE_ClientUsername"`
	NewX_AVM_DE_ClientRegistrar         string   `xml:"NewX_AVM-DE_ClientRegistrar"`
	NewX_AVM_DE_ClientRegistrarPort     uint16   `xml:"NewX_AVM-DE_ClientRegistrarPort"`
	NewX_AVM_DE_PhoneName               string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_ClientId                string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_OutGoingNumber          string   `xml:"NewX_AVM-DE_OutGoingNumber"`
	NewX_AVM_DE_InComingNumbers         string   `xml:"NewX_AVM-DE_InComingNumbers"`
	NewX_AVM_DE_ExternalRegistration    bool     `xml:"NewX_AVM-DE_ExternalRegistration"`
	NewX_AVM_DE_InternalNumber          string   `xml:"NewX_AVM-DE_InternalNumber"`
	NewX_AVM_DE_DelayedCallNotification bool     `xml:"NewX_AVM-DE_DelayedCallNotification"`
}

func (client *ServiceClient) X_AVM_DE_GetClient3(in *X_AVM_DE_GetClient3Request, out *X_AVM_DE_GetClient3Response) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetClient3Request]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetClient3Request]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetClient3Response]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetClient3Response]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient3", soapRequest, soapResponse)
}

type X_AVM_DE_GetClientByClientIdRequest struct {
	XMLName              xml.Name `xml:"u:X_AVM-DE_GetClientByClientIdRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientId string   `xml:"NewX_AVM-DE_ClientId"`
}

type X_AVM_DE_GetClientByClientIdResponse struct {
	XMLName                             xml.Name `xml:"X_AVM-DE_GetClientByClientIdResponse"`
	NewX_AVM_DE_ClientId                string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_ClientIndex             uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_ClientUsername          string   `xml:"NewX_AVM-DE_ClientUsername"`
	NewX_AVM_DE_ClientRegistrar         string   `xml:"NewX_AVM-DE_ClientRegistrar"`
	NewX_AVM_DE_ClientRegistrarPort     uint16   `xml:"NewX_AVM-DE_ClientRegistrarPort"`
	NewX_AVM_DE_PhoneName               string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_OutGoingNumber          string   `xml:"NewX_AVM-DE_OutGoingNumber"`
	NewX_AVM_DE_InComingNumbers         string   `xml:"NewX_AVM-DE_InComingNumbers"`
	NewX_AVM_DE_ExternalRegistration    bool     `xml:"NewX_AVM-DE_ExternalRegistration"`
	NewX_AVM_DE_InternalNumber          string   `xml:"NewX_AVM-DE_InternalNumber"`
	NewX_AVM_DE_DelayedCallNotification bool     `xml:"NewX_AVM-DE_DelayedCallNotification"`
}

func (client *ServiceClient) X_AVM_DE_GetClientByClientId(in *X_AVM_DE_GetClientByClientIdRequest, out *X_AVM_DE_GetClientByClientIdResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetClientByClientIdRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetClientByClientIdRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetClientByClientIdResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetClientByClientIdResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClientByClientId", soapRequest, soapResponse)
}

type X_AVM_DE_SetClient3Request struct {
	XMLName                          xml.Name `xml:"u:X_AVM-DE_SetClient3Request"`
	XMLNameSpace                     string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex          uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_ClientPassword       string   `xml:"NewX_AVM-DE_ClientPassword"`
	NewX_AVM_DE_ClientId             string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_PhoneName            string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_OutGoingNumber       string   `xml:"NewX_AVM-DE_OutGoingNumber"`
	NewX_AVM_DE_InComingNumbers      string   `xml:"NewX_AVM-DE_InComingNumbers"`
	NewX_AVM_DE_ExternalRegistration bool     `xml:"NewX_AVM-DE_ExternalRegistration"`
}

type X_AVM_DE_SetClient3Response struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetClient3Response"`
}

func (client *ServiceClient) X_AVM_DE_SetClient3(in *X_AVM_DE_SetClient3Request) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetClient3Request]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetClient3Request]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetClient3Response{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetClient3Response]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetClient3Response]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient3", soapRequest, soapResponse)
}

type X_AVM_DE_SetClient4Request struct {
	XMLName                     xml.Name `xml:"u:X_AVM-DE_SetClient4Request"`
	XMLNameSpace                string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex     uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_ClientPassword  string   `xml:"NewX_AVM-DE_ClientPassword"`
	NewX_AVM_DE_ClientUsername  string   `xml:"NewX_AVM-DE_ClientUsername"`
	NewX_AVM_DE_PhoneName       string   `xml:"NewX_AVM-DE_PhoneName"`
	NewX_AVM_DE_ClientId        string   `xml:"NewX_AVM-DE_ClientId"`
	NewX_AVM_DE_OutGoingNumber  string   `xml:"NewX_AVM-DE_OutGoingNumber"`
	NewX_AVM_DE_InComingNumbers string   `xml:"NewX_AVM-DE_InComingNumbers"`
}

type X_AVM_DE_SetClient4Response struct {
	XMLName                    xml.Name `xml:"X_AVM-DE_SetClient4Response"`
	NewX_AVM_DE_InternalNumber string   `xml:"NewX_AVM-DE_InternalNumber"`
}

func (client *ServiceClient) X_AVM_DE_SetClient4(in *X_AVM_DE_SetClient4Request, out *X_AVM_DE_SetClient4Response) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetClient4Request]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetClient4Request]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetClient4Response]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetClient4Response]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient4", soapRequest, soapResponse)
}

type X_AVM_DE_GetClientsRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetClientsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetClientsResponse struct {
	XMLName                xml.Name `xml:"X_AVM-DE_GetClientsResponse"`
	NewX_AVM_DE_ClientList string   `xml:"NewX_AVM-DE_ClientList"`
}

func (client *ServiceClient) X_AVM_DE_GetClients(out *X_AVM_DE_GetClientsResponse) error {
	in := &X_AVM_DE_GetClientsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetClientsRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetClientsRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetClientsResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetClientsResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClients", soapRequest, soapResponse)
}

type X_AVM_DE_GetNumberOfNumbersRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetNumberOfNumbersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetNumberOfNumbersResponse struct {
	XMLName            xml.Name `xml:"X_AVM-DE_GetNumberOfNumbersResponse"`
	NewNumberOfNumbers uint32   `xml:"NewNumberOfNumbers"`
}

func (client *ServiceClient) X_AVM_DE_GetNumberOfNumbers(out *X_AVM_DE_GetNumberOfNumbersResponse) error {
	in := &X_AVM_DE_GetNumberOfNumbersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetNumberOfNumbersRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetNumberOfNumbersRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetNumberOfNumbersResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetNumberOfNumbersResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfNumbers", soapRequest, soapResponse)
}

type X_AVM_DE_GetNumbersRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetNumbersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetNumbersResponse struct {
	XMLName       xml.Name `xml:"X_AVM-DE_GetNumbersResponse"`
	NewNumberList string   `xml:"NewNumberList"`
}

func (client *ServiceClient) X_AVM_DE_GetNumbers(out *X_AVM_DE_GetNumbersResponse) error {
	in := &X_AVM_DE_GetNumbersRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetNumbersRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetNumbersRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetNumbersResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetNumbersResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumbers", soapRequest, soapResponse)
}

type X_AVM_DE_DeleteClientRequest struct {
	XMLName                 xml.Name `xml:"u:X_AVM-DE_DeleteClientRequest"`
	XMLNameSpace            string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex uint16   `xml:"NewX_AVM-DE_ClientIndex"`
}

type X_AVM_DE_DeleteClientResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DeleteClientResponse"`
}

func (client *ServiceClient) X_AVM_DE_DeleteClient(in *X_AVM_DE_DeleteClientRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DeleteClientRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DeleteClientRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DeleteClientResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DeleteClientResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DeleteClientResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DeleteClient", soapRequest, soapResponse)
}

type X_AVM_DE_DialGetConfigRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_DialGetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_DialGetConfigResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_DialGetConfigResponse"`
	NewX_AVM_DE_PhoneName string   `xml:"NewX_AVM-DE_PhoneName"`
}

func (client *ServiceClient) X_AVM_DE_DialGetConfig(out *X_AVM_DE_DialGetConfigResponse) error {
	in := &X_AVM_DE_DialGetConfigRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DialGetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DialGetConfigRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DialGetConfigResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DialGetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialGetConfig", soapRequest, soapResponse)
}

type X_AVM_DE_DialSetConfigRequest struct {
	XMLName               xml.Name `xml:"u:X_AVM-DE_DialSetConfigRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_PhoneName string   `xml:"NewX_AVM-DE_PhoneName"`
}

type X_AVM_DE_DialSetConfigResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DialSetConfigResponse"`
}

func (client *ServiceClient) X_AVM_DE_DialSetConfig(in *X_AVM_DE_DialSetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DialSetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DialSetConfigRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DialSetConfigResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DialSetConfigResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DialSetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialSetConfig", soapRequest, soapResponse)
}

type X_AVM_DE_DialNumberRequest struct {
	XMLName                 xml.Name `xml:"u:X_AVM-DE_DialNumberRequest"`
	XMLNameSpace            string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_PhoneNumber string   `xml:"NewX_AVM-DE_PhoneNumber"`
}

type X_AVM_DE_DialNumberResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DialNumberResponse"`
}

func (client *ServiceClient) X_AVM_DE_DialNumber(in *X_AVM_DE_DialNumberRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DialNumberRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DialNumberRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DialNumberResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DialNumberResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DialNumberResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialNumber", soapRequest, soapResponse)
}

type X_AVM_DE_DialHangupRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_DialHangupRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_DialHangupResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DialHangupResponse"`
}

func (client *ServiceClient) X_AVM_DE_DialHangup() error {
	in := &X_AVM_DE_DialHangupRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DialHangupRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DialHangupRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DialHangupResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DialHangupResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DialHangupResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialHangup", soapRequest, soapResponse)
}

type X_AVM_DE_GetPhonePortRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetPhonePortRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type X_AVM_DE_GetPhonePortResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_GetPhonePortResponse"`
	NewX_AVM_DE_PhoneName string   `xml:"NewX_AVM-DE_PhoneName"`
}

func (client *ServiceClient) X_AVM_DE_GetPhonePort(in *X_AVM_DE_GetPhonePortRequest, out *X_AVM_DE_GetPhonePortResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetPhonePortRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetPhonePortRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetPhonePortResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetPhonePortResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetPhonePort", soapRequest, soapResponse)
}

type X_AVM_DE_SetDelayedCallNotificationRequest struct {
	XMLName                             xml.Name `xml:"u:X_AVM-DE_SetDelayedCallNotificationRequest"`
	XMLNameSpace                        string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_ClientIndex             uint16   `xml:"NewX_AVM-DE_ClientIndex"`
	NewX_AVM_DE_DelayedCallNotification bool     `xml:"NewX_AVM-DE_DelayedCallNotification"`
}

type X_AVM_DE_SetDelayedCallNotificationResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetDelayedCallNotificationResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetDelayedCallNotification(in *X_AVM_DE_SetDelayedCallNotificationRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetDelayedCallNotificationRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetDelayedCallNotificationRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetDelayedCallNotificationResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetDelayedCallNotificationResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetDelayedCallNotificationResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetDelayedCallNotification", soapRequest, soapResponse)
}

type GetVoIPCommonCountryCodeRequest struct {
	XMLName      xml.Name `xml:"u:GetVoIPCommonCountryCodeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetVoIPCommonCountryCodeResponse struct {
	XMLName            xml.Name `xml:"GetVoIPCommonCountryCodeResponse"`
	NewVoIPCountryCode string   `xml:"NewVoIPCountryCode"`
}

func (client *ServiceClient) GetVoIPCommonCountryCode(out *GetVoIPCommonCountryCodeResponse) error {
	in := &GetVoIPCommonCountryCodeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetVoIPCommonCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetVoIPCommonCountryCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetVoIPCommonCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[GetVoIPCommonCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetVoIPCommonCountryCode", soapRequest, soapResponse)
}

type X_AVM_DE_GetVoIPCommonCountryCodeRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetVoIPCommonCountryCodeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetVoIPCommonCountryCodeResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_GetVoIPCommonCountryCodeResponse"`
	NewX_AVM_DE_LKZ       string   `xml:"NewX_AVM-DE_LKZ"`
	NewX_AVM_DE_LKZPrefix string   `xml:"NewX_AVM-DE_LKZPrefix"`
}

func (client *ServiceClient) X_AVM_DE_GetVoIPCommonCountryCode(out *X_AVM_DE_GetVoIPCommonCountryCodeResponse) error {
	in := &X_AVM_DE_GetVoIPCommonCountryCodeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetVoIPCommonCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetVoIPCommonCountryCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetVoIPCommonCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetVoIPCommonCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPCommonCountryCode", soapRequest, soapResponse)
}

type SetVoIPCommonCountryCodeRequest struct {
	XMLName            xml.Name `xml:"u:SetVoIPCommonCountryCodeRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewVoIPCountryCode string   `xml:"NewVoIPCountryCode"`
}

type SetVoIPCommonCountryCodeResponse struct {
	XMLName xml.Name `xml:"SetVoIPCommonCountryCodeResponse"`
}

func (client *ServiceClient) SetVoIPCommonCountryCode(in *SetVoIPCommonCountryCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetVoIPCommonCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetVoIPCommonCountryCodeRequest]{
			In: in,
		},
	}
	out := &SetVoIPCommonCountryCodeResponse{}
	soapResponse := &tr064.SOAPResponse[SetVoIPCommonCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[SetVoIPCommonCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPCommonCountryCode", soapRequest, soapResponse)
}

type X_AVM_DE_SetVoIPCommonCountryCodeRequest struct {
	XMLName               xml.Name `xml:"u:X_AVM-DE_SetVoIPCommonCountryCodeRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_LKZ       string   `xml:"NewX_AVM-DE_LKZ"`
	NewX_AVM_DE_LKZPrefix string   `xml:"NewX_AVM-DE_LKZPrefix"`
}

type X_AVM_DE_SetVoIPCommonCountryCodeResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetVoIPCommonCountryCodeResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetVoIPCommonCountryCode(in *X_AVM_DE_SetVoIPCommonCountryCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetVoIPCommonCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetVoIPCommonCountryCodeRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetVoIPCommonCountryCodeResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetVoIPCommonCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetVoIPCommonCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetVoIPCommonCountryCode", soapRequest, soapResponse)
}

type GetVoIPEnableCountryCodeRequest struct {
	XMLName             xml.Name `xml:"u:GetVoIPEnableCountryCodeRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex uint16   `xml:"NewVoIPAccountIndex"`
}

type GetVoIPEnableCountryCodeResponse struct {
	XMLName                  xml.Name `xml:"GetVoIPEnableCountryCodeResponse"`
	NewVoIPEnableCountryCode bool     `xml:"NewVoIPEnableCountryCode"`
}

func (client *ServiceClient) GetVoIPEnableCountryCode(in *GetVoIPEnableCountryCodeRequest, out *GetVoIPEnableCountryCodeResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetVoIPEnableCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetVoIPEnableCountryCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetVoIPEnableCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[GetVoIPEnableCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetVoIPEnableCountryCode", soapRequest, soapResponse)
}

type SetVoIPEnableCountryCodeRequest struct {
	XMLName                  xml.Name `xml:"u:SetVoIPEnableCountryCodeRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex      uint16   `xml:"NewVoIPAccountIndex"`
	NewVoIPEnableCountryCode bool     `xml:"NewVoIPEnableCountryCode"`
}

type SetVoIPEnableCountryCodeResponse struct {
	XMLName xml.Name `xml:"SetVoIPEnableCountryCodeResponse"`
}

func (client *ServiceClient) SetVoIPEnableCountryCode(in *SetVoIPEnableCountryCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetVoIPEnableCountryCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetVoIPEnableCountryCodeRequest]{
			In: in,
		},
	}
	out := &SetVoIPEnableCountryCodeResponse{}
	soapResponse := &tr064.SOAPResponse[SetVoIPEnableCountryCodeResponse]{
		Body: &tr064.SOAPResponseBody[SetVoIPEnableCountryCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPEnableCountryCode", soapRequest, soapResponse)
}

type GetVoIPCommonAreaCodeRequest struct {
	XMLName      xml.Name `xml:"u:GetVoIPCommonAreaCodeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetVoIPCommonAreaCodeResponse struct {
	XMLName         xml.Name `xml:"GetVoIPCommonAreaCodeResponse"`
	NewVoIPAreaCode string   `xml:"NewVoIPAreaCode"`
}

func (client *ServiceClient) GetVoIPCommonAreaCode(out *GetVoIPCommonAreaCodeResponse) error {
	in := &GetVoIPCommonAreaCodeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetVoIPCommonAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetVoIPCommonAreaCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetVoIPCommonAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[GetVoIPCommonAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetVoIPCommonAreaCode", soapRequest, soapResponse)
}

type X_AVM_DE_GetVoIPCommonAreaCodeRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetVoIPCommonAreaCodeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetVoIPCommonAreaCodeResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_GetVoIPCommonAreaCodeResponse"`
	NewX_AVM_DE_OKZ       string   `xml:"NewX_AVM-DE_OKZ"`
	NewX_AVM_DE_OKZPrefix string   `xml:"NewX_AVM-DE_OKZPrefix"`
}

func (client *ServiceClient) X_AVM_DE_GetVoIPCommonAreaCode(out *X_AVM_DE_GetVoIPCommonAreaCodeResponse) error {
	in := &X_AVM_DE_GetVoIPCommonAreaCodeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetVoIPCommonAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetVoIPCommonAreaCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetVoIPCommonAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetVoIPCommonAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPCommonAreaCode", soapRequest, soapResponse)
}

type SetVoIPCommonAreaCodeRequest struct {
	XMLName         xml.Name `xml:"u:SetVoIPCommonAreaCodeRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewVoIPAreaCode string   `xml:"NewVoIPAreaCode"`
}

type SetVoIPCommonAreaCodeResponse struct {
	XMLName xml.Name `xml:"SetVoIPCommonAreaCodeResponse"`
}

func (client *ServiceClient) SetVoIPCommonAreaCode(in *SetVoIPCommonAreaCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetVoIPCommonAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetVoIPCommonAreaCodeRequest]{
			In: in,
		},
	}
	out := &SetVoIPCommonAreaCodeResponse{}
	soapResponse := &tr064.SOAPResponse[SetVoIPCommonAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[SetVoIPCommonAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPCommonAreaCode", soapRequest, soapResponse)
}

type X_AVM_DE_SetVoIPCommonAreaCodeRequest struct {
	XMLName               xml.Name `xml:"u:X_AVM-DE_SetVoIPCommonAreaCodeRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_OKZ       string   `xml:"NewX_AVM-DE_OKZ"`
	NewX_AVM_DE_OKZPrefix string   `xml:"NewX_AVM-DE_OKZPrefix"`
}

type X_AVM_DE_SetVoIPCommonAreaCodeResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetVoIPCommonAreaCodeResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetVoIPCommonAreaCode(in *X_AVM_DE_SetVoIPCommonAreaCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetVoIPCommonAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetVoIPCommonAreaCodeRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetVoIPCommonAreaCodeResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetVoIPCommonAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetVoIPCommonAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetVoIPCommonAreaCode", soapRequest, soapResponse)
}

type GetVoIPEnableAreaCodeRequest struct {
	XMLName             xml.Name `xml:"u:GetVoIPEnableAreaCodeRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex uint16   `xml:"NewVoIPAccountIndex"`
}

type GetVoIPEnableAreaCodeResponse struct {
	XMLName               xml.Name `xml:"GetVoIPEnableAreaCodeResponse"`
	NewVoIPEnableAreaCode bool     `xml:"NewVoIPEnableAreaCode"`
}

func (client *ServiceClient) GetVoIPEnableAreaCode(in *GetVoIPEnableAreaCodeRequest, out *GetVoIPEnableAreaCodeResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetVoIPEnableAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetVoIPEnableAreaCodeRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetVoIPEnableAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[GetVoIPEnableAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetVoIPEnableAreaCode", soapRequest, soapResponse)
}

type SetVoIPEnableAreaCodeRequest struct {
	XMLName               xml.Name `xml:"u:SetVoIPEnableAreaCodeRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewVoIPAccountIndex   uint16   `xml:"NewVoIPAccountIndex"`
	NewVoIPEnableAreaCode bool     `xml:"NewVoIPEnableAreaCode"`
}

type SetVoIPEnableAreaCodeResponse struct {
	XMLName xml.Name `xml:"SetVoIPEnableAreaCodeResponse"`
}

func (client *ServiceClient) SetVoIPEnableAreaCode(in *SetVoIPEnableAreaCodeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetVoIPEnableAreaCodeRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[SetVoIPEnableAreaCodeRequest]{
			In: in,
		},
	}
	out := &SetVoIPEnableAreaCodeResponse{}
	soapResponse := &tr064.SOAPResponse[SetVoIPEnableAreaCodeResponse]{
		Body: &tr064.SOAPResponseBody[SetVoIPEnableAreaCodeResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPEnableAreaCode", soapRequest, soapResponse)
}

type X_AVM_DE_GetAlarmClockRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetAlarmClockRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type X_AVM_DE_GetAlarmClockResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetAlarmClockResponse"`
	NewX_AVM_DE_AlarmClockEnable    bool     `xml:"NewX_AVM-DE_AlarmClockEnable"`
	NewX_AVM_DE_AlarmClockName      string   `xml:"NewX_AVM-DE_AlarmClockName"`
	NewX_AVM_DE_AlarmClockTime      string   `xml:"NewX_AVM-DE_AlarmClockTime"`
	NewX_AVM_DE_AlarmClockWeekdays  string   `xml:"NewX_AVM-DE_AlarmClockWeekdays"`
	NewX_AVM_DE_AlarmClockPhoneName string   `xml:"NewX_AVM-DE_AlarmClockPhoneName"`
}

func (client *ServiceClient) X_AVM_DE_GetAlarmClock(in *X_AVM_DE_GetAlarmClockRequest, out *X_AVM_DE_GetAlarmClockResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetAlarmClockRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetAlarmClockRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetAlarmClockResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetAlarmClockResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetAlarmClock", soapRequest, soapResponse)
}

type X_AVM_DE_SetAlarmClockEnableRequest struct {
	XMLName                      xml.Name `xml:"u:X_AVM-DE_SetAlarmClockEnableRequest"`
	XMLNameSpace                 string   `xml:"xmlns:u,attr"`
	NewIndex                     uint16   `xml:"NewIndex"`
	NewX_AVM_DE_AlarmClockEnable bool     `xml:"NewX_AVM-DE_AlarmClockEnable"`
}

type X_AVM_DE_SetAlarmClockEnableResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetAlarmClockEnableResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetAlarmClockEnable(in *X_AVM_DE_SetAlarmClockEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetAlarmClockEnableRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetAlarmClockEnableRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetAlarmClockEnableResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetAlarmClockEnableResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetAlarmClockEnableResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetAlarmClockEnable", soapRequest, soapResponse)
}

type X_AVM_DE_GetNumberOfAlarmClocksRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetNumberOfAlarmClocksRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetNumberOfAlarmClocksResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetNumberOfAlarmClocksResponse"`
	NewX_AVM_DE_NumberOfAlarmClocks uint16   `xml:"NewX_AVM-DE_NumberOfAlarmClocks"`
}

func (client *ServiceClient) X_AVM_DE_GetNumberOfAlarmClocks(out *X_AVM_DE_GetNumberOfAlarmClocksResponse) error {
	in := &X_AVM_DE_GetNumberOfAlarmClocksRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetNumberOfAlarmClocksRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetNumberOfAlarmClocksRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetNumberOfAlarmClocksResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetNumberOfAlarmClocksResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfAlarmClocks", soapRequest, soapResponse)
}
