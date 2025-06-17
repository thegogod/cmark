package logging

type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Debugln(v ...any)

	Info(v ...any)
	Infof(format string, v ...any)
	Infoln(v ...any)

	Warn(v ...any)
	Warnf(format string, v ...any)
	Warnln(v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)
	Errorln(v ...any)

	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)

	Log(level Level, v ...any)
	Logf(level Level, format string, v ...any)
	Logln(level Level, v ...any)
}
