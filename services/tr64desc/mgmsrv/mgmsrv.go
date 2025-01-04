// generated from spec version: 1.0
package mgmsrv

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
	XMLName                      xml.Name `xml:"GetInfoResponse"`
	NewURL                       string   `xml:"NewURL"`
	NewUsername                  string   `xml:"NewUsername"`
	NewPeriodicInformEnable      bool     `xml:"NewPeriodicInformEnable"`
	NewPeriodicInformInterval    uint32   `xml:"NewPeriodicInformInterval"`
	NewPeriodicInformTime        string   `xml:"NewPeriodicInformTime"`
	NewParameterKey              string   `xml:"NewParameterKey"`
	NewParameterHash             string   `xml:"NewParameterHash"`
	NewConnectionRequestURL      string   `xml:"NewConnectionRequestURL"`
	NewConnectionRequestUsername string   `xml:"NewConnectionRequestUsername"`
	NewUpgradesManaged           bool     `xml:"NewUpgradesManaged"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetManagementServerURLRequest struct {
	XMLName      xml.Name `xml:"u:SetManagementServerURLRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewURL       string   `xml:"NewURL"`
}

type SetManagementServerURLResponse struct {
	XMLName xml.Name `xml:"SetManagementServerURLResponse"`
}

func (client *ServiceClient) SetManagementServerURL(in *SetManagementServerURLRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetManagementServerURLResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetManagementServerURL", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetManagementServerUsernameRequest struct {
	XMLName      xml.Name `xml:"u:SetManagementServerUsernameRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewUsername  string   `xml:"NewUsername"`
}

type SetManagementServerUsernameResponse struct {
	XMLName xml.Name `xml:"SetManagementServerUsernameResponse"`
}

func (client *ServiceClient) SetManagementServerUsername(in *SetManagementServerUsernameRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetManagementServerUsernameResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetManagementServerUsername", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetManagementServerPasswordRequest struct {
	XMLName      xml.Name `xml:"u:SetManagementServerPasswordRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewPassword  string   `xml:"NewPassword"`
}

type SetManagementServerPasswordResponse struct {
	XMLName xml.Name `xml:"SetManagementServerPasswordResponse"`
}

func (client *ServiceClient) SetManagementServerPassword(in *SetManagementServerPasswordRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetManagementServerPasswordResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetManagementServerPassword", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetPeriodicInformRequest struct {
	XMLName                   xml.Name `xml:"u:SetPeriodicInformRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPeriodicInformEnable   bool     `xml:"NewPeriodicInformEnable"`
	NewPeriodicInformInterval uint32   `xml:"NewPeriodicInformInterval"`
	NewPeriodicInformTime     string   `xml:"NewPeriodicInformTime"`
}

type SetPeriodicInformResponse struct {
	XMLName xml.Name `xml:"SetPeriodicInformResponse"`
}

func (client *ServiceClient) SetPeriodicInform(in *SetPeriodicInformRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetPeriodicInformResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetPeriodicInform", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetConnectionRequestAuthenticationRequest struct {
	XMLName                      xml.Name `xml:"u:SetConnectionRequestAuthenticationRequest"`
	XMLNameSpace                 string   `xml:"xmlns:u,attr"`
	NewConnectionRequestUsername string   `xml:"NewConnectionRequestUsername"`
	NewConnectionRequestPassword string   `xml:"NewConnectionRequestPassword"`
}

type SetConnectionRequestAuthenticationResponse struct {
	XMLName xml.Name `xml:"SetConnectionRequestAuthenticationResponse"`
}

func (client *ServiceClient) SetConnectionRequestAuthentication(in *SetConnectionRequestAuthenticationRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetConnectionRequestAuthenticationResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConnectionRequestAuthentication", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetUpgradeManagementRequest struct {
	XMLName            xml.Name `xml:"u:SetUpgradeManagementRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewUpgradesManaged bool     `xml:"NewUpgradesManaged"`
}

type SetUpgradeManagementResponse struct {
	XMLName xml.Name `xml:"SetUpgradeManagementResponse"`
}

func (client *ServiceClient) SetUpgradeManagement(in *SetUpgradeManagementRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetUpgradeManagementResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetUpgradeManagement", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_SetTR069EnableRequest struct {
	XMLName         xml.Name `xml:"u:X_SetTR069EnableRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewTR069Enabled bool     `xml:"NewTR069Enabled"`
}

type X_SetTR069EnableResponse struct {
	XMLName xml.Name `xml:"X_SetTR069EnableResponse"`
}

func (client *ServiceClient) X_SetTR069Enable(in *X_SetTR069EnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &X_SetTR069EnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_SetTR069Enable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetTR069FirmwareDownloadEnabledRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetTR069FirmwareDownloadEnabledRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetTR069FirmwareDownloadEnabledResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetTR069FirmwareDownloadEnabledResponse"`
	NewTR069FirmwareDownloadEnabled bool     `xml:"NewTR069FirmwareDownloadEnabled"`
}

func (client *ServiceClient) X_AVM_DE_GetTR069FirmwareDownloadEnabled(out *X_AVM_DE_GetTR069FirmwareDownloadEnabledResponse) error {
	in := &X_AVM_DE_GetTR069FirmwareDownloadEnabledRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetTR069FirmwareDownloadEnabled", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_SetTR069FirmwareDownloadEnabledRequest struct {
	XMLName                         xml.Name `xml:"u:X_AVM-DE_SetTR069FirmwareDownloadEnabledRequest"`
	XMLNameSpace                    string   `xml:"xmlns:u,attr"`
	NewTR069FirmwareDownloadEnabled bool     `xml:"NewTR069FirmwareDownloadEnabled"`
}

type X_AVM_DE_SetTR069FirmwareDownloadEnabledResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetTR069FirmwareDownloadEnabledResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetTR069FirmwareDownloadEnabled(in *X_AVM_DE_SetTR069FirmwareDownloadEnabledRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &X_AVM_DE_SetTR069FirmwareDownloadEnabledResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetTR069FirmwareDownloadEnabled", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
