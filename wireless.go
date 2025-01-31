package tendone

import (
	"encoding/json"
	"errors"
)

type basicGetIndoorRequestWrap struct {
	BasicGetIndoor basicGetIndoorRequest `json:"wifiBasicGetIndoor"`
}

type basicGetIndoorRequest struct {
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
	rbody, err := json.Marshal(basicGetIndoorRequestWrap{basicGetIndoorRequest{Radio: radio, SSIDIndex: ssidIndex}})

	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr basicGetIndoorResponseWrap
	err = json.Unmarshal(resp, &wr)
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
	SetIndoorResponse basicSetIndoorResponse `json:"wifiBasicSetIndoor"`
}

type basicSetIndoorResponse string

// BasicSetIndoor sets the indoor settings for a given SSID.
func (s *Session) BasicSetIndoor(bs BasicSetIndoorRequest) (bool, error) {
	rbody, err := json.Marshal(basicSetIndoorRequestWrap{bs})

	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var wr basicSetIndoorResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return false, err
	}

	if wr.SetIndoorResponse != "ok" {
		return false, errors.New("Failed to set indoor")
	}

	return true, nil
}

type ssidSecurityRequestWrap struct {
	Req ssidSecurityRequest `json:"apSecurityGet"`
}

type ssidSecurityRequest struct {
	Radio     RadioType `json:"radio"`
	SSIDIndex string    `json:"ssidIndex"`
}

type ssidSecurityResponseWrap struct {
	Resp SSIDSecurityResponse `json:"apSecurityGet"`
}

type SSIDSecurityType string

const (
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

const (
	WepSecurityModeOpen   WepSecurityMode = "open"
	WepSecurityModeShared WepSecurityMode = "shared"
)

type WepKeyFormat string

const (
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
	rbody, err := json.Marshal(ssidSecurityRequestWrap{ssidSecurityRequest{Radio: radio, SSIDIndex: ssidIndex}})
	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr ssidSecurityResponseWrap
	err = json.Unmarshal(resp, &wr)
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
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return false, err
	}

	if wr.Resp != "ok" {
		return false, errors.New("Failed to set security")
	}

	return true, nil
}

// TODO: RF Settings

// TODO: RF Optimization

type channelAnalyseRequestWrap struct {
	Req channelAnalyseRequest `json:"wifiChannelAnalyse"`
}

type ChannelList string

const (
	ChannelList2_4G ChannelList = "1,2,3,4,5,6,7,8,9,10,11,12,13"
	ChannelList5G   ChannelList = "36,40,44,48,149,153,157,161"
)

type channelAnalyseRequest struct {
	Channels ChannelList `json:"channel"`
	Radio    RadioType   `json:"radio"`
}

type channelAnalyseResponseWrap struct {
	Resp ChannelAnalyseResponse `json:"wifiChannelAnalyse"`
}

type ChannelAnalyseResponse struct {
	TotalSSID string `json:"channelNum"`
	Percent   string `json:"channelPercent"`
}

// In the webui this is the 'Frequency Analysis' even though it's called
// 'Channel Analysis' in the API. For the Channel Scan, see [Session.APAnalyse]
func (s *Session) ChannelAnalyse(radio RadioType) (*ChannelAnalyseResponse, error) {
	var channels ChannelList
	switch radio {
	case Radio2_4G:
		channels = ChannelList2_4G
	case Radio5G:
		channels = ChannelList5G
	default:
		return nil, ErrInvalidRadio
	}

	rbody, err := json.Marshal(channelAnalyseRequestWrap{channelAnalyseRequest{Channels: channels, Radio: radio}})
	if err != nil {
		return nil, err
	}

	resp, err := fetch(s, rbody)

	var wr channelAnalyseResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return nil, err
	}

	return &wr.Resp, nil
}

type apAnalyseRequestWrap struct {
	Req apAnalyseRequest `json:"wifiApAnalyse"`
}

type apAnalyseRequest struct {
	Radio RadioType `json:"radio"`
}

type apAnalyseResponseWrap struct {
	Resp []ApAnalyseResponse `json:"wifiApAnalyse"`
}

type ApAnalyseResponse struct {
	Index     int    `json:"index"`
	SSID      string `json:"ssid"`
	Channel   int    `json:"channel"`
	SSIDEcode string `json:"ssidEncode"`
	Mac       string `json:"mac"`
	Security  string `json:"secType"`
	Bandwidth string `json:"bandwidth"`
	Signal    int    `json:"signal"`
	Netmode   string `json:"netmode"`
}

