package tendone

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type loginRequestWrap struct {
	SysLogin loginRequest `json:"sysLogin"`
}

type loginResponseWrap struct {
	SysLogin loginResponse `json:"sysLogin"`
}

// loginRequest is the struct used to login to the AP. It's also used to logout
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Time is formatted as "2006;1;2;15;4;5" and is used to be printed in the logs
	Time     string `json:"time"`
	TimeZone int    `json:"timeZone"`
	Logoff   bool   `json:"logoff"`
}

// loginResponse is the struct that is returned when logging in. It's also used
// to logout
type loginResponse struct {
	UserType string `json:"userType"`
	Login    bool   `json:"Login"`
	// When loggin in Logoff is boolean, when logging out Logoff is a string :(
	Logoff interface{} `json:"logoff"`
}

// IsAutheticated checks if the session is authenticated. This is done by checking
// if the main index.html page redirects to the login page.
//
// TODO: Check if this is the only way or I could use the API used for all the
// library
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
	// contains the strings from the login page.
	sbody := string(body)
	if strings.Contains(sbody, "login-body") && strings.Contains(sbody, "login-title") {
		return false, nil
	}

	return true, nil
}

// Login is used to login a session with the AP. This should be the first function
// to be called after creating a new session.
//
// You can't have multiple sessions for the same AP. You should always call
// logout after you are done with the session.
// For more information, please refer to the known issues in the docs about session management.
func (s *Session) Login(user, passwd string) (bool, error) {
	bpasswd := base64.StdEncoding.EncodeToString([]byte(passwd))

	sysLogin := loginRequestWrap{
		SysLogin: loginRequest{
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

	var LoginResponse loginResponseWrap
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

// Logout is used to logout a session with the AP. This should be the last function
// to be called before the session is destroyed. You should call this function
// even though the session is destroyed because the AP has NOT a strong security
// and session management, so if you leave the session open there is a possiblity
// that someone could use it to access the access point.
func (s *Session) Logout() (bool, error) {
	sysLogin := loginRequestWrap{
		SysLogin: loginRequest{
			Logoff: true,
		},
	}

	rbody, err := json.Marshal(sysLogin)
	if err != nil {
		return false, err
	}

	resp, err := fetch(s, rbody)
	if err != nil {
		return false, err
	}

	var SysLoginResponse loginResponseWrap
	if err := json.Unmarshal(resp, &SysLoginResponse); err != nil {
		return false, errors.Join(err, errors.New("Body: "+string(resp)))
	}

	logoff, ok := SysLoginResponse.SysLogin.Logoff.(string)
	if !ok {
		return false, fmt.Errorf("Logoff is not a string")
	}

	if logoff == "ok" {
		return true, nil
	}

	return false, nil
}
