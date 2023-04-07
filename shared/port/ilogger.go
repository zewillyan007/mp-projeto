package port

import "net/http"

type ILogger interface {
	Level(level int) ILogger
	SetExtraPart(key, v string) ILogger
	WithReqInf(r *http.Request) ILogger
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}
