// generated from spec version: 1.0
package lanhostconfigmgm

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
	XMLName                   xml.Name `xml:"GetInfoResponse"`
	NewDHCPServerConfigurable bool     `xml:"NewDHCPServerConfigurable"`
	NewDHCPRelay              bool     `xml:"NewDHCPRelay"`
	NewMinAddress             string   `xml:"NewMinAddress"`
	NewMaxAddress             string   `xml:"NewMaxAddress"`
	NewReservedAddresses      string   `xml:"NewReservedAddresses"`
	NewDHCPServerEnable       bool     `xml:"NewDHCPServerEnable"`
	NewDNSServers             string   `xml:"NewDNSServers"`
	NewDomainName             string   `xml:"NewDomainName"`
	NewIPRouters              string   `xml:"NewIPRouters"`
	NewSubnetMask             string   `xml:"NewSubnetMask"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetDHCPServerEnableRequest struct {
	XMLName             xml.Name `xml:"u:SetDHCPServerEnableRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewDHCPServerEnable bool     `xml:"NewDHCPServerEnable"`
}

type SetDHCPServerEnableResponse struct {
	XMLName xml.Name `xml:"SetDHCPServerEnableResponse"`
}

func (client *ServiceClient) SetDHCPServerEnable(in *SetDHCPServerEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDHCPServerEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDHCPServerEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetIPInterfaceRequest struct {
	XMLName             xml.Name `xml:"u:SetIPInterfaceRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewEnable           bool     `xml:"NewEnable"`
	NewIPAddress        string   `xml:"NewIPAddress"`
	NewSubnetMask       string   `xml:"NewSubnetMask"`
	NewIPAddressingType string   `xml:"NewIPAddressingType"`
}

type SetIPInterfaceResponse struct {
	XMLName xml.Name `xml:"SetIPInterfaceResponse"`
}

func (client *ServiceClient) SetIPInterface(in *SetIPInterfaceRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetIPInterfaceResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetIPInterface", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetAddressRangeRequest struct {
	XMLName      xml.Name `xml:"u:GetAddressRangeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAddressRangeResponse struct {
	XMLName       xml.Name `xml:"GetAddressRangeResponse"`
	NewMinAddress string   `xml:"NewMinAddress"`
	NewMaxAddress string   `xml:"NewMaxAddress"`
}

func (client *ServiceClient) GetAddressRange(out *GetAddressRangeResponse) error {
	in := &GetAddressRangeRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetAddressRange", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetAddressRangeRequest struct {
	XMLName       xml.Name `xml:"u:SetAddressRangeRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMinAddress string   `xml:"NewMinAddress"`
	NewMaxAddress string   `xml:"NewMaxAddress"`
}

type SetAddressRangeResponse struct {
	XMLName xml.Name `xml:"SetAddressRangeResponse"`
}

func (client *ServiceClient) SetAddressRange(in *SetAddressRangeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetAddressRangeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetAddressRange", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetIPRoutersListRequest struct {
	XMLName      xml.Name `xml:"u:GetIPRoutersListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIPRoutersListResponse struct {
	XMLName      xml.Name `xml:"GetIPRoutersListResponse"`
	NewIPRouters string   `xml:"NewIPRouters"`
}

func (client *ServiceClient) GetIPRoutersList(out *GetIPRoutersListResponse) error {
	in := &GetIPRoutersListRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetIPRoutersList", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetIPRouterRequest struct {
	XMLName      xml.Name `xml:"u:SetIPRouterRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIPRouters string   `xml:"NewIPRouters"`
}

type SetIPRouterResponse struct {
	XMLName xml.Name `xml:"SetIPRouterResponse"`
}

func (client *ServiceClient) SetIPRouter(in *SetIPRouterRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetIPRouterResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetIPRouter", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetSubnetMaskRequest struct {
	XMLName      xml.Name `xml:"u:GetSubnetMaskRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetSubnetMaskResponse struct {
	XMLName       xml.Name `xml:"GetSubnetMaskResponse"`
	NewSubnetMask string   `xml:"NewSubnetMask"`
}

func (client *ServiceClient) GetSubnetMask(out *GetSubnetMaskResponse) error {
	in := &GetSubnetMaskRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetSubnetMask", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetSubnetMaskRequest struct {
	XMLName       xml.Name `xml:"u:SetSubnetMaskRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewSubnetMask string   `xml:"NewSubnetMask"`
}

type SetSubnetMaskResponse struct {
	XMLName xml.Name `xml:"SetSubnetMaskResponse"`
}

func (client *ServiceClient) SetSubnetMask(in *SetSubnetMaskRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetSubnetMaskResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetSubnetMask", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDNSServersRequest struct {
	XMLName      xml.Name `xml:"u:GetDNSServersRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDNSServersResponse struct {
	XMLName       xml.Name `xml:"GetDNSServersResponse"`
	NewDNSServers string   `xml:"NewDNSServers"`
}

func (client *ServiceClient) GetDNSServers(out *GetDNSServersResponse) error {
	in := &GetDNSServersRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDNSServers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetIPInterfaceNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetIPInterfaceNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIPInterfaceNumberOfEntriesResponse struct {
	XMLName                       xml.Name `xml:"GetIPInterfaceNumberOfEntriesResponse"`
	NewIPInterfaceNumberOfEntries uint16   `xml:"NewIPInterfaceNumberOfEntries"`
}

func (client *ServiceClient) GetIPInterfaceNumberOfEntries(out *GetIPInterfaceNumberOfEntriesResponse) error {
	in := &GetIPInterfaceNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetIPInterfaceNumberOfEntries", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
