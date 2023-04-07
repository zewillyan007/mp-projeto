package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
)

const (
	ColorReset    = 0
	ColorBold     = 1
	ColorRed      = 31
	ColorGreen    = 32
	ColorYellow   = 33
	ColorBlue     = 34
	ColorMagenta  = 35
	ColorCyan     = 36
	ColorGray     = 37
	ColorDarkGray = 90
)

func Colorize(s interface{}, color int, enabled bool) string {
	if !enabled {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, s)
}

func NewConsoleWritePhoenix(noColor bool, PartsOrder, FieldsExclude []string) zerolog.ConsoleWriter {

	output := zerolog.ConsoleWriter{Out: os.Stdout,
		NoColor:    false,
		TimeFormat: "2006-01-02 15:04:05.000000",
	}

	if PartsOrder != nil {
		output.PartsOrder = PartsOrder
	}

	if FieldsExclude != nil {
		output.FieldsExclude = FieldsExclude
	}

	output.FormatLevel = func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case zerolog.LevelTraceValue:
				l = Colorize("TRA", ColorMagenta, noColor)
			case zerolog.LevelDebugValue:
				l = Colorize("DEB", ColorYellow, noColor)
			case zerolog.LevelInfoValue:
				l = Colorize("INF", ColorGreen, noColor)
			case zerolog.LevelWarnValue:
				l = Colorize("WAR", ColorRed, noColor)
			case zerolog.LevelErrorValue:
				l = Colorize("ERR", ColorRed, noColor)
			case zerolog.LevelFatalValue:
				l = Colorize(Colorize("FAT", ColorRed, noColor), ColorBold, noColor)
			case zerolog.LevelPanicValue:
				l = Colorize(Colorize("PAN", ColorRed, noColor), ColorBold, noColor)
			default:
				l = Colorize("???", ColorBold, noColor)
			}
		} else {
			if i == nil {
				l = Colorize("???", ColorBold, noColor)
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return l
	}

	output.FormatMessage = func(i interface{}) string {

		val := fmt.Sprintf("%s", i)
		if len(val) > 1 {
			val = fmt.Sprintf("- %s", i)
		}
		return val
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%v", "")
	}

	output.FormatFieldValue = func(i interface{}) string {

		return fmt.Sprintf("[%s]", i)
	}

	output.FormatCaller = func(i interface{}) string {
		var c string
		if cc, ok := i.(string); ok {
			c = cc
		}
		if len(c) > 0 {
			if cwd, err := os.Getwd(); err == nil {
				if rel, err := filepath.Rel(cwd, c); err == nil {
					c = rel
				}
			}
			c = "[" + Colorize(c, ColorBold, noColor) + "]"
		}
		return c
	}

	return output
}
