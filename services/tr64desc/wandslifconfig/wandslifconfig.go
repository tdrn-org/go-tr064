// generated from spec version: 1.0
package wandslifconfig

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
	XMLName                  xml.Name `xml:"GetInfoResponse"`
	NewEnable                bool     `xml:"NewEnable"`
	NewStatus                string   `xml:"NewStatus"`
	NewDataPath              string   `xml:"NewDataPath"`
	NewUpstreamCurrRate      int64    `xml:"NewUpstreamCurrRate"`
	NewDownstreamCurrRate    uint64   `xml:"NewDownstreamCurrRate"`
	NewUpstreamMaxRate       uint64   `xml:"NewUpstreamMaxRate"`
	NewDownstreamMaxRate     uint64   `xml:"NewDownstreamMaxRate"`
	NewUpstreamNoiseMargin   uint64   `xml:"NewUpstreamNoiseMargin"`
	NewDownstreamNoiseMargin uint64   `xml:"NewDownstreamNoiseMargin"`
	NewUpstreamAttenuation   uint64   `xml:"NewUpstreamAttenuation"`
	NewDownstreamAttenuation uint64   `xml:"NewDownstreamAttenuation"`
	NewATURVendor            string   `xml:"NewATURVendor"`
	NewATURCountry           string   `xml:"NewATURCountry"`
	NewUpstreamPower         uint16   `xml:"NewUpstreamPower"`
	NewDownstreamPower       uint16   `xml:"NewDownstreamPower"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetStatisticsTotalRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsTotalRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsTotalResponse struct {
	XMLName                xml.Name `xml:"GetStatisticsTotalResponse"`
	NewReceiveBlocks       uint64   `xml:"NewReceiveBlocks"`
	NewTransmitBlocks      uint64   `xml:"NewTransmitBlocks"`
	NewCellDelin           uint64   `xml:"NewCellDelin"`
	NewLinkRetrain         uint64   `xml:"NewLinkRetrain"`
	NewInitErrors          uint64   `xml:"NewInitErrors"`
	NewInitTimeouts        uint64   `xml:"NewInitTimeouts"`
	NewLossOfFraming       uint64   `xml:"NewLossOfFraming"`
	NewErroredSecs         uint64   `xml:"NewErroredSecs"`
	NewSeverelyErroredSecs uint64   `xml:"NewSeverelyErroredSecs"`
	NewFECErrors           uint64   `xml:"NewFECErrors"`
	NewATUCFECErrors       uint64   `xml:"NewATUCFECErrors"`
	NewHECErrors           uint64   `xml:"NewHECErrors"`
	NewATUCHECErrors       uint64   `xml:"NewATUCHECErrors"`
	NewCRCErrors           uint64   `xml:"NewCRCErrors"`
	NewATUCCRCErrors       uint64   `xml:"NewATUCCRCErrors"`
}

func (client *ServiceClient) GetStatisticsTotal(out *GetStatisticsTotalResponse) error {
	in := &GetStatisticsTotalRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetStatisticsTotal", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetDSLDiagnoseInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetDSLDiagnoseInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDSLDiagnoseInfoResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetDSLDiagnoseInfoResponse"`
	NewX_AVM_DE_DSLDiagnoseState    string   `xml:"NewX_AVM-DE_DSLDiagnoseState"`
	NewX_AVM_DE_CableNokDistance    int64    `xml:"NewX_AVM-DE_CableNokDistance"`
	NewX_AVM_DE_DSLLastDiagnoseTime uint64   `xml:"NewX_AVM-DE_DSLLastDiagnoseTime"`
	NewX_AVM_DE_DSLSignalLossTime   uint64   `xml:"NewX_AVM-DE_DSLSignalLossTime"`
	NewX_AVM_DE_DSLActive           bool     `xml:"NewX_AVM-DE_DSLActive"`
	NewX_AVM_DE_DSLSync             bool     `xml:"NewX_AVM-DE_DSLSync"`
}

func (client *ServiceClient) X_AVM_DE_GetDSLDiagnoseInfo(out *X_AVM_DE_GetDSLDiagnoseInfoResponse) error {
	in := &X_AVM_DE_GetDSLDiagnoseInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetDSLDiagnoseInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type X_AVM_DE_GetDSLInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetDSLInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDSLInfoResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetDSLInfoResponse"`
	NewSNRGds                uint64   `xml:"NewSNRGds"`
	NewSNRGus                uint64   `xml:"NewSNRGus"`
	NewSNRpsds               string   `xml:"NewSNRpsds"`
	NewSNRpsus               string   `xml:"NewSNRpsus"`
	NewSNRMTds               uint64   `xml:"NewSNRMTds"`
	NewSNRMTus               uint64   `xml:"NewSNRMTus"`
	NewLATNds                string   `xml:"NewLATNds"`
	NewLATNus                string   `xml:"NewLATNus"`
	NewFECErrors             uint64   `xml:"NewFECErrors"`
	NewCRCErrors             uint64   `xml:"NewCRCErrors"`
	NewLinkStatus            string   `xml:"NewLinkStatus"`
	NewModulationType        string   `xml:"NewModulationType"`
	NewCurrentProfile        string   `xml:"NewCurrentProfile"`
	NewUpstreamCurrRate      int64    `xml:"NewUpstreamCurrRate"`
	NewDownstreamCurrRate    uint64   `xml:"NewDownstreamCurrRate"`
	NewUpstreamMaxRate       uint64   `xml:"NewUpstreamMaxRate"`
	NewDownstreamMaxRate     uint64   `xml:"NewDownstreamMaxRate"`
	NewUpstreamNoiseMargin   uint64   `xml:"NewUpstreamNoiseMargin"`
	NewDownstreamNoiseMargin uint64   `xml:"NewDownstreamNoiseMargin"`
	NewUpstreamAttenuation   uint64   `xml:"NewUpstreamAttenuation"`
	NewDownstreamAttenuation uint64   `xml:"NewDownstreamAttenuation"`
	NewATURVendor            string   `xml:"NewATURVendor"`
	NewATURCountry           string   `xml:"NewATURCountry"`
	NewUpstreamPower         uint16   `xml:"NewUpstreamPower"`
	NewDownstreamPower       uint16   `xml:"NewDownstreamPower"`
}

func (client *ServiceClient) X_AVM_DE_GetDSLInfo(out *X_AVM_DE_GetDSLInfoResponse) error {
	in := &X_AVM_DE_GetDSLInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetDSLInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
