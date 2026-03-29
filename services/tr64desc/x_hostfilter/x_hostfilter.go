// generated from spec version: 1.0
package x_hostfilter

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type MarkTicketRequest struct {
	XMLName      xml.Name `xml:"u:MarkTicketRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type MarkTicketResponse struct {
	XMLName     xml.Name `xml:"MarkTicketResponse"`
	NewTicketID string   `xml:"NewTicketID"`
}

func (client *ServiceClient) MarkTicket(out *MarkTicketResponse) error {
	in := &MarkTicketRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "MarkTicket", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetTicketIDStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetTicketIDStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewTicketID  string   `xml:"NewTicketID"`
}

type GetTicketIDStatusResponse struct {
	XMLName           xml.Name `xml:"GetTicketIDStatusResponse"`
	NewTicketIDStatus string   `xml:"NewTicketIDStatus"`
}

func (client *ServiceClient) GetTicketIDStatus(in *GetTicketIDStatusRequest, out *GetTicketIDStatusResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetTicketIDStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DiscardAllTicketsRequest struct {
	XMLName      xml.Name `xml:"u:DiscardAllTicketsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type DiscardAllTicketsResponse struct {
	XMLName xml.Name `xml:"DiscardAllTicketsResponse"`
}

func (client *ServiceClient) DiscardAllTickets() error {
	in := &DiscardAllTicketsRequest{XMLNameSpace: client.Service.Type()}
	out := &DiscardAllTicketsResponse{}
	return client.TR064Client.InvokeService(client.Service, "DiscardAllTickets", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DisallowWANAccessByIPRequest struct {
	XMLName        xml.Name `xml:"u:DisallowWANAccessByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
	NewDisallow    bool     `xml:"NewDisallow"`
}

type DisallowWANAccessByIPResponse struct {
	XMLName xml.Name `xml:"DisallowWANAccessByIPResponse"`
}

func (client *ServiceClient) DisallowWANAccessByIP(in *DisallowWANAccessByIPRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DisallowWANAccessByIPResponse{}
	return client.TR064Client.InvokeService(client.Service, "DisallowWANAccessByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetWANAccessByIPRequest struct {
	XMLName        xml.Name `xml:"u:GetWANAccessByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
}

type GetWANAccessByIPResponse struct {
	XMLName      xml.Name `xml:"GetWANAccessByIPResponse"`
	NewDisallow  bool     `xml:"NewDisallow"`
	NewWANAccess string   `xml:"NewWANAccess"`
}

func (client *ServiceClient) GetWANAccessByIP(in *GetWANAccessByIPRequest, out *GetWANAccessByIPResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetWANAccessByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetHostEntryByIPRequest struct {
	XMLName        xml.Name `xml:"u:GetHostEntryByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
}

type GetHostEntryByIPResponse struct {
	XMLName             xml.Name `xml:"GetHostEntryByIPResponse"`
	NewHostName         string   `xml:"NewHostName"`
	NewFilterProfileID  string   `xml:"NewFilterProfileID"`
	NewTimeUsed         uint32   `xml:"NewTimeUsed"`
	NewTimeMax          uint32   `xml:"NewTimeMax"`
	NewTicketsInAdvance uint32   `xml:"NewTicketsInAdvance"`
	NewTicketValid      uint32   `xml:"NewTicketValid"`
	NewIsTimeShared     bool     `xml:"NewIsTimeShared"`
	NewWANAccess        string   `xml:"NewWANAccess"`
}

func (client *ServiceClient) GetHostEntryByIP(in *GetHostEntryByIPRequest, out *GetHostEntryByIPResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetHostEntryByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetFilterProfileByIDRequest struct {
	XMLName            xml.Name `xml:"u:GetFilterProfileByIDRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewFilterProfileID string   `xml:"NewFilterProfileID"`
}

type GetFilterProfileByIDResponse struct {
	XMLName          xml.Name `xml:"GetFilterProfileByIDResponse"`
	NewFilterProfile string   `xml:"NewFilterProfile"`
}

func (client *ServiceClient) GetFilterProfileByID(in *GetFilterProfileByIDRequest, out *GetFilterProfileByIDResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetFilterProfileByID", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type AddTicketTimeToHostEntryByIPRequest struct {
	XMLName        xml.Name `xml:"u:AddTicketTimeToHostEntryByIPRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewIPv4Address string   `xml:"NewIPv4Address"`
}

type AddTicketTimeToHostEntryByIPResponse struct {
	XMLName             xml.Name `xml:"AddTicketTimeToHostEntryByIPResponse"`
	NewTimeUsed         uint32   `xml:"NewTimeUsed"`
	NewTimeMax          uint32   `xml:"NewTimeMax"`
	NewTicketsInAdvance uint32   `xml:"NewTicketsInAdvance"`
	NewTicketValid      uint32   `xml:"NewTicketValid"`
}

func (client *ServiceClient) AddTicketTimeToHostEntryByIP(in *AddTicketTimeToHostEntryByIPRequest, out *AddTicketTimeToHostEntryByIPResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "AddTicketTimeToHostEntryByIP", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type AddHostEntryToFilterProfileRequest struct {
	XMLName            xml.Name `xml:"u:AddHostEntryToFilterProfileRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewIPv4Address     string   `xml:"NewIPv4Address"`
	NewFilterProfileID string   `xml:"NewFilterProfileID"`
}

type AddHostEntryToFilterProfileResponse struct {
	XMLName xml.Name `xml:"AddHostEntryToFilterProfileResponse"`
}

func (client *ServiceClient) AddHostEntryToFilterProfile(in *AddHostEntryToFilterProfileRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &AddHostEntryToFilterProfileResponse{}
	return client.TR064Client.InvokeService(client.Service, "AddHostEntryToFilterProfile", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetFilterProfilesRequest struct {
	XMLName      xml.Name `xml:"u:GetFilterProfilesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetFilterProfilesResponse struct {
	XMLName              xml.Name `xml:"GetFilterProfilesResponse"`
	NewFilterProfileList string   `xml:"NewFilterProfileList"`
}

func (client *ServiceClient) GetFilterProfiles(out *GetFilterProfilesResponse) error {
	in := &GetFilterProfilesRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetFilterProfiles", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
