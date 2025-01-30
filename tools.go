package tendone

import "encoding/json"

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
