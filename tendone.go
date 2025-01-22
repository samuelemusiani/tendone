package tendone

import (
	"net/http"
)

const MODULES_PATH = "/goform/modules"

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
