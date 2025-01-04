// generated from spec version: 1.0
package x_contact

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
	XMLName        xml.Name `xml:"GetInfoResponse"`
	NewEnable      bool     `xml:"NewEnable"`
	NewStatus      string   `xml:"NewStatus"`
	NewLastConnect string   `xml:"NewLastConnect"`
	NewUrl         string   `xml:"NewUrl"`
	NewServiceId   string   `xml:"NewServiceId"`
	NewUsername    string   `xml:"NewUsername"`
	NewName        string   `xml:"NewName"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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
	out := &SetEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetConfigRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
	NewUrl       string   `xml:"NewUrl"`
	NewServiceId string   `xml:"NewServiceId"`
	NewUsername  string   `xml:"NewUsername"`
	NewPassword  string   `xml:"NewPassword"`
	NewName      string   `xml:"NewName"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetConfigResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetInfoByIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetInfoByIndexResponse struct {
	XMLName        xml.Name `xml:"GetInfoByIndexResponse"`
	NewEnable      bool     `xml:"NewEnable"`
	NewStatus      string   `xml:"NewStatus"`
	NewLastConnect string   `xml:"NewLastConnect"`
	NewUrl         string   `xml:"NewUrl"`
	NewServiceId   string   `xml:"NewServiceId"`
	NewUsername    string   `xml:"NewUsername"`
	NewName        string   `xml:"NewName"`
}

