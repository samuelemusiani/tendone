package tendone

import (
	"encoding/json"
	"errors"
	"strconv"
)

type sysStatusWrap struct {
	SysStatus SysStatus `json:"sysStatus"`
}

type SysStatus struct {
	RunningTime     string `json:"runningTime"`
	DeviceName      string `json:"deviceName"`
	SoftwareVersion string `json:"softwareVersion"`
	SysTime         string `json:"sysTime"`
	CPU             string `json:"cpu"`
	RAM             string `json:"ram"`
	HardwareVersion string `json:"hardwareVersion"`
	SerialNumber    string `json:"serialNumber"`
	UcmanageEn      bool   `json:"ucmanageEn"`
	Mode            string `json:"mode"`
	BridgeStatus    bool   `json:"bridgeStatus"`
}

type lanStatusWrap struct {
	LanStatus LanStatus `json:"lanStatus"`
}

type LanStatus struct {
	MAC     string `json:"lanMac"`
	IP      string `json:"lanIp"`
	Mask    string `json:"lanMask"`
	Gateway string `json:"lanGw"`
	DNS0    string `json:"preDns"`
	DNS1    string `json:"altDns"`
}

type clientNumWrap struct {
	ClientNum ClientNum `json:"wifiClientNum"`
}

type ClientNum struct {
	Num string `json:"clientNum"`
}

func (s *Session) GetSysStatus() (*SysStatus, error) {
	rbody, err := json.Marshal(sysStatusWrap{})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var ss sysStatusWrap
	err = json.Unmarshal(resp, &ss)
	if err != nil {
		return nil, err
	}

	return &ss.SysStatus, nil
}

func (s *Session) LanStatus() (*LanStatus, error) {
	rbody, err := json.Marshal(lanStatusWrap{})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var ls lanStatusWrap
	err = json.Unmarshal(resp, &ls)
	if err != nil {
		return nil, err
	}

	return &ls.LanStatus, nil
}

func (s *Session) ClientNum() (*ClientNum, error) {
	rbody, err := json.Marshal(clientNumWrap{})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wc clientNumWrap
	err = json.Unmarshal(resp, &wc)
	if err != nil {
		return nil, err
	}

	return &wc.ClientNum, nil
}

var ErrInvalidRadio = errors.New("Invalid radio type")

type RadioType string

const (
	Radio2_4G RadioType = "2.4G"
	Radio5G   RadioType = "5G"
)

func isValidRadio(radio RadioType) bool {
	switch radio {
	case Radio2_4G, Radio5G:
		return true
	default:
		return false
	}
}

type radioStatusRequest struct {
	Radio RadioType `json:"radio"`
}

type radioStatusRequestWrap struct {
	RadioStatus radioStatusRequest `json:"wifiRadioStatus"`
}

type RadioStatusResponse struct {
	NetMode string `json:"netMode"`
	Channel string `json:"channel"`
	RadioEn bool   `json:"wifiRadioEn"`
}

type radioStatusResponseWrap struct {
	RadioStatus RadioStatusResponse `json:"wifiRadioStatus"`
}

func (s *Session) RadioStatus(radio RadioType) (*RadioStatusResponse, error) {
	if !isValidRadio(radio) {
		return nil, ErrInvalidRadio
	}

	rbody, err := json.Marshal(radioStatusRequestWrap{radioStatusRequest{Radio: radio}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr radioStatusResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return nil, err
	}

	return &wr.RadioStatus, nil
}

type ssidListRequestWrap struct {
	SSIDList ssidListRequest `json:"wifiSsidList"`
}

type ssidListRequest struct {
	Radio RadioType `json:"radio"`
}

type ssidListResponseWrap struct {
	SSIDList SSIDListResponse `json:"wifiSsidList"`
}

type SSIDListResponse []SSID

type SSID struct {
	SSID     string `json:"ssid"`
	Enabled  bool   `json:"wifiSsidEn"`
	Mac      string `json:"mac"`
	Security string `json:"security"`
}

func (s *Session) SSIDList(radio RadioType) (*SSIDListResponse, error) {
	rbody, err := json.Marshal(ssidListRequestWrap{ssidListRequest{Radio: radio}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var ws ssidListResponseWrap
	err = json.Unmarshal(resp, &ws)
	if err != nil {
		return nil, err
	}

	return &ws.SSIDList, nil
}

type trafficRequestWrap struct {
	Traffic trafficRequest `json:"wifiTraffic"`
}

type trafficRequest struct {
	Radio RadioType `json:"radio"`
}

type trafficResponseWrap struct {
	Traffic TrafficResponse `json:"wifiTraffic"`
}

type TrafficResponse []SSIDTraffic

type SSIDTraffic struct {
	SSID string `json:"ssid"`
	// Mesurements are in MB
	RxTraffic   string `json:"rxTraffic"`
	RxPacketNum string `json:"rxPacketNum"`
	// Mesurements are in MB
	TxTraffic   string `json:"txTraffic"`
	TxPacketNum string `json:"txPacketNum"`
}

func (s *Session) Traffic(radio RadioType) (*TrafficResponse, error) {
	if !isValidRadio(radio) {
		return nil, ErrInvalidRadio
	}

	rbody, err := json.Marshal(trafficRequestWrap{trafficRequest{Radio: radio}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var tr trafficResponseWrap
	err = json.Unmarshal(resp, &tr)
	if err != nil {
		return nil, err
	}

	return &tr.Traffic, nil
}

type clientListRequestWrap struct {
	ClientListRequest clientListRequest `json:"wifiClientList"`
}

type clientListRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`
}

type clientListResponseWrap struct {
	ClientListResponse ClientListResponse `json:"wifiClientList"`
}

type ClientListResponse []WifiClient

type WifiClient struct {
	Select string `json:"select"`
	Index  string `json:"index"`
	IP     string `json:"ip"`
	MAC    string `json:"mac"`
	// Mesurements are in seconds
	ConnectTime string `json:"connectTime"`
	// Mesurements are in Mbps
	TxRate string `json:"txRate"`
	// Mesurements are in Mbps
	RxRate    string `json:"rxRate"`
	SignNoise string `json:"signNoise"`
	CCQ       string `json:"ccq"`
	OsEnable  string `json:"osEnable"`
	OsType    string `json:"osType"`
}

func (s *Session) ClientList(radio RadioType, ssidIndex int) (*ClientListResponse, error) {
	rbody, err := json.Marshal(clientListRequestWrap{clientListRequest{Radio: radio, SSIDIndex: strconv.Itoa(ssidIndex)}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var cl clientListResponseWrap
	err = json.Unmarshal(resp, &cl)
	if err != nil {
		return nil, err
	}

	return &cl.ClientListResponse, nil
}
