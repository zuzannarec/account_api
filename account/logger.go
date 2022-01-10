package accountapi

import (
	"log"
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type StdLogger struct {
	logger *log.Logger
}

func NewStdOutLogger() *StdLogger {
	return &StdLogger{logger: log.New(os.Stdout, "ACCOUNT_API | ", 5)}
}

func (l *StdLogger) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (l *DefaultLogger) Debug(args ...interface{}) {
}

func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
}
