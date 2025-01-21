package tendone

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Session struct {
	uri     string
	cookies []*http.Cookie
}

func NewSession(uri string) *Session {
	return &Session{uri: uri}
}

func (s *Session) GetURI() string {
	return s.uri
}

func (s *Session) IsAutheticated() (bool, error) {
	req, err := http.NewRequest("GET", s.uri, nil)
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

	if resp.StatusCode != 200 {
		return false, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// To check if we are authenticated we simply check if the body
	// contains the strings from the login page
	sbody := string(body)
	if strings.Contains(sbody, "login-body") && strings.Contains(sbody, "login-title") {
		return false, nil
	}

	return true, nil
}

func (s *Session) Login(user, passwd string) error {
	bpasswd := base64.StdEncoding.EncodeToString([]byte(passwd))

	rbody := []byte(fmt.Sprintf(`
{
	"sysLogin": {
		"logoff": false,
		"password": "%s",
		"time": "%s",
		"timeZone": 12,
		"username": "%s"
	}
}
`, bpasswd, "2025;1;21;20;8;42", user))

	resp, err := http.Post(s.uri+"/goform/modules", "application/json", bytes.NewReader(rbody))

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	s.cookies = resp.Cookies()

	fmt.Println(string(body))

	return nil
}
