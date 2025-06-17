package logging

import "strings"

type Level uint8

const (
	Invalid Level = iota
	Debug
	Info
	Warn
	Error
)

func LevelFromString(level string) Level {
	level = strings.TrimSpace(strings.ToLower(level))

	switch level {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "":
		return Info
	default:
		return Invalid
	}
}

func (self Level) String() string {
	switch self {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	default:
		return "invalid"
	}
}
