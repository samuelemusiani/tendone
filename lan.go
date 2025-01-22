package tendone

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type LanConfigGetRequestWrap struct {
	Config LanConfigGetRequest `json:"lanCfgGet"`
}

type LanConfigGetRequest struct{}

type LanConfigGetResponseWrap struct {
	Config LanConfig `json:"lanCfgGet"`
}

type EthMode string

var EthModeAuto EthMode = "auto"
var EthMode10M EthMode = "10M"

type LanConfig struct {
	Type       string  `json:"lanType"`
	DeviceName string  `json:"deviceName"`
	MAC        string  `json:"lanMac"`
	IP         string  `json:"lanIp"`
	Mask       string  `json:"lanMask"`
	Gateway    string  `json:"lanGw"`
	DNS0       string  `json:"preDns"`
	DNS1       string  `json:"altDns"`
	EthMode    EthMode `json:"ethMode"`
}

type LanConfigSetRequestWrap struct {
	Config LanConfig `json:"lanCfgSet"`
}

type LanConfigSetResponseWrap struct {
	Resp string `json:"lanCfgSet"`
}

func (s *Session) LanConfigGet() (*LanConfig, error) {
	rbody, err := json.Marshal(LanConfigGetRequestWrap{})
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

	var lcr LanConfigGetResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&lcr)
	if err != nil {
		return nil, err
	}

	return &lcr.Config, nil
}

func (s *Session) LanConfigSet(lc *LanConfig) (bool, error) {
	rbody, err := json.Marshal(LanConfigSetRequestWrap{Config: *lc})
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

	var lcs LanConfigSetResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&lcs)
	if err != nil {
		return false, err
	}

	if lcs.Resp == "ok" {
		return true, nil
	}

	return false, nil
}
