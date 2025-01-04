// generated from spec version: 1.0
package igdconn

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetConnectionTypeRequest struct {
	XMLName           xml.Name `xml:"u:SetConnectionTypeRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewConnectionType string   `xml:"NewConnectionType"`
}

type SetConnectionTypeResponse struct {
	XMLName xml.Name `xml:"SetConnectionTypeResponse"`
}

func (client *ServiceClient) SetConnectionType(in *SetConnectionTypeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetConnectionTypeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetConnectionType", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetConnectionTypeInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetConnectionTypeInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetConnectionTypeInfoResponse struct {
	XMLName                    xml.Name `xml:"GetConnectionTypeInfoResponse"`
	NewConnectionType          string   `xml:"NewConnectionType"`
	NewPossibleConnectionTypes string   `xml:"NewPossibleConnectionTypes"`
}

func (client *ServiceClient) GetConnectionTypeInfo(out *GetConnectionTypeInfoResponse) error {
	in := &GetConnectionTypeInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetConnectionTypeInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetAutoDisconnectTimeRequest struct {
	XMLName      xml.Name `xml:"u:GetAutoDisconnectTimeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAutoDisconnectTimeResponse struct {
	XMLName               xml.Name `xml:"GetAutoDisconnectTimeResponse"`
	NewAutoDisconnectTime uint32   `xml:"NewAutoDisconnectTime"`
}

func (client *ServiceClient) GetAutoDisconnectTime(out *GetAutoDisconnectTimeResponse) error {
	in := &GetAutoDisconnectTimeRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetAutoDisconnectTime", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetIdleDisconnectTimeRequest struct {
	XMLName      xml.Name `xml:"u:GetIdleDisconnectTimeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetIdleDisconnectTimeResponse struct {
	XMLName               xml.Name `xml:"GetIdleDisconnectTimeResponse"`
	NewIdleDisconnectTime uint32   `xml:"NewIdleDisconnectTime"`
}

func (client *ServiceClient) GetIdleDisconnectTime(out *GetIdleDisconnectTimeResponse) error {
	in := &GetIdleDisconnectTimeRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetIdleDisconnectTime", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type RequestConnectionRequest struct {
	XMLName      xml.Name `xml:"u:RequestConnectionRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RequestConnectionResponse struct {
	XMLName xml.Name `xml:"RequestConnectionResponse"`
}

