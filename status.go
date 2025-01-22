package tendone

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type SysStatusWrap struct {
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

type LanStatusWrap struct {
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

type ClientNumWrap struct {
	ClientNum ClientNum `json:"wifiClientNum"`
}

type ClientNum struct {
	Num string `json:"clientNum"`
}

func (s *Session) GetSysStatus() (*SysStatus, error) {
	rbody, err := json.Marshal(SysStatusWrap{})

	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ss SysStatusWrap
	err = json.NewDecoder(resp.Body).Decode(&ss)
	if err != nil {
		return nil, err
	}

	return &ss.SysStatus, nil
}

func (s *Session) LanStatus() (*LanStatus, error) {
	rbody, err := json.Marshal(LanStatusWrap{})
	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ls LanStatusWrap
	err = json.NewDecoder(resp.Body).Decode(&ls)
	if err != nil {
		return nil, err
	}

	return &ls.LanStatus, nil
}

func (s *Session) ClientNum() (*ClientNum, error) {
	rbody, err := json.Marshal(ClientNumWrap{})
	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var wc ClientNumWrap
	err = json.NewDecoder(resp.Body).Decode(&wc)
	if err != nil {
		return nil, err
	}

	return &wc.ClientNum, nil
}

type RadioType string

var Radio2_4G RadioType = "2.4G"
var Radio5G RadioType = "5G"

type RadioStatusRequest struct {
	Radio RadioType `json:"radio"`
}

type RadioStatusRequestWrap struct {
	RadioStatus RadioStatusRequest `json:"wifiRadioStatus"`
}

type RadioStatusResponse struct {
	NetMode string `json:"netMode"`
	Channel string `json:"channel"`
	RadioEn bool   `json:"wifiRadioEn"`
}

type RadioStatusResponseWrap struct {
	RadioStatus RadioStatusResponse `json:"wifiRadioStatus"`
}

func (s *Session) RadioStatus(radio RadioType) (*RadioStatusResponse, error) {
	rbody, err := json.Marshal(RadioStatusRequestWrap{RadioStatusRequest{Radio: radio}})
	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var wr RadioStatusResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return nil, err
	}

	return &wr.RadioStatus, nil
}

type SSIDListRequestWrap struct {
	SSIDList SSIDListRequest `json:"wifiSsidList"`
}

type SSIDListRequest struct {
	Radio RadioType `json:"radio"`
}

type SSIDListResponseWrap struct {
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
	rbody, err := json.Marshal(SSIDListRequestWrap{SSIDListRequest{Radio: radio}})

	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ws SSIDListResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&ws)
	if err != nil {
		return nil, err
	}

	return &ws.SSIDList, nil
}

type TrafficRequestWrap struct {
	Traffic TrafficRequest `json:"wifiTraffic"`
}

type TrafficRequest struct {
	Radio RadioType `json:"radio"`
}

type TrafficResponseWrap struct {
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
	rbody, err := json.Marshal(TrafficRequestWrap{TrafficRequest{Radio: radio}})

	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var tr TrafficResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return nil, err
	}

	return &tr.Traffic, nil
}

type ClientListRequestWrap struct {
	ClientListRequest ClientListRequest `json:"wifiClientList"`
}

type ClientListRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`
}

type ClientListResponseWrap struct {
	ClientListResponse ClientListResponse `json:"wifiClientList"`
}

type ClientListResponse []WifiClient

type WifiClient struct {
	Select      string `json:"select"`
	Index       string `json:"index"`
	IP          string `json:"ip"`
	MAC         string `json:"mac"`
	ConnectTime string `json:"connectTime"`
	TxRate      string `json:"txRate"`
	RxRate      string `json:"rxRate"`
	SignNoise   string `json:"signNoise"`
	CCQ         string `json:"ccq"`
	OsEnable    string `json:"osEnable"`
	OsType      string `json:"osType"`
}

func (s *Session) ClientList(radio RadioType, ssidIndex int) (*ClientListResponse, error) {
	rbody, err := json.Marshal(ClientListRequestWrap{ClientListRequest{Radio: radio, SSIDIndex: strconv.Itoa(ssidIndex)}})

	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return nil, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var cl ClientListResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&cl)
	if err != nil {
		return nil, err
	}

	return &cl.ClientListResponse, nil
}
