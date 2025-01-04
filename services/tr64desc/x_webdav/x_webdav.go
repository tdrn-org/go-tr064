// generated from spec version: 1.0
package x_webdav

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
	XMLName           xml.Name `xml:"GetInfoResponse"`
	NewEnable         bool     `xml:"NewEnable"`
	NewHostURL        string   `xml:"NewHostURL"`
	NewUsername       string   `xml:"NewUsername"`
	NewMountpointName string   `xml:"NewMountpointName"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetConfigRequest struct {
	XMLName           xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewEnable         bool     `xml:"NewEnable"`
	NewHostURL        string   `xml:"NewHostURL"`
	NewUsername       string   `xml:"NewUsername"`
	NewPassword       string   `xml:"NewPassword"`
	NewMountpointName string   `xml:"NewMountpointName"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetConfigResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
