package tendone

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type rebootRequestWrap struct {
	Req RebootRequest `json:"sysReboot"`
}

type RebootRequest string

func (s *Session) Reboot() error {
	rbody, err := json.Marshal(rebootRequestWrap{
		Req: "",
	})
	if err != nil {
		return err
	}

	// No Response for the reboot command
	_, err = fetch(s, rbody)
	return err
}

func (s *Session) Backup(path string) error {
	urlPath := "/cgi-bin/DownloadCfg/APCfm.cfg"

	if _, err := os.Stat(path); err == nil {
		return os.ErrExist
	} else if !os.IsNotExist(err) {
		return err
	}

	req, err := http.NewRequest("GET", s.uri+urlPath, nil)
	if err != nil {
		return err
	}

	for _, c := range s.cookies {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

type ledCtrlRequestWrap struct {
	Req interface{} `json:"sysLedCtrlGet"`
}

type ledCtrlGetResponseWrap struct {
	Resp LedCtrlGetResponse `json:"sysLedCtrlGet"`
}

type LedCtrlGetResponse struct {
	Enable string `json:"enable"`
}

func (s *Session) LedCtrlGet() (bool, error) {
	rbody, err := json.Marshal(ledCtrlRequestWrap{})
	if err != nil {
		return false, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var lcr ledCtrlGetResponseWrap
	err = json.Unmarshal(resp, &lcr)
	if err != nil {
		return false, err
	}

	return lcr.Resp.Enable == "on", nil
}

type ledCtrlSetRequestWrap struct {
	Req LedCtrlSetRequest `json:"sysLedCtrlSet"`
}

type LedCtrlSetRequest struct {
	Enable string `json:"enable"`
}

type ledCtrlSetResponseWrap struct {
	Resp string `json:"enable"`
}

func (s *Session) LedCtrlSet(enable bool) (bool, error) {
	var enableStr string
	if enable {
		enableStr = "on"
	} else {
		enableStr = "off"
	}

	rbody, err := json.Marshal(ledCtrlSetRequestWrap{
		Req: LedCtrlSetRequest{
			Enable: enableStr,
		},
	})
	if err != nil {
		return false, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var lcr ledCtrlSetResponseWrap
	err = json.Unmarshal(resp, &lcr)
	if err != nil {
		return false, err
	}

	return lcr.Resp == "ok", nil
}

type logsRequestWrap struct {
	Req interface{} `json:"sysLogGet"`
}

type logsResponseWrap struct {
	Resp []LogsResponse `json:"sysLogGet"`
}

type LogsResponse struct {
	Index int    `json:"index"`
	Time  string `json:"time"`
	Type  string `json:"type"`
	Info  string `json:"info"`
}

func (s *Session) Logs() ([]LogsResponse, error) {
	rbody, err := json.Marshal(logsRequestWrap{})
	if err != nil {
		return nil, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return nil, err
	}

	var lr logsResponseWrap
	err = json.Unmarshal(resp, &lr)
	if err != nil {
		return nil, err
	}

	return lr.Resp, nil
}

type logsClearRequestWrap struct {
	Req interface{} `json:"sysLogClear"`
}

func (s *Session) LogsClear() error {
	rbody, err := json.Marshal(logsClearRequestWrap{})
	if err != nil {
		return err
	}

	_, err = fetch(s, rbody)
	return err
}
