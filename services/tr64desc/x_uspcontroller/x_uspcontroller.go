// generated from spec version: 1.0
package x_uspcontroller

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
	XMLName                        xml.Name `xml:"GetInfoResponse"`
	NewMinCharsEndpointID          uint16   `xml:"NewMinCharsEndpointID"`
	NewMaxCharsEndpointID          uint16   `xml:"NewMaxCharsEndpointID"`
	NewAllowedCharsEndpointID      string   `xml:"NewAllowedCharsEndpointID"`
	NewMinCharsHostname            uint16   `xml:"NewMinCharsHostname"`
	NewMaxCharsHostname            uint16   `xml:"NewMaxCharsHostname"`
	NewMinCharsPath                uint16   `xml:"NewMinCharsPath"`
	NewMaxCharsPath                uint16   `xml:"NewMaxCharsPath"`
	NewMinCharsMQTTControllerTopic uint16   `xml:"NewMinCharsMQTTControllerTopic"`
	NewMaxCharsMQTTControllerTopic uint16   `xml:"NewMaxCharsMQTTControllerTopic"`
	NewMinCharsMQTTResponseTopic   uint16   `xml:"NewMinCharsMQTTResponseTopic"`
	NewMaxCharsMQTTResponseTopic   uint16   `xml:"NewMaxCharsMQTTResponseTopic"`
	NewMinCharsUsername            uint16   `xml:"NewMinCharsUsername"`
	NewMaxCharsUsername            uint16   `xml:"NewMaxCharsUsername"`
	NewMinCharsPassword            uint16   `xml:"NewMinCharsPassword"`
	NewMaxCharsPassword            uint16   `xml:"NewMaxCharsPassword"`
	NewUSPMyFRITZEnabled           bool     `xml:"NewUSPMyFRITZEnabled"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type GetUSPControllerByIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetUSPControllerByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type GetUSPControllerByIndexResponse struct {
	XMLName                  xml.Name `xml:"GetUSPControllerByIndexResponse"`
	NewEnable                bool     `xml:"NewEnable"`
	NewEndpointID            string   `xml:"NewEndpointID"`
	NewMTP                   string   `xml:"NewMTP"`
	NewHostname              string   `xml:"NewHostname"`
	NewPath                  string   `xml:"NewPath"`
	NewPort                  uint16   `xml:"NewPort"`
	NewUseTLS                bool     `xml:"NewUseTLS"`
	NewMQTTControllerTopic   string   `xml:"NewMQTTControllerTopic"`
	NewMQTTResponseTopic     string   `xml:"NewMQTTResponseTopic"`
	NewAccessRightSmarthome  bool     `xml:"NewAccessRightSmarthome"`
	NewAccessRightMesh       bool     `xml:"NewAccessRightMesh"`
	NewAccessRightInternet   bool     `xml:"NewAccessRightInternet"`
	NewAccessRightSystem     bool     `xml:"NewAccessRightSystem"`
	NewAccessRightController bool     `xml:"NewAccessRightController"`
	NewAccessRightWiFi       bool     `xml:"NewAccessRightWiFi"`
	NewAccessRightVoIP       bool     `xml:"NewAccessRightVoIP"`
	NewUsername              string   `xml:"NewUsername"`
}

func (client *ServiceClient) GetUSPControllerByIndex(in *GetUSPControllerByIndexRequest, out *GetUSPControllerByIndexResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetUSPControllerByIndex", soapRequest, soapResponse)
}

type GetUSPControllerNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetUSPControllerNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetUSPControllerNumberOfEntriesResponse struct {
	XMLName                         xml.Name `xml:"GetUSPControllerNumberOfEntriesResponse"`
	NewUSPControllerNumberOfEntries uint32   `xml:"NewUSPControllerNumberOfEntries"`
}

func (client *ServiceClient) GetUSPControllerNumberOfEntries(out *GetUSPControllerNumberOfEntriesResponse) error {
	in := &GetUSPControllerNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetUSPControllerNumberOfEntries", soapRequest, soapResponse)
}

type AddUSPControllerRequest struct {
	XMLName                  xml.Name `xml:"u:AddUSPControllerRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewEnable                bool     `xml:"NewEnable"`
	NewEndpointID            string   `xml:"NewEndpointID"`
	NewMTP                   string   `xml:"NewMTP"`
	NewHostname              string   `xml:"NewHostname"`
	NewPath                  string   `xml:"NewPath"`
	NewPort                  uint16   `xml:"NewPort"`
	NewUseTLS                bool     `xml:"NewUseTLS"`
	NewMQTTControllerTopic   string   `xml:"NewMQTTControllerTopic"`
	NewMQTTResponseTopic     string   `xml:"NewMQTTResponseTopic"`
	NewAccessRightSmarthome  bool     `xml:"NewAccessRightSmarthome"`
	NewAccessRightMesh       bool     `xml:"NewAccessRightMesh"`
	NewAccessRightInternet   bool     `xml:"NewAccessRightInternet"`
	NewAccessRightSystem     bool     `xml:"NewAccessRightSystem"`
	NewAccessRightController bool     `xml:"NewAccessRightController"`
	NewAccessRightWiFi       bool     `xml:"NewAccessRightWiFi"`
	NewAccessRightVoIP       bool     `xml:"NewAccessRightVoIP"`
	NewUsername              string   `xml:"NewUsername"`
	NewPassword              string   `xml:"NewPassword"`
}

