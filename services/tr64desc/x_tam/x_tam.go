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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetEnable", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetMessageList", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &MarkMessageResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "MarkMessage", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &DeleteMessageResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "DeleteMessage", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetList", soapRequest, soapResponse)
}
