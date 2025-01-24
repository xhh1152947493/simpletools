package logger

import (
	"github.com/rs/zerolog"
	"net/http"
)

type LogEvent struct {
	*zerolog.Event
}

type DebugStr interface {
	DebugStr() string
}

func (e *LogEvent) DebugStr(key string, msg DebugStr) *LogEvent {
	if e.Event != nil && msg != nil {
		e.Str(key, msg.DebugStr())
	}
	return e
}

type CErr interface {
	GetCode() int
	GetErr() string
}

func (e *LogEvent) CErr(err CErr) *LogEvent {
	if e.Event != nil && err != nil {
		e.Int("code", err.GetCode())
		e.Str("err", err.GetErr())
	}
	return e
}

func (e *LogEvent) HttpRequest(r *http.Request) *LogEvent {
	if e.Event != nil {
		e.Str("method", r.Method)
		e.Str("url", r.URL.String())
		e.Str("remoteaddr", r.RemoteAddr)
		e.Str("cookie", r.Header.Get("Cookie"))
		e.Str("content-type", r.Header.Get("Content-Type"))
		e.Str("host", r.Header.Get("Host"))
		e.Str("x-real-ip", r.Header.Get("X-Real-IP"))
		e.Str("x-forwarded-for", r.Header.Get("X-Forwarded-For"))
	}
	return e
}

func (e *LogEvent) UID(uid int64) *LogEvent {
	e.Int64("uid", uid)
	return e
}

func (e *LogEvent) OpenID(openid string) *LogEvent {
	e.Str("openid", openid)
	return e
}

type UserInfo interface {
	Username() string
	Platform() string
}

func (e *LogEvent) User(userinfo UserInfo) *LogEvent {
	if e.Event != nil && userinfo != nil {
		e.Str("username", userinfo.Username())
		e.Str("platform", userinfo.Platform())
	}
	return e
}
