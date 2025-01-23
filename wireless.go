package tendone

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type basicGetIndoorRequestWrap struct {
	BasicGetIndoor BasicGetIndoorRequest `json:"wifiBasicGetIndoor"`
}

type BasicGetIndoorRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`
}

type basicGetIndoorResponseWrap struct {
	BasicGetIndoor BasicGetIndoorResponse `json:"wifiBasicGetIndoor"`
}

type BasicGetIndoorResponse struct {
	SSID         string `json:"ssid"`
	SSIDEncode   string `json:"ssidEncode"`
	MaxClientNum string `json:"maxClientNum"`
	Enabled      bool   `json:"ssidEn"`
	// If the SSID should be broadcasted
	Broadcast        bool   `json:"broadcastSsid"`
	Guest            bool   `json:"isGuest"`
	IsolateClient    bool   `json:"staIsolate"`
	IsolateSSID      bool   `json:"ssidIsolate"`
	WMF              bool   `json:"wmf"`
	AvailableClients string `json:"avalidClientNum"`
	ScheduleEN       bool   `json:"scheduleEn"`
}

func (s *Session) BasicGetIndoor(radio RadioType, ssidIndex string) (*BasicGetIndoorResponse, error) {
	rbody, err := json.Marshal(basicGetIndoorRequestWrap{BasicGetIndoorRequest{Radio: radio, SSIDIndex: ssidIndex}})

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

	var wr basicGetIndoorResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return nil, err
	}

	return &wr.BasicGetIndoor, nil
}

type basicSetIndoorRequestWrap struct {
	BasicSetIndoor BasicSetIndoorRequest `json:"wifiBasicSetIndoor"`
}

type BasicSetIndoorRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`

	SSID       string `json:"ssid"`
	SSIDEncode string `json:"ssidEncode"`
	Enabled    bool   `json:"ssidEn"`
	// If the SSID should be broadcasted
	Broadcast     bool   `json:"broadcastSsid"`
	Guest         bool   `json:"isGuest"`
	IsolateClient bool   `json:"staIsolate"`
	IsolateSSID   bool   `json:"ssidIsolate"`
	WMF           bool   `json:"wmf"`
	MaxClientNum  string `json:"maxClientNum"`
}

type basicSetIndoorResponseWrap struct {
	SetIndoorResponse BasicSetIndoorResponse `json:"wifiBasicSetIndoor"`
}

type BasicSetIndoorResponse string

// BasicSetIndoor sets the indoor settings for a given SSID.
func (s *Session) BasicSetIndoor(bs BasicSetIndoorRequest) (bool, error) {
	rbody, err := json.Marshal(basicSetIndoorRequestWrap{bs})

	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(rbody))
	if err != nil {
		return false, err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	var wr basicSetIndoorResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return false, err
	}

	if wr.SetIndoorResponse != "ok" {
		return false, errors.New("Failed to set indoor")
	}

	return true, nil
}
