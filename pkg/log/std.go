package log

import (
	"fmt"
	"time"
)

type StdLogger struct {
}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) log(level, format string, v ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if format == "" {
		fmt.Printf("[%s] %s: %v\n", timestamp, level, fmt.Sprint(v...))
	} else {
		fmt.Printf("[%s] %s: %s\n", timestamp, level, fmt.Sprintf(format, v...))
	}
}

func (l *StdLogger) Debug(v ...any) {
	l.log("DEBUG", "", v...)
}

func (l *StdLogger) Debugf(format string, v ...any) {
	l.log("DEBUG", format, v...)
}

func (l *StdLogger) Info(v ...any) {
	l.log("INFO", "", v...)
}

func (l *StdLogger) Infof(format string, v ...any) {
	l.log("INFO", format, v...)
}

func (l *StdLogger) Warn(v ...any) {
	l.log("WARN", "", v...)
}

func (l *StdLogger) Warnf(format string, v ...any) {
	l.log("WARN", format, v...)
}

func (l *StdLogger) Error(v ...any) {
	l.log("ERROR", "", v...)
}

func (l *StdLogger) Errorf(format string, v ...any) {
	l.log("ERROR", format, v...)
}
