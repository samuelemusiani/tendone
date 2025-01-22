package tendone

import (
	"bytes"
	"encoding/json"
	"net/http"
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

type WifiClientNumWrap struct {
	WifiClientNum WifiClientNum `json:"wifiClientNum"`
}

type WifiClientNum struct {
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

func (s *Session) WifiClientNum() (*WifiClientNum, error) {
	rbody, err := json.Marshal(WifiClientNumWrap{})
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

	var wc WifiClientNumWrap
	err = json.NewDecoder(resp.Body).Decode(&wc)
	if err != nil {
		return nil, err
	}

	return &wc.WifiClientNum, nil
}

type RadioType string

var Radio2_4G RadioType = "2.4G"
var Radio5G RadioType = "5G"

type WifiRadioStatusRequest struct {
	Radio RadioType `json:"radio"`
}

type WifiRadioStatusRequestWrap struct {
	WifiRadioStatus WifiRadioStatusRequest `json:"wifiRadioStatus"`
}

type WifiRadioStatusResponse struct {
	NetMode     string `json:"netMode"`
	Channel     string `json:"channel"`
	WifiRadioEn bool   `json:"wifiRadioEn"`
}

type WifiRadioStatusResponseWrap struct {
	WifiRadioStatus WifiRadioStatusResponse `json:"wifiRadioStatus"`
}

func (s *Session) WifiRadioStatus(radio RadioType) (*WifiRadioStatusResponse, error) {
	rbody, err := json.Marshal(WifiRadioStatusRequestWrap{WifiRadioStatusRequest{Radio: radio}})
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

	var wr WifiRadioStatusResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return nil, err
	}

	return &wr.WifiRadioStatus, nil
}

type WifiSSIDListRequestWrap struct {
	WifiSSIDList WifiSSIDListRequest `json:"wifiSsidList"`
}

type WifiSSIDListRequest struct {
	Radio RadioType `json:"radio"`
}

type WifiSSIDListResponseWrap struct {
	WifiSSIDList WifiSSIDListResponse `json:"wifiSsidList"`
}

type WifiSSIDListResponse []SSID

type SSID struct {
	SSID     string `json:"ssid"`
	Enabled  bool   `json:"wifiSsidEn"`
	Mac      string `json:"mac"`
	Security string `json:"security"`
}

func (s *Session) WifiSSIDList(radio RadioType) (*WifiSSIDListResponse, error) {
	rbody, err := json.Marshal(WifiSSIDListRequestWrap{WifiSSIDListRequest{Radio: radio}})

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

	var ws WifiSSIDListResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&ws)
	if err != nil {
		return nil, err
	}

	return &ws.WifiSSIDList, nil
}
