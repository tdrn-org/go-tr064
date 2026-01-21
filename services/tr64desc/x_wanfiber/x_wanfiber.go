// generated from spec version: 1.0
package x_wanfiber

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
	XMLName                        xml.Name `xml:"GetInfoResponse"`
	NewOpticalSignalLevel          int32    `xml:"NewOpticalSignalLevel"`
	NewLowerOpticalThreshold       int32    `xml:"NewLowerOpticalThreshold"`
	NewUpperOpticalThreshold       int32    `xml:"NewUpperOpticalThreshold"`
	NewTransmitOpticalLevel        int32    `xml:"NewTransmitOpticalLevel"`
	NewLowerTransmitPowerThreshold int32    `xml:"NewLowerTransmitPowerThreshold"`
	NewUpperTransmitPowerThreshold int32    `xml:"NewUpperTransmitPowerThreshold"`
	NewSFPVendor                   string   `xml:"NewSFPVendor"`
	NewSFPPartNumber               string   `xml:"NewSFPPartNumber"`
	NewSFPSerialNumber             string   `xml:"NewSFPSerialNumber"`
	NewSFPType                     int32    `xml:"NewSFPType"`
	NewTXWaveLength                uint32   `xml:"NewTXWaveLength"`
	NewFiberMode                   string   `xml:"NewFiberMode"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetInfoGPONRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoGPONRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoGPONResponse struct {
	XMLName         xml.Name `xml:"GetInfoGPONResponse"`
	NewGPONSerial   string   `xml:"NewGPONSerial"`
	NewPONId        string   `xml:"NewPONId"`
	NewONUId        uint32   `xml:"NewONUId"`
	NewUNIType      string   `xml:"NewUNIType"`
	NewGEMPortCount uint32   `xml:"NewGEMPortCount"`
}

func (client *ServiceClient) GetInfoGPON(out *GetInfoGPONResponse) error {
	in := &GetInfoGPONRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfoGPON", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsResponse struct {
	XMLName                 xml.Name `xml:"GetStatisticsResponse"`
	NewBytesSent            uint32   `xml:"NewBytesSent"`
	NewBytesReceived        uint32   `xml:"NewBytesReceived"`
	NewPacketsSent          uint32   `xml:"NewPacketsSent"`
	NewPacketsReceived      uint32   `xml:"NewPacketsReceived"`
	NewPacketErrorsSent     uint32   `xml:"NewPacketErrorsSent"`
	NewPacketErrorsReceived uint32   `xml:"NewPacketErrorsReceived"`
	NewPacketsMulticast     uint32   `xml:"NewPacketsMulticast"`
	NewConnectionRateDown   uint32   `xml:"NewConnectionRateDown"`
	NewConnectionRateUp     uint32   `xml:"NewConnectionRateUp"`
	NewBestTrainState       uint32   `xml:"NewBestTrainState"`
	NewResyncs              uint32   `xml:"NewResyncs"`
	NewMinutesInShowtime    uint32   `xml:"NewMinutesInShowtime"`
}

func (client *ServiceClient) GetStatistics(out *GetStatisticsResponse) error {
	in := &GetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetStatistics", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
