// generated from spec version: 1.0
package wandsllinkconfig

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
	NewEnable                 bool     `xml:"NewEnable"`
	NewLinkStatus             string   `xml:"NewLinkStatus"`
	NewLinkType               string   `xml:"NewLinkType"`
	NewDestinationAddress     string   `xml:"NewDestinationAddress"`
	NewATMEncapsulation       string   `xml:"NewATMEncapsulation"`
	NewAutoConfig             bool     `xml:"NewAutoConfig"`
	NewATMQoS                 string   `xml:"NewATMQoS"`
	NewATMPeakCellRate        uint32   `xml:"NewATMPeakCellRate"`
	NewATMSustainableCellRate uint32   `xml:"NewATMSustainableCellRate"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableResponse struct {
	XMLName xml.Name `xml:"SetEnableResponse"`
}

func (client *ServiceClient) SetEnable(in *SetEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetEnableResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetEnable", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
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

type GetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsResponse struct {
	XMLName                 xml.Name `xml:"GetStatisticsResponse"`
	NewATMTransmittedBlocks uint32   `xml:"NewATMTransmittedBlocks"`
	NewATMReceivedBlocks    uint32   `xml:"NewATMReceivedBlocks"`
	NewAAL5CRCErrors        uint32   `xml:"NewAAL5CRCErrors"`
	NewATMCRCErrors         uint32   `xml:"NewATMCRCErrors"`
}

func (client *ServiceClient) GetStatistics(out *GetStatisticsResponse) error {
	in := &GetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetStatistics", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
