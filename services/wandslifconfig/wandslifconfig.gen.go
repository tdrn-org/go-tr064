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
	NewUpstreamCurrRate      int32    `xml:"NewUpstreamCurrRate"`
	NewDownstreamCurrRate    uint32   `xml:"NewDownstreamCurrRate"`
	NewUpstreamMaxRate       uint32   `xml:"NewUpstreamMaxRate"`
	NewDownstreamMaxRate     uint32   `xml:"NewDownstreamMaxRate"`
	NewUpstreamNoiseMargin   uint32   `xml:"NewUpstreamNoiseMargin"`
	NewDownstreamNoiseMargin uint32   `xml:"NewDownstreamNoiseMargin"`
	NewUpstreamAttenuation   uint32   `xml:"NewUpstreamAttenuation"`
	NewDownstreamAttenuation uint32   `xml:"NewDownstreamAttenuation"`
	NewATURVendor            string   `xml:"NewATURVendor"`
	NewATURCountry           string   `xml:"NewATURCountry"`
	NewUpstreamPower         uint16   `xml:"NewUpstreamPower"`
	NewDownstreamPower       uint16   `xml:"NewDownstreamPower"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type GetStatisticsTotalRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsTotalRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsTotalResponse struct {
	XMLName                xml.Name `xml:"GetStatisticsTotalResponse"`
	NewReceiveBlocks       uint32   `xml:"NewReceiveBlocks"`
	NewTransmitBlocks      uint32   `xml:"NewTransmitBlocks"`
	NewCellDelin           uint32   `xml:"NewCellDelin"`
	NewLinkRetrain         uint32   `xml:"NewLinkRetrain"`
	NewInitErrors          uint32   `xml:"NewInitErrors"`
	NewInitTimeouts        uint32   `xml:"NewInitTimeouts"`
	NewLossOfFraming       uint32   `xml:"NewLossOfFraming"`
	NewErroredSecs         uint32   `xml:"NewErroredSecs"`
	NewSeverelyErroredSecs uint32   `xml:"NewSeverelyErroredSecs"`
	NewFECErrors           uint32   `xml:"NewFECErrors"`
	NewATUCFECErrors       uint32   `xml:"NewATUCFECErrors"`
	NewHECErrors           uint32   `xml:"NewHECErrors"`
	NewATUCHECErrors       uint32   `xml:"NewATUCHECErrors"`
	NewCRCErrors           uint32   `xml:"NewCRCErrors"`
	NewATUCCRCErrors       uint32   `xml:"NewATUCCRCErrors"`
}

func (client *ServiceClient) GetStatisticsTotal(out *GetStatisticsTotalResponse) error {
	in := &GetStatisticsTotalRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetStatisticsTotalRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[GetStatisticsTotalRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetStatisticsTotalResponse]{
		Body: &tr064.SOAPResponseBody[GetStatisticsTotalResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetStatisticsTotal", soapRequest, soapResponse)
}

type X_AVM_DE_GetDSLDiagnoseInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetDSLDiagnoseInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDSLDiagnoseInfoResponse struct {
	XMLName                         xml.Name `xml:"X_AVM-DE_GetDSLDiagnoseInfoResponse"`
	NewX_AVM_DE_DSLDiagnoseState    string   `xml:"NewX_AVM-DE_DSLDiagnoseState"`
	NewX_AVM_DE_CableNokDistance    int32    `xml:"NewX_AVM-DE_CableNokDistance"`
	NewX_AVM_DE_DSLLastDiagnoseTime uint32   `xml:"NewX_AVM-DE_DSLLastDiagnoseTime"`
	NewX_AVM_DE_DSLSignalLossTime   uint32   `xml:"NewX_AVM-DE_DSLSignalLossTime"`
	NewX_AVM_DE_DSLActive           bool     `xml:"NewX_AVM-DE_DSLActive"`
	NewX_AVM_DE_DSLSync             bool     `xml:"NewX_AVM-DE_DSLSync"`
}

func (client *ServiceClient) X_AVM_DE_GetDSLDiagnoseInfo(out *X_AVM_DE_GetDSLDiagnoseInfoResponse) error {
	in := &X_AVM_DE_GetDSLDiagnoseInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetDSLDiagnoseInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetDSLDiagnoseInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetDSLDiagnoseInfoResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetDSLDiagnoseInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetDSLDiagnoseInfo", soapRequest, soapResponse)
}

type X_AVM_DE_GetDSLInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetDSLInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetDSLInfoResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetDSLInfoResponse"`
	NewSNRGds                uint32   `xml:"NewSNRGds"`
	NewSNRGus                uint32   `xml:"NewSNRGus"`
	NewSNRpsds               string   `xml:"NewSNRpsds"`
	NewSNRpsus               string   `xml:"NewSNRpsus"`
	NewSNRMTds               uint32   `xml:"NewSNRMTds"`
	NewSNRMTus               uint32   `xml:"NewSNRMTus"`
	NewLATNds                string   `xml:"NewLATNds"`
	NewLATNus                string   `xml:"NewLATNus"`
	NewFECErrors             uint32   `xml:"NewFECErrors"`
	NewCRCErrors             uint32   `xml:"NewCRCErrors"`
	NewLinkStatus            string   `xml:"NewLinkStatus"`
	NewModulationType        string   `xml:"NewModulationType"`
	NewCurrentProfile        string   `xml:"NewCurrentProfile"`
	NewUpstreamCurrRate      int32    `xml:"NewUpstreamCurrRate"`
	NewDownstreamCurrRate    uint32   `xml:"NewDownstreamCurrRate"`
	NewUpstreamMaxRate       uint32   `xml:"NewUpstreamMaxRate"`
	NewDownstreamMaxRate     uint32   `xml:"NewDownstreamMaxRate"`
	NewUpstreamNoiseMargin   uint32   `xml:"NewUpstreamNoiseMargin"`
	NewDownstreamNoiseMargin uint32   `xml:"NewDownstreamNoiseMargin"`
	NewUpstreamAttenuation   uint32   `xml:"NewUpstreamAttenuation"`
	NewDownstreamAttenuation uint32   `xml:"NewDownstreamAttenuation"`
	NewATURVendor            string   `xml:"NewATURVendor"`
	NewATURCountry           string   `xml:"NewATURCountry"`
	NewUpstreamPower         uint16   `xml:"NewUpstreamPower"`
	NewDownstreamPower       uint16   `xml:"NewDownstreamPower"`
}

func (client *ServiceClient) X_AVM_DE_GetDSLInfo(out *X_AVM_DE_GetDSLInfoResponse) error {
	in := &X_AVM_DE_GetDSLInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetDSLInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetDSLInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetDSLInfoResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetDSLInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetDSLInfo", soapRequest, soapResponse)
}