func (s *Session) APAnalyse(radio RadioType) ([]ApAnalyseResponse, error) {
	if !isValidRadio(radio) {
		return nil, ErrInvalidRadio
	}

	rbody, err := json.Marshal(apAnalyseRequestWrap{apAnalyseRequest{Radio: radio}})
	if err != nil {
		return nil, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr apAnalyseResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return nil, err
	}

	return wr.Resp, nil
}

type wmmGetRequestWrap struct {
	Req wmmGetRequest `json:"wifiWmmGet"`
}

type wmmGetRequest struct {
	Radio RadioType `json:"radio"`
}

type wmmGetResponseWrap struct {
	Resp WmmGetResponse `json:"wifiWmmGet"`
}

type WmmConfig struct {
	Be string `json:"be"`
	Bk string `json:"bk"`
	Vi string `json:"vi"`
	Vo string `json:"vo"`
}

type WmmStaConfig struct {
	Be string `json:"be"`
	Bk string `json:"bk"`
	Vi string `json:"vi"`
	Vo string `json:"vo"`
}

type WmmMode string

const (
	// Optimize for scenarios with 1-10 users
	WmmModeNormal WmmMode = "normal"
	// Optimize for scenarios with more than 10 users
	WmmModeHigh WmmMode = "high"
	// Custom settings
	WmmModeCustom WmmMode = "custom"
)

func IsValidWmmMode(m WmmMode) bool {
	switch m {
	case WmmModeHigh, WmmModeNormal, WmmModeCustom:
		return true
	default:
		return false
	}
}

var ErrInvalidWmmMode = errors.New("Invalid WMM mode")

type WmmGetResponse struct {
	Enabled   bool         `json:"wmmEn"`
	NoAck     bool         `json:"noAck"`
	Mode      WmmMode      `json:"wmmMode"`
	Config    WmmConfig    `json:"wmmConfig"`
	StaConfig WmmStaConfig `json:"wmmStaConfig"`
}

type wmmSetRequestWrap struct {
	Req WmmSetRequest `json:"wifiWmmSet"`
}

type WmmSetRequest struct {
	Radio     RadioType    `json:"radio"`
	NoAck     string       `json:"noAck"`
	Mode      WmmMode      `json:"wmmMode"`
	Enabled   bool         `json:"wmmEn"`
	Config    WmmConfig    `json:"wmmConfig"`
	StaConfig WmmStaConfig `json:"wmmStaConfig"`
}

type wmmSetResponseWrap struct {
	Resp string `json:"wifiWmmSet"`
}

func (r WmmGetResponse) IntoSetRequest(radio RadioType) WmmSetRequest {
	var ack string
	switch r.NoAck {
	case true:
		ack = "true"
	case false:
		ack = "false"
	}
	return WmmSetRequest{
		Radio:     radio,
		NoAck:     ack,
		Mode:      r.Mode,
		Enabled:   r.Enabled,
		Config:    r.Config,
		StaConfig: r.StaConfig,
	}
}

func (s *Session) WmmGet(radio RadioType) (*WmmGetResponse, error) {
	if !isValidRadio(radio) {
		return nil, ErrInvalidRadio
	}

	rbody, err := json.Marshal(wmmGetRequestWrap{wmmGetRequest{Radio: radio}})
	if err != nil {
		return nil, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var wr wmmGetResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return nil, err
	}

	return &wr.Resp, nil
}

func (s *Session) WmmSet(wsr WmmSetRequest) (bool, error) {
	if !isValidRadio(wsr.Radio) {
		return false, ErrInvalidRadio
	}

	if wsr.NoAck != "true" && wsr.NoAck != "false" {
		return false, errors.New("Invalid NoAck value")
	}

	if !IsValidWmmMode(wsr.Mode) {
		return false, ErrInvalidWmmMode
	}

	rbody, err := json.Marshal(wmmSetRequestWrap{Req: wsr})
	if err != nil {
		return false, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var wr wmmSetResponseWrap
	err = json.Unmarshal(resp, &wr)
	if err != nil {
		return false, err
	}

	if wr.Resp != "ok" {
		return false, errors.New("Failed to set WMM: " + wr.Resp)
	}

	return true, nil
}
