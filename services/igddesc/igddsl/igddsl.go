// generated from spec version: 1.0
package igddsl

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetDSLLinkTypeRequest struct {
	XMLName      xml.Name `xml:"u:SetDSLLinkTypeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewLinkType  string   `xml:"NewLinkType"`
}

type SetDSLLinkTypeResponse struct {
	XMLName xml.Name `xml:"SetDSLLinkTypeResponse"`
}

func (client *ServiceClient) SetDSLLinkType(in *SetDSLLinkTypeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDSLLinkTypeResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDSLLinkType", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDSLLinkInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetDSLLinkInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDSLLinkInfoResponse struct {
	XMLName       xml.Name `xml:"GetDSLLinkInfoResponse"`
	NewLinkType   string   `xml:"NewLinkType"`
	NewLinkStatus string   `xml:"NewLinkStatus"`
}

func (client *ServiceClient) GetDSLLinkInfo(out *GetDSLLinkInfoResponse) error {
	in := &GetDSLLinkInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDSLLinkInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetAutoConfigRequest struct {
	XMLName      xml.Name `xml:"u:GetAutoConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetAutoConfigResponse struct {
	XMLName       xml.Name `xml:"GetAutoConfigResponse"`
	NewAutoConfig bool     `xml:"NewAutoConfig"`
}

func (client *ServiceClient) GetAutoConfig(out *GetAutoConfigResponse) error {
	in := &GetAutoConfigRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetAutoConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetModulationTypeRequest struct {
	XMLName      xml.Name `xml:"u:GetModulationTypeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetModulationTypeResponse struct {
	XMLName           xml.Name `xml:"GetModulationTypeResponse"`
	NewModulationType string   `xml:"NewModulationType"`
}

func (client *ServiceClient) GetModulationType(out *GetModulationTypeResponse) error {
	in := &GetModulationTypeRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetModulationType", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetDestinationAddressRequest struct {
	XMLName               xml.Name `xml:"u:SetDestinationAddressRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewDestinationAddress string   `xml:"NewDestinationAddress"`
}

type SetDestinationAddressResponse struct {
	XMLName xml.Name `xml:"SetDestinationAddressResponse"`
}

func (client *ServiceClient) SetDestinationAddress(in *SetDestinationAddressRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDestinationAddressResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDestinationAddress", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDestinationAddressRequest struct {
	XMLName      xml.Name `xml:"u:GetDestinationAddressRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDestinationAddressResponse struct {
	XMLName               xml.Name `xml:"GetDestinationAddressResponse"`
	NewDestinationAddress string   `xml:"NewDestinationAddress"`
}

func (client *ServiceClient) GetDestinationAddress(out *GetDestinationAddressResponse) error {
	in := &GetDestinationAddressRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDestinationAddress", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetATMEncapsulationRequest struct {
	XMLName             xml.Name `xml:"u:SetATMEncapsulationRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewATMEncapsulation string   `xml:"NewATMEncapsulation"`
}

type SetATMEncapsulationResponse struct {
	XMLName xml.Name `xml:"SetATMEncapsulationResponse"`
}

func (client *ServiceClient) SetATMEncapsulation(in *SetATMEncapsulationRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetATMEncapsulationResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetATMEncapsulation", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetATMEncapsulationRequest struct {
	XMLName      xml.Name `xml:"u:GetATMEncapsulationRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetATMEncapsulationResponse struct {
	XMLName             xml.Name `xml:"GetATMEncapsulationResponse"`
	NewATMEncapsulation string   `xml:"NewATMEncapsulation"`
}

func (client *ServiceClient) GetATMEncapsulation(out *GetATMEncapsulationResponse) error {
	in := &GetATMEncapsulationRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetATMEncapsulation", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetFCSPreservedRequest struct {
	XMLName         xml.Name `xml:"u:SetFCSPreservedRequest"`
	XMLNameSpace    string   `xml:"xmlns:u,attr"`
	NewFCSPreserved bool     `xml:"NewFCSPreserved"`
}

type SetFCSPreservedResponse struct {
	XMLName xml.Name `xml:"SetFCSPreservedResponse"`
}

func (client *ServiceClient) SetFCSPreserved(in *SetFCSPreservedRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetFCSPreservedResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetFCSPreserved", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetFCSPreservedRequest struct {
	XMLName      xml.Name `xml:"u:GetFCSPreservedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetFCSPreservedResponse struct {
	XMLName         xml.Name `xml:"GetFCSPreservedResponse"`
	NewFCSPreserved bool     `xml:"NewFCSPreserved"`
}

func (client *ServiceClient) GetFCSPreserved(out *GetFCSPreservedResponse) error {
	in := &GetFCSPreservedRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetFCSPreserved", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