func (client *ServiceClient) GetInfoByIndex(in *GetInfoByIndexRequest, out *GetInfoByIndexResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetInfoByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetEnableByIndexRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableByIndexResponse struct {
	XMLName xml.Name `xml:"SetEnableByIndexResponse"`
}

func (client *ServiceClient) SetEnableByIndex(in *SetEnableByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetEnableByIndexResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetEnableByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetConfigByIndexRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
	NewEnable    bool     `xml:"NewEnable"`
	NewUrl       string   `xml:"NewUrl"`
	NewServiceId string   `xml:"NewServiceId"`
	NewUsername  string   `xml:"NewUsername"`
	NewPassword  string   `xml:"NewPassword"`
	NewName      string   `xml:"NewName"`
}

type SetConfigByIndexResponse struct {
	XMLName xml.Name `xml:"SetConfigByIndexResponse"`
}

func (client *ServiceClient) SetConfigByIndex(in *SetConfigByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetConfigByIndexResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConfigByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeleteByIndexRequest struct {
	XMLName      xml.Name `xml:"u:DeleteByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type DeleteByIndexResponse struct {
	XMLName xml.Name `xml:"DeleteByIndexResponse"`
}

func (client *ServiceClient) DeleteByIndex(in *DeleteByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeleteByIndexResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeleteByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfEntriesResponse struct {
	XMLName                 xml.Name `xml:"GetNumberOfEntriesResponse"`
	NewOnTelNumberOfEntries uint16   `xml:"NewOnTelNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfEntries(out *GetNumberOfEntriesResponse) error {
	in := &GetNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfEntries", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetCallListRequest struct {
	XMLName      xml.Name `xml:"u:GetCallListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetCallListResponse struct {
	XMLName        xml.Name `xml:"GetCallListResponse"`
	NewCallListURL string   `xml:"NewCallListURL"`
}

func (client *ServiceClient) GetCallList(out *GetCallListResponse) error {
	in := &GetCallListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetCallList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetPhonebookListRequest struct {
	XMLName      xml.Name `xml:"u:GetPhonebookListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetPhonebookListResponse struct {
	XMLName          xml.Name `xml:"GetPhonebookListResponse"`
	NewPhonebookList string   `xml:"NewPhonebookList"`
}

func (client *ServiceClient) GetPhonebookList(out *GetPhonebookListResponse) error {
	in := &GetPhonebookListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetPhonebookRequest struct {
	XMLName        xml.Name `xml:"u:GetPhonebookRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

type GetPhonebookResponse struct {
	XMLName             xml.Name `xml:"GetPhonebookResponse"`
	NewPhonebookName    string   `xml:"NewPhonebookName"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
	NewPhonebookURL     string   `xml:"NewPhonebookURL"`
}

func (client *ServiceClient) GetPhonebook(in *GetPhonebookRequest, out *GetPhonebookResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetPhonebook", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type AddPhonebookRequest struct {
	XMLName             xml.Name `xml:"u:AddPhonebookRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
	NewPhonebookName    string   `xml:"NewPhonebookName"`
}

type AddPhonebookResponse struct {
	XMLName xml.Name `xml:"AddPhonebookResponse"`
}

func (client *ServiceClient) AddPhonebook(in *AddPhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &AddPhonebookResponse{}
	return client.TR064Client.InvokeService(client.Service, "AddPhonebook", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeletePhonebookRequest struct {
	XMLName             xml.Name `xml:"u:DeletePhonebookRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookExtraID string   `xml:"NewPhonebookExtraID"`
}

type DeletePhonebookResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookResponse"`
}

func (client *ServiceClient) DeletePhonebook(in *DeletePhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeletePhonebookResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebook", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetPhonebookEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetPhonebookEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type GetPhonebookEntryResponse struct {
	XMLName               xml.Name `xml:"GetPhonebookEntryResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetPhonebookEntry(in *GetPhonebookEntryRequest, out *GetPhonebookEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetPhonebookEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:GetPhonebookEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookID            uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type GetPhonebookEntryUIDResponse struct {
	XMLName               xml.Name `xml:"GetPhonebookEntryUIDResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetPhonebookEntryUID(in *GetPhonebookEntryUIDRequest, out *GetPhonebookEntryUIDResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetPhonebookEntryUID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetPhonebookEntryRequest struct {
	XMLName               xml.Name `xml:"u:SetPhonebookEntryRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID   uint32   `xml:"NewPhonebookEntryID"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetPhonebookEntryResponse struct {
	XMLName xml.Name `xml:"SetPhonebookEntryResponse"`
}

func (client *ServiceClient) SetPhonebookEntry(in *SetPhonebookEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetPhonebookEntryResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetPhonebookEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetPhonebookEntryUIDRequest struct {
	XMLName               xml.Name `xml:"u:SetPhonebookEntryUIDRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetPhonebookEntryUIDResponse struct {
	XMLName                   xml.Name `xml:"SetPhonebookEntryUIDResponse"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

func (client *ServiceClient) SetPhonebookEntryUID(in *SetPhonebookEntryUIDRequest, out *SetPhonebookEntryUIDResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "SetPhonebookEntryUID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeletePhonebookEntryRequest struct {
	XMLName             xml.Name `xml:"u:DeletePhonebookEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookID      uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type DeletePhonebookEntryResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookEntryResponse"`
}

func (client *ServiceClient) DeletePhonebookEntry(in *DeletePhonebookEntryRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeletePhonebookEntryResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebookEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeletePhonebookEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:DeletePhonebookEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookID            uint16   `xml:"NewPhonebookID"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type DeletePhonebookEntryUIDResponse struct {
	XMLName xml.Name `xml:"DeletePhonebookEntryUIDResponse"`
}

func (client *ServiceClient) DeletePhonebookEntryUID(in *DeletePhonebookEntryUIDRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeletePhonebookEntryUIDResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeletePhonebookEntryUID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetCallBarringEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetCallBarringEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryID uint32   `xml:"NewPhonebookEntryID"`
}

type GetCallBarringEntryResponse struct {
	XMLName               xml.Name `xml:"GetCallBarringEntryResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetCallBarringEntry(in *GetCallBarringEntryRequest, out *GetCallBarringEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetCallBarringEntryByNumRequest struct {
	XMLName      xml.Name `xml:"u:GetCallBarringEntryByNumRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewNumber    string   `xml:"NewNumber"`
}

type GetCallBarringEntryByNumResponse struct {
	XMLName               xml.Name `xml:"GetCallBarringEntryByNumResponse"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

func (client *ServiceClient) GetCallBarringEntryByNum(in *GetCallBarringEntryByNumRequest, out *GetCallBarringEntryByNumResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringEntryByNum", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetCallBarringListRequest struct {
	XMLName      xml.Name `xml:"u:GetCallBarringListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetCallBarringListResponse struct {
	XMLName         xml.Name `xml:"GetCallBarringListResponse"`
	NewPhonebookURL string   `xml:"NewPhonebookURL"`
}

func (client *ServiceClient) GetCallBarringList(out *GetCallBarringListResponse) error {
	in := &GetCallBarringListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetCallBarringList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetCallBarringEntryRequest struct {
	XMLName               xml.Name `xml:"u:SetCallBarringEntryRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryData string   `xml:"NewPhonebookEntryData"`
}

type SetCallBarringEntryResponse struct {
	XMLName                   xml.Name `xml:"SetCallBarringEntryResponse"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

func (client *ServiceClient) SetCallBarringEntry(in *SetCallBarringEntryRequest, out *SetCallBarringEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "SetCallBarringEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeleteCallBarringEntryUIDRequest struct {
	XMLName                   xml.Name `xml:"u:DeleteCallBarringEntryUIDRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewPhonebookEntryUniqueID uint32   `xml:"NewPhonebookEntryUniqueID"`
}

type DeleteCallBarringEntryUIDResponse struct {
	XMLName xml.Name `xml:"DeleteCallBarringEntryUIDResponse"`
}

func (client *ServiceClient) DeleteCallBarringEntryUID(in *DeleteCallBarringEntryUIDRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeleteCallBarringEntryUIDResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeleteCallBarringEntryUID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDECTHandsetListRequest struct {
	XMLName      xml.Name `xml:"u:GetDECTHandsetListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDECTHandsetListResponse struct {
	XMLName       xml.Name `xml:"GetDECTHandsetListResponse"`
	NewDectIDList string   `xml:"NewDectIDList"`
}

func (client *ServiceClient) GetDECTHandsetList(out *GetDECTHandsetListResponse) error {
	in := &GetDECTHandsetListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDECTHandsetList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDECTHandsetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetDECTHandsetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewDectID    uint16   `xml:"NewDectID"`
}

type GetDECTHandsetInfoResponse struct {
	XMLName        xml.Name `xml:"GetDECTHandsetInfoResponse"`
	NewHandsetName string   `xml:"NewHandsetName"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

func (client *ServiceClient) GetDECTHandsetInfo(in *GetDECTHandsetInfoRequest, out *GetDECTHandsetInfoResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetDECTHandsetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetDECTHandsetPhonebookRequest struct {
	XMLName        xml.Name `xml:"u:SetDECTHandsetPhonebookRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewDectID      uint16   `xml:"NewDectID"`
	NewPhonebookID uint16   `xml:"NewPhonebookID"`
}

type SetDECTHandsetPhonebookResponse struct {
	XMLName xml.Name `xml:"SetDECTHandsetPhonebookResponse"`
}

func (client *ServiceClient) SetDECTHandsetPhonebook(in *SetDECTHandsetPhonebookRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDECTHandsetPhonebookResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDECTHandsetPhonebook", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetNumberOfDeflectionsRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDeflectionsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDeflectionsResponse struct {
	XMLName                xml.Name `xml:"GetNumberOfDeflectionsResponse"`
	NewNumberOfDeflections uint16   `xml:"NewNumberOfDeflections"`
}

func (client *ServiceClient) GetNumberOfDeflections(out *GetNumberOfDeflectionsResponse) error {
	in := &GetNumberOfDeflectionsRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDeflections", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDeflectionRequest struct {
	XMLName         xml.Name `xml:"u:GetDeflectionRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewDeflectionId uint16   `xml:"NewDeflectionId"`
}

type GetDeflectionResponse struct {
	XMLName               xml.Name `xml:"GetDeflectionResponse"`
	NewEnable             bool     `xml:"NewEnable"`
	NewType               string   `xml:"NewType"`
	NewNumber             string   `xml:"NewNumber"`
	NewDeflectionToNumber string   `xml:"NewDeflectionToNumber"`
	NewMode               string   `xml:"NewMode"`
	NewOutgoing           string   `xml:"NewOutgoing"`
	NewPhonebookID        uint16   `xml:"NewPhonebookID"`
}

func (client *ServiceClient) GetDeflection(in *GetDeflectionRequest, out *GetDeflectionResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetDeflection", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDeflectionsRequest struct {
	XMLName      xml.Name `xml:"u:GetDeflectionsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDeflectionsResponse struct {
	XMLName           xml.Name `xml:"GetDeflectionsResponse"`
	NewDeflectionList string   `xml:"NewDeflectionList"`
}

func (client *ServiceClient) GetDeflections(out *GetDeflectionsResponse) error {
	in := &GetDeflectionsRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDeflections", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetDeflectionEnableRequest struct {
	XMLName         xml.Name `xml:"u:SetDeflectionEnableRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewDeflectionId uint16   `xml:"NewDeflectionId"`
	NewEnable       bool     `xml:"NewEnable"`
}

type SetDeflectionEnableResponse struct {
	XMLName xml.Name `xml:"SetDeflectionEnableResponse"`
}

func (client *ServiceClient) SetDeflectionEnable(in *SetDeflectionEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDeflectionEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDeflectionEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
