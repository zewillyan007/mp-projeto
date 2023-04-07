package logger

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	port_shared "mp-projeto/shared/port"

	"github.com/rs/zerolog"
)

const (
	CallerPart = "callerp"
	FromPart   = "fromp"
)

type LoggerPhoenix struct {
	from        string
	noColor     bool
	callerLevel int
	logger      zerolog.Logger
	extraParts  map[string]map[string]interface{}
}

func NewLoggerPhoenix(from string, callerLevel int, noColor bool) port_shared.ILogger {

	ExtraParts := make(map[string]map[string]interface{})
	ExtraParts[CallerPart] = make(map[string]interface{})
	ExtraParts[FromPart] = make(map[string]interface{})
	ExtraParts["reqid"] = make(map[string]interface{})
	ExtraParts["method"] = make(map[string]interface{})
	ExtraParts["url"] = make(map[string]interface{})

	ExtraParts[CallerPart]["color"] = nil
	ExtraParts[FromPart]["color"] = nil
	ExtraParts["reqid"]["color"] = ColorCyan
	ExtraParts["method"]["color"] = ColorMagenta
	ExtraParts["url"]["color"] = 38

	PartsOrder := []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		FromPart,
		//CallerPart,
		"method",
		"url",
		"reqid",
		zerolog.MessageFieldName}

	l := LoggerPhoenix{from: from, callerLevel: callerLevel, noColor: noColor}
	l.extraParts = ExtraParts

	var FieldsExlude []string

	for k := range l.extraParts {
		FieldsExlude = append(FieldsExlude, k)
	}

	output := NewConsoleWritePhoenix(true, PartsOrder, FieldsExlude)
	l.logger = zerolog.New(output)

	ctx := l.logger.With().Timestamp().Caller()
	for _, extraPart := range FieldsExlude {
		ctx.Str(extraPart, "")
	}

	l.logger = ctx.Logger()

	for _, ep := range FieldsExlude {
		l.SetExtraPart(ep, "")
	}

	l.SetExtraPart(FromPart, from)

	return &l
}

func (l *LoggerPhoenix) Level(level int) port_shared.ILogger {
	l.callerLevel = level
	return l
}

func (l *LoggerPhoenix) getCallerName() string {
	_, file, line, _ := runtime.Caller(l.callerLevel)
	return filepath.Base(file) + ":" + fmt.Sprintf("%v", line)
}

func (l *LoggerPhoenix) WithReqInf(r *http.Request) port_shared.ILogger {
	return NewLoggerPhoenix(l.from, l.callerLevel, l.noColor).SetExtraPart("reqid", r.Header.Get("reqid")).SetExtraPart("method", r.Method).SetExtraPart("url", r.URL.String())
}

func (l *LoggerPhoenix) SetExtraPart(key, value string) port_shared.ILogger {

	if l.extraParts[key]["color"] != nil {
		value = Colorize(value, l.extraParts[key]["color"].(int), true)
	}
	l.logger = l.logger.With().Str(key, value).Logger()
	return l
}

func (l *LoggerPhoenix) Info(format string, v ...interface{}) {
	l.SetExtraPart(CallerPart, l.getCallerName())
	l.logger.Info().Msg(fmt.Sprintf(format, v...))
}

func (l *LoggerPhoenix) Debug(format string, v ...interface{}) {
	l.SetExtraPart(CallerPart, l.getCallerName())
	l.logger.Debug().Msg(fmt.Sprintf(format, v...))
}

func (l *LoggerPhoenix) Error(format string, v ...interface{}) {
	l.SetExtraPart(CallerPart, l.getCallerName())
	l.logger.Error().Msg(fmt.Sprintf(format, v...))
}

func (l *LoggerPhoenix) Warn(format string, v ...interface{}) {
	l.SetExtraPart(CallerPart, l.getCallerName())
	l.logger.Warn().Msg(fmt.Sprintf(format, v...))
}

func (l *LoggerPhoenix) Fatal(format string, v ...interface{}) {
	l.SetExtraPart(CallerPart, l.getCallerName())
	l.logger.Fatal().Msg(fmt.Sprintf(format, v...))
}
