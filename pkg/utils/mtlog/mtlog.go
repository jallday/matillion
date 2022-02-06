package mtlog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Level = logrus.Level

type Fields = logrus.Fields

const (
	TraceLevel Level = logrus.TraceLevel
	DebugLevel Level = logrus.DebugLevel
	InfoLevel  Level = logrus.InfoLevel
	WarnLevel  Level = logrus.WarnLevel
	ErrorLevel Level = logrus.ErrorLevel
	FatalLevel Level = logrus.FatalLevel
)

type Logger struct {
	level  Level
	logger *logrus.Logger
}

func New(level Level) *Logger {
	l := &Logger{
		level:  level,
		logger: logrus.New(),
	}
	l.logger.SetOutput(os.Stdout)
	l.logger.SetFormatter(&logrus.JSONFormatter{})
	l.logger.SetLevel(level)
	return l
}

func (l *Logger) Trace(msg string, fields Fields) {
	l.logger.WithFields(fields).Trace(msg)
}

func (l *Logger) Debug(msg string, fields Fields) {
	l.logger.WithFields(fields).Debug(msg)
}

func (l *Logger) Info(msg string, fields Fields) {
	l.logger.WithFields(fields).Info(msg)
}

func (l *Logger) Warn(msg string, fields Fields) {
	l.logger.WithFields(fields).Warn(msg)
}

func (l *Logger) Error(msg string, fields Fields) {
	l.logger.WithFields(fields).Error(msg)
}

func (l *Logger) Fatal(msg string, fields Fields) {
	l.logger.WithFields(fields).Fatal(msg)
}
