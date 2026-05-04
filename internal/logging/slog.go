package logging

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

func InitSlog() {
	s := slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: attrHandler,
		},
	))
	slog.SetDefault(s)
}

func attrHandler(groups []string, a slog.Attr) slog.Attr {
	var newValue any

	switch a.Key {
	case "time":
		value := a.Value.Time()
		newValue = value.Format(time.DateTime)
	case "source":
		value := a.Value.String()
		cutOff := strings.LastIndex(value, "/internal/")
		newValue = strings.Replace(
			strings.TrimSuffix(value[cutOff:], "}"),
			" ",
			":",
			-1,
		)
	default:
		newValue = a.Value
	}

	return slog.Any(a.Key, newValue)
}
