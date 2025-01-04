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
	return client.TR064Client.InvokeService(client.Service, "GetInfoEx", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_AddVoIPAccountResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_AddVoIPAccount", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPAccount", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_DelVoIPAccountResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DelVoIPAccount", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPAccounts", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetConfigResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetMaxVoIPNumbers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetExistingVoIPNumbers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfClients", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient2", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetClientResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetClient2Response{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient2", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClient3", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClientByClientId", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetClient3Response{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient3", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetClient4", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetClients", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfNumbers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumbers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_DeleteClientResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DeleteClient", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialGetConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_DialSetConfigResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialSetConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_DialNumberResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialNumber", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_DialHangupResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DialHangup", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetPhonePort", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetDelayedCallNotificationResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetDelayedCallNotification", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetVoIPCommonCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPCommonCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetVoIPCommonCountryCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPCommonCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetVoIPCommonCountryCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetVoIPCommonCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetVoIPEnableCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetVoIPEnableCountryCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPEnableCountryCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetVoIPCommonAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetVoIPCommonAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetVoIPCommonAreaCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPCommonAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetVoIPCommonAreaCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetVoIPCommonAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "GetVoIPEnableAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetVoIPEnableAreaCodeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetVoIPEnableAreaCode", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetAlarmClock", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &X_AVM_DE_SetAlarmClockEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetAlarmClockEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNumberOfAlarmClocks", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
