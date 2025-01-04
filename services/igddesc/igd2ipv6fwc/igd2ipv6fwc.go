// generated from spec version: 1.0
package igd2ipv6fwc

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetFirewallStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetFirewallStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetFirewallStatusResponse struct {
	XMLName               xml.Name `xml:"GetFirewallStatusResponse"`
	FirewallEnabled       bool     `xml:"FirewallEnabled"`
	InboundPinholeAllowed bool     `xml:"InboundPinholeAllowed"`
}

func (client *ServiceClient) GetFirewallStatus(out *GetFirewallStatusResponse) error {
	in := &GetFirewallStatusRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetFirewallStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetOutboundPinholeTimeoutRequest struct {
	XMLName        xml.Name `xml:"u:GetOutboundPinholeTimeoutRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	RemoteHost     string   `xml:"RemoteHost"`
	RemotePort     uint16   `xml:"RemotePort"`
	InternalClient string   `xml:"InternalClient"`
	InternalPort   uint16   `xml:"InternalPort"`
	Protocol       uint16   `xml:"Protocol"`
}

type GetOutboundPinholeTimeoutResponse struct {
	XMLName                xml.Name `xml:"GetOutboundPinholeTimeoutResponse"`
	OutboundPinholeTimeout uint32   `xml:"OutboundPinholeTimeout"`
}

func (client *ServiceClient) GetOutboundPinholeTimeout(in *GetOutboundPinholeTimeoutRequest, out *GetOutboundPinholeTimeoutResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetOutboundPinholeTimeout", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type AddPinholeRequest struct {
	XMLName        xml.Name `xml:"u:AddPinholeRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	RemoteHost     string   `xml:"RemoteHost"`
	RemotePort     uint16   `xml:"RemotePort"`
	InternalClient string   `xml:"InternalClient"`
	InternalPort   uint16   `xml:"InternalPort"`
	Protocol       uint16   `xml:"Protocol"`
	LeaseTime      uint32   `xml:"LeaseTime"`
}

type AddPinholeResponse struct {
	XMLName  xml.Name `xml:"AddPinholeResponse"`
	UniqueID uint16   `xml:"UniqueID"`
}

func (client *ServiceClient) AddPinhole(in *AddPinholeRequest, out *AddPinholeResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "AddPinhole", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type UpdatePinholeRequest struct {
	XMLName      xml.Name `xml:"u:UpdatePinholeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	UniqueID     uint16   `xml:"UniqueID"`
	NewLeaseTime uint32   `xml:"NewLeaseTime"`
}

type UpdatePinholeResponse struct {
	XMLName xml.Name `xml:"UpdatePinholeResponse"`
}

func (client *ServiceClient) UpdatePinhole(in *UpdatePinholeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &UpdatePinholeResponse{}
	return client.TR064Client.InvokeService(client.Service, "UpdatePinhole", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeletePinholeRequest struct {
	XMLName      xml.Name `xml:"u:DeletePinholeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	UniqueID     uint16   `xml:"UniqueID"`
}

type DeletePinholeResponse struct {
	XMLName xml.Name `xml:"DeletePinholeResponse"`
}

func (client *ServiceClient) DeletePinhole(in *DeletePinholeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeletePinholeResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeletePinhole", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetPinholePacketsRequest struct {
	XMLName      xml.Name `xml:"u:GetPinholePacketsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	UniqueID     uint16   `xml:"UniqueID"`
}

type GetPinholePacketsResponse struct {
	XMLName        xml.Name `xml:"GetPinholePacketsResponse"`
	PinholePackets uint32   `xml:"PinholePackets"`
}

func (client *ServiceClient) GetPinholePackets(in *GetPinholePacketsRequest, out *GetPinholePacketsResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetPinholePackets", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type CheckPinholeWorkingRequest struct {
	XMLName      xml.Name `xml:"u:CheckPinholeWorkingRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	UniqueID     uint16   `xml:"UniqueID"`
}

type CheckPinholeWorkingResponse struct {
	XMLName   xml.Name `xml:"CheckPinholeWorkingResponse"`
	IsWorking bool     `xml:"IsWorking"`
}

func (client *ServiceClient) CheckPinholeWorking(in *CheckPinholeWorkingRequest, out *CheckPinholeWorkingResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "CheckPinholeWorking", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
