package tendone

import (
	"encoding/json"
	"errors"
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

	resp, err := fetch(s, rbody)
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

	resp, err := fetch(s, rbody)
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

type ssidSecurityRequestWrap struct {
	Req SSIDSecurityRequest `json:"apSecurityGet"`
}

type SSIDSecurityRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`
}

type ssidSecurityResponseWrap struct {
	Resp SSIDSecurityResponse `json:"apSecurityGet"`
}

type SSIDSecurityType string

var (
	SSIDSecurityTypeNone           SSIDSecurityType = "none"
	SSIDSecurityTypeWep            SSIDSecurityType = "wep"
	SSIDSecurityTypeWpaPsk         SSIDSecurityType = "wpa-psk"
	SSIDSecurityTypeWpa2Psk        SSIDSecurityType = "wpa2-psk"
	SSIDSecurityTypeWpaWpa2Psk     SSIDSecurityType = "mixed wpa/wpa2-psk"
	SSIDSecurityTypeWpa            SSIDSecurityType = "wpa"
	SSIDSecurityTypeWpa2           SSIDSecurityType = "wpa2"
	SSIDSecurityTypeWpa3Sae        SSIDSecurityType = "wpa3sae"
	SSIDSecurityTypeWpa3SaeWpa2Psk SSIDSecurityType = "wpa3saewpa2psk"
)

type WepSecurityMode string

var (
	WepSecurityModeOpen   WepSecurityMode = "open"
	WepSecurityModeShared WepSecurityMode = "shared"
)

type WepKeyFormat string

var (
	WepKeyFormatAscii WepKeyFormat = "ascii"
	WepKeyFormatHex   WepKeyFormat = "hex"
)

type SSIDSecurityResponse struct {
	Radio         RadioType        `json:"radio"`
	SSIDIndex     string           `json:"ssidIndex"`
	Type          SSIDSecurityType `json:"secType"`
	WepAuth       WepSecurityMode  `json:"wepAuth"`
	WepDefaultKey string           `json:"wepDefaultKey"`
	WepKey1       string           `json:"wepKey1"`
	WepKeyFormat1 WepKeyFormat     `json:"wepKeyType1"`
	WepKey2       string           `json:"wepKey2"`
	WepKeyFormat2 WepKeyFormat     `json:"wepKeyType2"`
	WepKey3       string           `json:"wepKey3"`
	WepKeyFormat3 WepKeyFormat     `json:"wepKeyType3"`
	WepKey4       string           `json:"wepKey4"`
	WepKeyFormat4 WepKeyFormat     `json:"wepKeyType4"`
	// I don't know what this is. The default value is 'aes'.
	// I don't think you can change it from the web interface
	WpaPskAuth        string `json:"wpapskAuth"`
	RadiusPwdInterval string `json:"radiusPwdInterval"`
	WpaPskPwdInterval string `json:"wpapskPwdInterval"`
	WpaPskPwd         string `json:"wpapskPwd"`
	RadiusKey         string `json:"radiusKey"`
	RadiusIP          string `json:"radiusIp"`
	RadiusPort        string `json:"radiusPort"`
	// I don't know what this is. The default value is 'aes'.
	// I don't think you can change it from the web interface
	RadiusAuth string `json:"radiusAuth"`
}

func (s *Session) SSIDSecurityGet(radio RadioType, ssidIndex string) (*SSIDSecurityResponse, error) {
	rbody, err := json.Marshal(ssidSecurityRequestWrap{SSIDSecurityRequest{Radio: radio, SSIDIndex: ssidIndex}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr ssidSecurityResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return nil, err
	}

	return &wr.Resp, nil
}

type ssidSecuritySetRequestWrap struct {
	Req SSIDSecurityResponse `json:"apSecuritySet"`
}

type ssidSecuritySetResponseWrap struct {
	Resp string `json:"apSecuritySet"`
}

func (s *Session) SSIDSecuritySet(ssr SSIDSecurityResponse) (bool, error) {
	rbody, err := json.Marshal(ssidSecuritySetRequestWrap{Req: ssr})
	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var wr ssidSecuritySetResponseWrap
	err = json.NewDecoder(resp.Body).Decode(&wr)
	if err != nil {
		return false, err
	}

	if wr.Resp != "ok" {
		return false, errors.New("Failed to set security")
	}

	return true, nil
}
