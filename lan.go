package tendone

import (
	"encoding/json"
	"errors"
)

type lanConfigGetRequestWrap struct {
	Config interface{} `json:"lanCfgGet"`
}

type lanConfigGetResponseWrap struct {
	Config LanConfig `json:"lanCfgGet"`
}

type EthMode string

var ErrInvalidEthMode = errors.New("Invalid EthMode")

const (
	EthModeAuto EthMode = "auto"
	EthMode10M  EthMode = "10M"
)

func isValidEthMode(em EthMode) bool {
	switch em {
	case EthModeAuto, EthMode10M:
		return true
	}
	return false
}

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

type lanConfigSetRequestWrap struct {
	Config LanConfig `json:"lanCfgSet"`
}

type lanConfigSetResponseWrap struct {
	Resp string `json:"lanCfgSet"`
}

func (s *Session) LanConfigGet() (*LanConfig, error) {
	rbody, err := json.Marshal(lanConfigGetRequestWrap{})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var lcr lanConfigGetResponseWrap
	err = json.Unmarshal(resp, &lcr)
	if err != nil {
		return nil, err
	}

	return &lcr.Config, nil
}

func (s *Session) LanConfigSet(lc *LanConfig) (bool, error) {
	rbody, err := json.Marshal(lanConfigSetRequestWrap{Config: *lc})
	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var lcs lanConfigSetResponseWrap
	err = json.Unmarshal(resp, &lcs)
	if err != nil {
		return false, err
	}

	if lcs.Resp == "ok" {
		return true, nil
	}

	return false, nil
}
