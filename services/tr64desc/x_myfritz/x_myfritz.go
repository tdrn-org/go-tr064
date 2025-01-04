// generated from spec version: 1.0
package x_myfritz

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
	XMLName             xml.Name `xml:"GetInfoResponse"`
	NewEnabled          bool     `xml:"NewEnabled"`
	NewDeviceRegistered bool     `xml:"NewDeviceRegistered"`
	NewDynDNSName       string   `xml:"NewDynDNSName"`
	NewPort             uint32   `xml:"NewPort"`
	NewState            string   `xml:"NewState"`
	NewEmail            string   `xml:"NewEmail"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetMyFRITZRequest struct {
	XMLName      xml.Name `xml:"u:SetMyFRITZRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnabled   bool     `xml:"NewEnabled"`
	NewEmail     string   `xml:"NewEmail"`
}

type SetMyFRITZResponse struct {
	XMLName xml.Name `xml:"SetMyFRITZResponse"`
}

func (client *ServiceClient) SetMyFRITZ(in *SetMyFRITZRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetMyFRITZResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetMyFRITZ", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetNumberOfServicesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfServicesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfServicesResponse struct {
	XMLName             xml.Name `xml:"GetNumberOfServicesResponse"`
	NewNumberOfServices uint32   `xml:"NewNumberOfServices"`
}

func (client *ServiceClient) GetNumberOfServices(out *GetNumberOfServicesResponse) error {
	in := &GetNumberOfServicesRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfServices", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetServiceByIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetServiceByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type GetServiceByIndexResponse struct {
	XMLName                  xml.Name `xml:"GetServiceByIndexResponse"`
	NewEnabled               bool     `xml:"NewEnabled"`
	NewName                  string   `xml:"NewName"`
	NewScheme                string   `xml:"NewScheme"`
	NewPort                  uint32   `xml:"NewPort"`
	NewURLPath               string   `xml:"NewURLPath"`
	NewType                  string   `xml:"NewType"`
	NewIPv4ForwardingWarning uint8    `xml:"NewIPv4ForwardingWarning"`
	NewIPv4Addresses         string   `xml:"NewIPv4Addresses"`
	NewIPv6Addresses         string   `xml:"NewIPv6Addresses"`
	NewIPv6InterfaceIDs      string   `xml:"NewIPv6InterfaceIDs"`
	NewMACAddress            string   `xml:"NewMACAddress"`
	NewHostName              string   `xml:"NewHostName"`
	NewDynDnsLabel           string   `xml:"NewDynDnsLabel"`
	NewStatus                uint32   `xml:"NewStatus"`
}

func (client *ServiceClient) GetServiceByIndex(in *GetServiceByIndexRequest, out *GetServiceByIndexResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetServiceByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetServiceByIndexRequest struct {
	XMLName            xml.Name `xml:"u:SetServiceByIndexRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewIndex           uint32   `xml:"NewIndex"`
	NewEnabled         bool     `xml:"NewEnabled"`
	NewName            string   `xml:"NewName"`
	NewScheme          string   `xml:"NewScheme"`
	NewPort            uint32   `xml:"NewPort"`
	NewURLPath         string   `xml:"NewURLPath"`
	NewType            string   `xml:"NewType"`
	NewIPv4Address     string   `xml:"NewIPv4Address"`
	NewIPv6Address     string   `xml:"NewIPv6Address"`
	NewIPv6InterfaceID string   `xml:"NewIPv6InterfaceID"`
	NewMACAddress      string   `xml:"NewMACAddress"`
	NewHostName        string   `xml:"NewHostName"`
}

type SetServiceByIndexResponse struct {
	XMLName xml.Name `xml:"SetServiceByIndexResponse"`
}

func (client *ServiceClient) SetServiceByIndex(in *SetServiceByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetServiceByIndexResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetServiceByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeleteServiceByIndexRequest struct {
	XMLName      xml.Name `xml:"u:DeleteServiceByIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint32   `xml:"NewIndex"`
}

type DeleteServiceByIndexResponse struct {
	XMLName xml.Name `xml:"DeleteServiceByIndexResponse"`
}

func (client *ServiceClient) DeleteServiceByIndex(in *DeleteServiceByIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeleteServiceByIndexResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeleteServiceByIndex", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
