package log

type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

var global Logger = NewStdLogger()

func SetLogger(logger Logger) {
	global = logger
}

func Debug(v ...any) {
	global.Debug(v...)
}

func Debugf(format string, v ...any) {
	global.Debugf(format, v...)
}

func Info(v ...any) {
	global.Info(v...)
}

func Infof(format string, v ...any) {
	global.Infof(format, v...)
}

func Warn(v ...any) {
	global.Warn(v...)
}

func Warnf(format string, v ...any) {
	global.Warnf(format, v...)
}

func Error(v ...any) {
	global.Error(v...)
}

func Errorf(format string, v ...any) {
	global.Errorf(format, v...)
}