func (client *ServiceClient) RequestConnection() error {
	in := &RequestConnectionRequest{XMLNameSpace: client.Service.Type()}
	out := &RequestConnectionResponse{}
	return client.TR064Client.InvokeService(client.Service, "RequestConnection", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type RequestTerminationRequest struct {
	XMLName      xml.Name `xml:"u:RequestTerminationRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RequestTerminationResponse struct {
	XMLName xml.Name `xml:"RequestTerminationResponse"`
}

func (client *ServiceClient) RequestTermination() error {
	in := &RequestTerminationRequest{XMLNameSpace: client.Service.Type()}
	out := &RequestTerminationResponse{}
	return client.TR064Client.InvokeService(client.Service, "RequestTermination", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type ForceTerminationRequest struct {
	XMLName      xml.Name `xml:"u:ForceTerminationRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type ForceTerminationResponse struct {
	XMLName xml.Name `xml:"ForceTerminationResponse"`
}

func (client *ServiceClient) ForceTermination() error {
	in := &ForceTerminationRequest{XMLNameSpace: client.Service.Type()}
	out := &ForceTerminationResponse{}
	return client.TR064Client.InvokeService(client.Service, "ForceTermination", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetStatusInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetStatusInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatusInfoResponse struct {
	XMLName                xml.Name `xml:"GetStatusInfoResponse"`
	NewConnectionStatus    string   `xml:"NewConnectionStatus"`
	NewLastConnectionError string   `xml:"NewLastConnectionError"`
	NewUptime              uint32   `xml:"NewUptime"`
}

func (client *ServiceClient) GetStatusInfo(out *GetStatusInfoResponse) error {
	in := &GetStatusInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetStatusInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetNATRSIPStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetNATRSIPStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNATRSIPStatusResponse struct {
	XMLName          xml.Name `xml:"GetNATRSIPStatusResponse"`
	NewRSIPAvailable bool     `xml:"NewRSIPAvailable"`
	NewNATEnabled    bool     `xml:"NewNATEnabled"`
}

func (client *ServiceClient) GetNATRSIPStatus(out *GetNATRSIPStatusResponse) error {
	in := &GetNATRSIPStatusRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetNATRSIPStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetGenericPortMappingEntryRequest struct {
	XMLName             xml.Name `xml:"u:GetGenericPortMappingEntryRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewPortMappingIndex uint16   `xml:"NewPortMappingIndex"`
}

type GetGenericPortMappingEntryResponse struct {
	XMLName                   xml.Name `xml:"GetGenericPortMappingEntryResponse"`
	NewRemoteHost             string   `xml:"NewRemoteHost"`
	NewExternalPort           uint16   `xml:"NewExternalPort"`
	NewProtocol               string   `xml:"NewProtocol"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

func (client *ServiceClient) GetGenericPortMappingEntry(in *GetGenericPortMappingEntryRequest, out *GetGenericPortMappingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetGenericPortMappingEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetSpecificPortMappingEntryRequest struct {
	XMLName         xml.Name `xml:"u:GetSpecificPortMappingEntryRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewRemoteHost   string   `xml:"NewRemoteHost"`
	NewExternalPort uint16   `xml:"NewExternalPort"`
	NewProtocol     string   `xml:"NewProtocol"`
}

type GetSpecificPortMappingEntryResponse struct {
	XMLName                   xml.Name `xml:"GetSpecificPortMappingEntryResponse"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

func (client *ServiceClient) GetSpecificPortMappingEntry(in *GetSpecificPortMappingEntryRequest, out *GetSpecificPortMappingEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetSpecificPortMappingEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type AddPortMappingRequest struct {
	XMLName                   xml.Name `xml:"u:AddPortMappingRequest"`
	XMLNameSpace              string   `xml:"xmlns:u,attr"`
	NewRemoteHost             string   `xml:"NewRemoteHost"`
	NewExternalPort           uint16   `xml:"NewExternalPort"`
	NewProtocol               string   `xml:"NewProtocol"`
	NewInternalPort           uint16   `xml:"NewInternalPort"`
	NewInternalClient         string   `xml:"NewInternalClient"`
	NewEnabled                bool     `xml:"NewEnabled"`
	NewPortMappingDescription string   `xml:"NewPortMappingDescription"`
	NewLeaseDuration          uint32   `xml:"NewLeaseDuration"`
}

type AddPortMappingResponse struct {
	XMLName xml.Name `xml:"AddPortMappingResponse"`
}

func (client *ServiceClient) AddPortMapping(in *AddPortMappingRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &AddPortMappingResponse{}
	return client.TR064Client.InvokeService(client.Service, "AddPortMapping", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DeletePortMappingRequest struct {
	XMLName         xml.Name `xml:"u:DeletePortMappingRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewRemoteHost   string   `xml:"NewRemoteHost"`
	NewExternalPort uint16   `xml:"NewExternalPort"`
	NewProtocol     string   `xml:"NewProtocol"`
}

type DeletePortMappingResponse struct {
	XMLName xml.Name `xml:"DeletePortMappingResponse"`
}

func (client *ServiceClient) DeletePortMapping(in *DeletePortMappingRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DeletePortMappingResponse{}
	return client.TR064Client.InvokeService(client.Service, "DeletePortMapping", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetExternalIPAddressRequest struct {
	XMLName      xml.Name `xml:"u:GetExternalIPAddressRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetExternalIPAddressResponse struct {
	XMLName              xml.Name `xml:"GetExternalIPAddressResponse"`
	NewExternalIPAddress string   `xml:"NewExternalIPAddress"`
}

func (client *ServiceClient) GetExternalIPAddress(out *GetExternalIPAddressResponse) error {
	in := &GetExternalIPAddressRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetExternalIPAddress", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetExternalIPv6AddressRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetExternalIPv6AddressRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetExternalIPv6AddressResponse struct {
	XMLName                xml.Name `xml:"X_AVM_DE_GetExternalIPv6AddressResponse"`
	NewExternalIPv6Address string   `xml:"NewExternalIPv6Address"`
	NewPrefixLength        uint8    `xml:"NewPrefixLength"`
	NewValidLifetime       uint32   `xml:"NewValidLifetime"`
	NewPreferedLifetime    uint32   `xml:"NewPreferedLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetExternalIPv6Address(out *X_AVM_DE_GetExternalIPv6AddressResponse) error {
	in := &X_AVM_DE_GetExternalIPv6AddressRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetExternalIPv6Address", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetIPv6PrefixRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetIPv6PrefixRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPv6PrefixResponse struct {
	XMLName             xml.Name `xml:"X_AVM_DE_GetIPv6PrefixResponse"`
	NewIPv6Prefix       string   `xml:"NewIPv6Prefix"`
	NewPrefixLength     uint8    `xml:"NewPrefixLength"`
	NewValidLifetime    uint32   `xml:"NewValidLifetime"`
	NewPreferedLifetime uint32   `xml:"NewPreferedLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetIPv6Prefix(out *X_AVM_DE_GetIPv6PrefixResponse) error {
	in := &X_AVM_DE_GetIPv6PrefixRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetIPv6Prefix", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetDNSServerRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetDNSServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDNSServerResponse struct {
	XMLName           xml.Name `xml:"X_AVM_DE_GetDNSServerResponse"`
	NewIPv4DNSServer1 string   `xml:"NewIPv4DNSServer1"`
	NewIPv4DNSServer2 string   `xml:"NewIPv4DNSServer2"`
}

func (client *ServiceClient) X_AVM_DE_GetDNSServer(out *X_AVM_DE_GetDNSServerResponse) error {
	in := &X_AVM_DE_GetDNSServerRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetDNSServer", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetIPv6DNSServerRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM_DE_GetIPv6DNSServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPv6DNSServerResponse struct {
	XMLName           xml.Name `xml:"X_AVM_DE_GetIPv6DNSServerResponse"`
	NewIPv6DNSServer1 string   `xml:"NewIPv6DNSServer1"`
	NewValidLifetime1 uint32   `xml:"NewValidLifetime1"`
	NewIPv6DNSServer2 string   `xml:"NewIPv6DNSServer2"`
	NewValidLifetime  uint32   `xml:"NewValidLifetime"`
}

func (client *ServiceClient) X_AVM_DE_GetIPv6DNSServer(out *X_AVM_DE_GetIPv6DNSServerResponse) error {
	in := &X_AVM_DE_GetIPv6DNSServerRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM_DE_GetIPv6DNSServer", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
