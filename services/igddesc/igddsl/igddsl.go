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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetDSLLinkTypeResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetDSLLinkType", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetDSLLinkInfo", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetAutoConfig", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetModulationType", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetDestinationAddressResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetDestinationAddress", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetDestinationAddress", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetATMEncapsulationResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetATMEncapsulation", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetATMEncapsulation", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetFCSPreservedResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetFCSPreserved", soapRequest, soapResponse)
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
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetFCSPreserved", soapRequest, soapResponse)
}
