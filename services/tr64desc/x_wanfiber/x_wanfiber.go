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
	NewOpticalSignalLevel          int64    `xml:"NewOpticalSignalLevel"`
	NewLowerOpticalThreshold       int64    `xml:"NewLowerOpticalThreshold"`
	NewUpperOpticalThreshold       int64    `xml:"NewUpperOpticalThreshold"`
	NewTransmitOpticalLevel        int64    `xml:"NewTransmitOpticalLevel"`
	NewLowerTransmitPowerThreshold int64    `xml:"NewLowerTransmitPowerThreshold"`
	NewUpperTransmitPowerThreshold int64    `xml:"NewUpperTransmitPowerThreshold"`
	NewSFPVendor                   string   `xml:"NewSFPVendor"`
	NewSFPPartNumber               string   `xml:"NewSFPPartNumber"`
	NewSFPSerialNumber             string   `xml:"NewSFPSerialNumber"`
	NewSFPType                     int64    `xml:"NewSFPType"`
	NewTXWaveLength                uint64   `xml:"NewTXWaveLength"`
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
	NewONUId        uint64   `xml:"NewONUId"`
	NewUNIType      string   `xml:"NewUNIType"`
	NewGEMPortCount uint64   `xml:"NewGEMPortCount"`
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
	NewBytesSent            uint64   `xml:"NewBytesSent"`
	NewBytesReceived        uint64   `xml:"NewBytesReceived"`
	NewPacketsSent          uint64   `xml:"NewPacketsSent"`
	NewPacketsReceived      uint64   `xml:"NewPacketsReceived"`
	NewPacketErrorsSent     uint64   `xml:"NewPacketErrorsSent"`
	NewPacketErrorsReceived uint64   `xml:"NewPacketErrorsReceived"`
	NewPacketsMulticast     uint64   `xml:"NewPacketsMulticast"`
	NewConnectionRateDown   uint64   `xml:"NewConnectionRateDown"`
	NewConnectionRateUp     uint64   `xml:"NewConnectionRateUp"`
	NewBestTrainState       uint64   `xml:"NewBestTrainState"`
	NewResyncs              uint64   `xml:"NewResyncs"`
	NewMinutesInShowtime    uint64   `xml:"NewMinutesInShowtime"`
}

func (client *ServiceClient) GetStatistics(out *GetStatisticsResponse) error {
	in := &GetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetStatistics", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
