package tendone

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type rebootRequestWrap struct {
	RebootRequest RebootRequest `json:"sysReboot"`
}

type RebootRequest string

func (s *Session) Reboot() error {
	rbody, err := json.Marshal(rebootRequestWrap{
		RebootRequest: "",
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