type AddUSPControllerResponse struct {
	XMLName  xml.Name `xml:"AddUSPControllerResponse"`
	NewIndex uint32   `xml:"NewIndex"`
}

func (client *ServiceClient) AddUSPController(in *AddUSPControllerRequest, out *AddUSPControllerResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "AddUSPController", soapRequest, soapResponse)
}

type DeleteUSPControllerByIndexRequest struct {
	XMLName      xml.Name `xml:"u:DeleteUSPControllerByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type DeleteUSPControllerByIndexResponse struct {
	XMLName xml.Name `xml:"DeleteUSPControllerByIndexResponse"`
}

func (client *ServiceClient) DeleteUSPControllerByIndex(in *DeleteUSPControllerByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &DeleteUSPControllerByIndexResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "DeleteUSPControllerByIndex", soapRequest, soapResponse)
}

type SetUSPControllerEnableByIndexRequest struct {
	XMLName      xml.Name `xml:"u:SetUSPControllerEnableByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetUSPControllerEnableByIndexResponse struct {
	XMLName xml.Name `xml:"SetUSPControllerEnableByIndexResponse"`
}

func (client *ServiceClient) SetUSPControllerEnableByIndex(in *SetUSPControllerEnableByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetUSPControllerEnableByIndexResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetUSPControllerEnableByIndex", soapRequest, soapResponse)
}

type GetUSPMyFRITZEnableRequest struct {
	XMLName      xml.Name `xml:"u:GetUSPMyFRITZEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetUSPMyFRITZEnableResponse struct {
	XMLName              xml.Name `xml:"GetUSPMyFRITZEnableResponse"`
	NewUSPMyFRITZEnabled bool     `xml:"NewUSPMyFRITZEnabled"`
}

func (client *ServiceClient) GetUSPMyFRITZEnable(out *GetUSPMyFRITZEnableResponse) error {
	in := &GetUSPMyFRITZEnableRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetUSPMyFRITZEnable", soapRequest, soapResponse)
}

type SetUSPMyFRITZEnableRequest struct {
	XMLName              xml.Name `xml:"u:SetUSPMyFRITZEnableRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewUSPMyFRITZEnabled bool     `xml:"NewUSPMyFRITZEnabled"`
}

type SetUSPMyFRITZEnableResponse struct {
	XMLName xml.Name `xml:"SetUSPMyFRITZEnableResponse"`
}

func (client *ServiceClient) SetUSPMyFRITZEnable(in *SetUSPMyFRITZEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetUSPMyFRITZEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetUSPMyFRITZEnable", soapRequest, soapResponse)
}
