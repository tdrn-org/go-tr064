// generated from spec version: 1.0
package x_tam

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
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetInfoResponse struct {
	XMLName         xml.Name `xml:"GetInfoResponse"`
	NewEnable       bool     `xml:"NewEnable"`
	NewName         string   `xml:"NewName"`
	NewTAMRunning   bool     `xml:"NewTAMRunning"`
	NewStick        uint16   `xml:"NewStick"`
	NewStatus       uint16   `xml:"NewStatus"`
	NewCapacity     uint32   `xml:"NewCapacity"`
	NewMode         string   `xml:"NewMode"`
	NewRingSeconds  uint16   `xml:"NewRingSeconds"`
	NewPhoneNumbers string   `xml:"NewPhoneNumbers"`
}

func (client *ServiceClient) GetInfo(in *GetInfoRequest, out *GetInfoResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
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

type GetMessageListRequest struct {
	XMLName      xml.Name `xml:"u:GetMessageListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetMessageListResponse struct {
	XMLName xml.Name `xml:"GetMessageListResponse"`
	NewURL  string   `xml:"NewURL"`
}

func (client *ServiceClient) GetMessageList(in *GetMessageListRequest, out *GetMessageListResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetMessageList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type MarkMessageRequest struct {
	XMLName         xml.Name `xml:"u:MarkMessageRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewIndex        uint16   `xml:"NewIndex"`
	NewMessageIndex uint16   `xml:"NewMessageIndex"`
	NewMarkedAsRead bool     `xml:"NewMarkedAsRead"`
}

type MarkMessageResponse struct {
	XMLName xml.Name `xml:"MarkMessageResponse"`
}

func (client *ServiceClient) MarkMessage(in *MarkMessageRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &MarkMessageResponse{}
	return client.TR064Client.InvokeService(client.Service, "MarkMessage", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeleteMessageRequest struct {
	XMLName         xml.Name `xml:"u:DeleteMessageRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewIndex        uint16   `xml:"NewIndex"`
	NewMessageIndex uint16   `xml:"NewMessageIndex"`
}

type DeleteMessageResponse struct {
	XMLName xml.Name `xml:"DeleteMessageResponse"`
}

func (client *ServiceClient) DeleteMessage(in *DeleteMessageRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeleteMessageResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeleteMessage", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetListRequest struct {
	XMLName      xml.Name `xml:"u:GetListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetListResponse struct {
	XMLName    xml.Name `xml:"GetListResponse"`
	NewTAMList string   `xml:"NewTAMList"`
}

func (client *ServiceClient) GetList(out *GetListResponse) error {
	in := &GetListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
