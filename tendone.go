/*
Package tendone provides a simple library to interact with a Tenda i27 access
point. This could be used to automate some tasks, monitor the device or to
create a new and better web interface with simpler and more powerful apis.
This could also be used to automate changes across multiple devices, without
exposing them to the internet with the Tenda could service.

# Basic Usage

To correctly use this library you need to create a new session with the Tenda i27
thorugh the [tendone.NewSession] function. This function will return a new session
that can be used to interact with the device.

	// Create a new session
	s := tendone.NewSession("http://tenda-AP.int")
	// Login the session
	lg, _ := s.Login("admin", "password")
	// Execute the desired functions, es. get the status of the device
	st, err := s.GetSysStatus()
	// Logout the session
	lg, err = s.Logout()

The code does not check for errors, but you should always do it.
*/
package tendone

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const MODULES_PATH = "/goform/modules"

// Session is an istance of a session with a Tenda i27 access point. At the moment
// ONLY ONE session is supported for acess point. This is NOT a library limitation,
// but rather a limitation of the access point itself. To create a new session use
// the [tendone.NewSession] function.
type Session struct {
	uri     string
	cookies []*http.Cookie
}

// NewSession creates a new session with the Tenda i27 access point at the given
// uri. The uri should be http://<ip> or http://<hostname>
func NewSession(uri string) *Session {
	return &Session{uri: uri}
}

// GetURI returns the uri of the session
func (s *Session) GetURI() string {
	return s.uri
}

var ErrLoggedOut = errors.New("Session is logged out")
var ErrMSGNotValid = errors.New("Message not valid")

type errCodeWrap struct {
	ErrCode string `json:"errCode"`
}

type notValidWrap struct {
	NotValid string `json:"resp"`
}

// As almost all the functions in this library are the same, this function is
// used to fetch the data from the access point avoiding code duplication.
// Returns the response body and an error if any.
func fetch(s *Session, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", s.uri+MODULES_PATH, bytes.NewReader(body))
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

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errCode errCodeWrap
	err = json.Unmarshal(rbody, &errCode)
	if err != nil {
		return nil, err
	}

	if len(errCode.ErrCode) > 0 {
		if errCode.ErrCode == "logout" {
			return nil, ErrLoggedOut
		} else {
			return nil, errors.New("Error code: " + errCode.ErrCode)
		}
	}

	var notValid notValidWrap
	err = json.Unmarshal(rbody, &notValid)
	if err != nil {
		return nil, err
	}

	if len(notValid.NotValid) > 0 {
		if notValid.NotValid == "not valid msg" {
			return nil, ErrMSGNotValid
		} else {
			return nil, errors.New("Not valid: " + notValid.NotValid)
		}
	}

	return rbody, nil
}
