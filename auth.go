package tendone

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type LoginRequest struct {
	SysLogin SysLoginRequest `json:"sysLogin"`
}

type LoginResponse struct {
	SysLogin SysLoginResponse `json:"sysLogin"`
}

type SysLoginRequest struct {
	Logoff   bool   `json:"logoff"`
	Password string `json:"password"`
	Time     string `json:"time"`
	TimeZone int    `json:"timeZone"`
	Username string `json:"username"`
}

type SysLoginResponse struct {
	UserType string `json:"userType"`
	Login    bool   `json:"Login"`
	// When loggin in Logoff is boolean, when logging out Logoff is a string :(
	Logoff interface{} `json:"logoff"`
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

func (s *Session) Login(user, passwd string) (bool, error) {
	bpasswd := base64.StdEncoding.EncodeToString([]byte(passwd))

	sysLogin := LoginRequest{
		SysLogin: SysLoginRequest{
			Logoff:   false,
			Password: bpasswd,
			Time:     time.Now().Format("2006;1;2;15;4;5"),
			TimeZone: 12,
			Username: user,
		},
	}

	rbody, err := json.Marshal(sysLogin)

	resp, err := http.Post(s.uri+MODULES_PATH, "application/json", bytes.NewReader(rbody))

	if err != nil {
		return false, err
	}

	var LoginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&LoginResponse); err != nil {
		return false, err
	}

	s.cookies = resp.Cookies()

	fmt.Println(LoginResponse)

	if LoginResponse.SysLogin.Login && LoginResponse.SysLogin.UserType == user {
		return true, nil
	}

	return false, nil
}

func (s *Session) Logout() (bool, error) {
	sysLogin := LoginRequest{
		SysLogin: SysLoginRequest{
			Logoff: true,
		},
	}

	rbody, err := json.Marshal(sysLogin)

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

	var SysLoginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&SysLoginResponse); err != nil {
		return false, err
	}

	if SysLoginResponse.SysLogin.Logoff.(string) == "ok" {
		return true, nil
	}

	return false, nil

}
