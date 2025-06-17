package logging

import (
	"fmt"
	"os"
	"strings"
)

type ConsoleLogger struct {
	Name     string
	Prefix   string
	Level    Level
	Disabled bool
}

func Console(name string) ConsoleLogger {
	pattern := os.Getenv("CMARK_LOG")

	if pattern == "" {
		pattern = "*"
	}

	return ConsoleLogger{
		Name:     name,
		Prefix:   os.Getenv("CMARK_LOG_PREFIX"),
		Level:    LevelFromString(os.Getenv("CMARK_LOG_LEVEL")),
		Disabled: !Match(name, pattern),
	}
}

func (self ConsoleLogger) Child(name string) ConsoleLogger {
	return Console(strings.Join([]string{self.Name, name}, "."))
}

func (self ConsoleLogger) Debug(v ...any) {
	self.Log(Debug, v...)
}

func (self ConsoleLogger) Debugf(format string, v ...any) {
	self.Logf(Debug, format, v...)
}

func (self ConsoleLogger) Debugln(v ...any) {
	self.Logln(Debug, v...)
}

func (self ConsoleLogger) Log(level Level, v ...any) {
	if self.Disabled || level > self.Level {
		return
	}

	lvl := Text("[" + self.Level.String() + "]")
	name := Text(self.Name)

	switch self.Level {
	case Debug:
		lvl = lvl.BlueForeground().Bold()
		name = name.BlueForeground().Bold()
	case Info:
		lvl = lvl.CyanForeground().Bold()
		name = name.CyanForeground().Bold()
	case Warn:
		lvl = lvl.YellowForeground().Bold()
		name = name.YellowForeground().Bold()
	case Error:
		lvl = lvl.RedForeground().Bold()
		name = name.RedForeground().Bold()
	}

	line := append(
		[]any{lvl.String(), name.String()},
		v...,
	)

	fmt.Fprint(os.Stdout, line...)
}

func (self ConsoleLogger) Logf(level Level, format string, v ...any) {
	if self.Disabled {
		return
	}

	lvl := Text("[" + self.Level.String() + "]")
	name := Text(self.Name)

	switch self.Level {
	case Debug:
		lvl = lvl.BlueForeground().Bold()
		name = name.BlueForeground().Bold()
	case Info:
		lvl = lvl.CyanForeground().Bold()
		name = name.CyanForeground().Bold()
	case Warn:
		lvl = lvl.YellowForeground().Bold()
		name = name.YellowForeground().Bold()
	case Error:
		lvl = lvl.RedForeground().Bold()
		name = name.RedForeground().Bold()
	}

	line := append(
		[]any{lvl.String(), name.String()},
		v...,
	)

	fmt.Fprintf(os.Stdout, format, line...)
}

func (self ConsoleLogger) Logln(level Level, v ...any) {
	if self.Disabled {
		return
	}

	lvl := Text("[" + level.String() + "]")
	name := Text(self.Name)

	switch self.Level {
	case Debug:
		lvl = lvl.BlueForeground().Bold()
		name = name.BlueForeground().Bold()
	case Info:
		lvl = lvl.CyanForeground().Bold()
		name = name.CyanForeground().Bold()
	case Warn:
		lvl = lvl.YellowForeground().Bold()
		name = name.YellowForeground().Bold()
	case Error:
		lvl = lvl.RedForeground().Bold()
		name = name.RedForeground().Bold()
	}

	line := append(
		[]any{lvl.String(), name.String()},
		v...,
	)

	fmt.Fprintln(os.Stdout, line...)
}
