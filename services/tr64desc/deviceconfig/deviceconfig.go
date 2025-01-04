// generated from spec version: 1.0
package deviceconfig

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetPersistentDataRequest struct {
	XMLName      xml.Name `xml:"u:GetPersistentDataRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetPersistentDataResponse struct {
	XMLName           xml.Name `xml:"GetPersistentDataResponse"`
	NewPersistentData string   `xml:"NewPersistentData"`
}

func (client *ServiceClient) GetPersistentData(out *GetPersistentDataResponse) error {
	in := &GetPersistentDataRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetPersistentData", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetPersistentDataRequest struct {
	XMLName           xml.Name `xml:"u:SetPersistentDataRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewPersistentData string   `xml:"NewPersistentData"`
}

type SetPersistentDataResponse struct {
	XMLName xml.Name `xml:"SetPersistentDataResponse"`
}

func (client *ServiceClient) SetPersistentData(in *SetPersistentDataRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetPersistentDataResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetPersistentData", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type ConfigurationStartedRequest struct {
	XMLName      xml.Name `xml:"u:ConfigurationStartedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewSessionID string   `xml:"NewSessionID"`
}

type ConfigurationStartedResponse struct {
	XMLName xml.Name `xml:"ConfigurationStartedResponse"`
}

func (client *ServiceClient) ConfigurationStarted(in *ConfigurationStartedRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &ConfigurationStartedResponse{}
	return client.TR064Client.InvokeService(client.Service, "ConfigurationStarted", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type ConfigurationFinishedRequest struct {
	XMLName      xml.Name `xml:"u:ConfigurationFinishedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type ConfigurationFinishedResponse struct {
	XMLName   xml.Name `xml:"ConfigurationFinishedResponse"`
	NewStatus string   `xml:"NewStatus"`
}

func (client *ServiceClient) ConfigurationFinished(out *ConfigurationFinishedResponse) error {
	in := &ConfigurationFinishedRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "ConfigurationFinished", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type FactoryResetRequest struct {
	XMLName      xml.Name `xml:"u:FactoryResetRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type FactoryResetResponse struct {
	XMLName xml.Name `xml:"FactoryResetResponse"`
}

func (client *ServiceClient) FactoryReset() error {
	in := &FactoryResetRequest{XMLNameSpace: client.Service.Type()}
	out := &FactoryResetResponse{}
	return client.TR064Client.InvokeService(client.Service, "FactoryReset", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type RebootRequest struct {
	XMLName      xml.Name `xml:"u:RebootRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RebootResponse struct {
	XMLName xml.Name `xml:"RebootResponse"`
}

func (client *ServiceClient) Reboot() error {
	in := &RebootRequest{XMLNameSpace: client.Service.Type()}
	out := &RebootResponse{}
	return client.TR064Client.InvokeService(client.Service, "Reboot", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_GenerateUUIDRequest struct {
	XMLName      xml.Name `xml:"u:X_GenerateUUIDRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_GenerateUUIDResponse struct {
	XMLName xml.Name `xml:"X_GenerateUUIDResponse"`
	NewUUID string   `xml:"NewUUID"`
}

func (client *ServiceClient) X_GenerateUUID(out *X_GenerateUUIDResponse) error {
	in := &X_GenerateUUIDRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_GenerateUUID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetConfigFileRequest struct {
	XMLName              xml.Name `xml:"u:X_AVM-DE_GetConfigFileRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_Password string   `xml:"NewX_AVM-DE_Password"`
}

type X_AVM_DE_GetConfigFileResponse struct {
	XMLName                   xml.Name `xml:"X_AVM-DE_GetConfigFileResponse"`
	NewX_AVM_DE_ConfigFileUrl string   `xml:"NewX_AVM-DE_ConfigFileUrl"`
}

func (client *ServiceClient) X_AVM_DE_GetConfigFile(in *X_AVM_DE_GetConfigFileRequest, out *X_AVM_DE_GetConfigFileResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetConfigFile", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_SetConfigFileRequest struct {
	XMLName                   xml.Name `xml:"u:X_AVM-DE_SetConfigFileRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_Password      string   `xml:"NewX_AVM-DE_Password"`
	NewX_AVM_DE_ConfigFileUrl string   `xml:"NewX_AVM-DE_ConfigFileUrl"`
}

type X_AVM_DE_SetConfigFileResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetConfigFileResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetConfigFile(in *X_AVM_DE_SetConfigFileRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &X_AVM_DE_SetConfigFileResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetConfigFile", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_CreateUrlSIDRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_CreateUrlSIDRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_CreateUrlSIDResponse struct {
	XMLName            xml.Name `xml:"X_AVM-DE_CreateUrlSIDResponse"`
	NewX_AVM_DE_UrlSID string   `xml:"NewX_AVM-DE_UrlSID"`
}

func (client *ServiceClient) X_AVM_DE_CreateUrlSID(out *X_AVM_DE_CreateUrlSIDResponse) error {
	in := &X_AVM_DE_CreateUrlSIDRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_CreateUrlSID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_SendSupportDataRequest struct {
	XMLName                     xml.Name `xml:"u:X_AVM-DE_SendSupportDataRequest"`
	XMLNameSpace                string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_SupportDataMode string   `xml:"NewX_AVM-DE_SupportDataMode"`
}

type X_AVM_DE_SendSupportDataResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SendSupportDataResponse"`
}

func (client *ServiceClient) X_AVM_DE_SendSupportData(in *X_AVM_DE_SendSupportDataRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &X_AVM_DE_SendSupportDataResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SendSupportData", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetSupportDataInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetSupportDataInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetSupportDataInfoResponse struct {
	XMLName                          xml.Name `xml:"X_AVM-DE_GetSupportDataInfoResponse"`
	NewX_AVM_DE_SupportDataMode      string   `xml:"NewX_AVM-DE_SupportDataMode"`
	NewX_AVM_DE_SupportDataID        string   `xml:"NewX_AVM-DE_SupportDataID"`
	NewX_AVM_DE_SupportDataTimestamp string   `xml:"NewX_AVM-DE_SupportDataTimestamp"`
	NewX_AVM_DE_SupportDataStatus    string   `xml:"NewX_AVM-DE_SupportDataStatus"`
	NewX_AVM_DE_SupportDataEnabled   bool     `xml:"NewX_AVM-DE_SupportDataEnabled"`
}

func (client *ServiceClient) X_AVM_DE_GetSupportDataInfo(out *X_AVM_DE_GetSupportDataInfoResponse) error {
	in := &X_AVM_DE_GetSupportDataInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetSupportDataInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetSupportDataEnableRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetSupportDataEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetSupportDataEnableResponse struct {
	XMLName                        xml.Name `xml:"X_AVM-DE_GetSupportDataEnableResponse"`
	NewX_AVM_DE_SupportDataEnabled bool     `xml:"NewX_AVM-DE_SupportDataEnabled"`
}

func (client *ServiceClient) X_AVM_DE_GetSupportDataEnable(out *X_AVM_DE_GetSupportDataEnableResponse) error {
	in := &X_AVM_DE_GetSupportDataEnableRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetSupportDataEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_SetSupportDataEnableRequest struct {
	XMLName                        xml.Name `xml:"u:X_AVM-DE_SetSupportDataEnableRequest"`
	XMLNameSpace                   string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_SupportDataEnabled bool     `xml:"NewX_AVM-DE_SupportDataEnabled"`
}

type X_AVM_DE_SetSupportDataEnableResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetSupportDataEnableResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetSupportDataEnable(in *X_AVM_DE_SetSupportDataEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &X_AVM_DE_SetSupportDataEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetSupportDataEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
