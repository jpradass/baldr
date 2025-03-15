package log

import (
	"go.uber.org/zap"

	"github.com/jpradass/baldr/utils"
)

var logger *Log

type Log struct {
	l *zap.Logger
}

type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string, ...any)
	Fatal(string)
}

func New() (Logger, error) {
	if logger == nil {
		l, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}

		logger = &Log{l: l}
	}

	defer logger.l.Sync()

	return logger, nil
}

func (log *Log) Debug(msg string) {
	log.l.Debug(msg)
}

func (log *Log) Info(msg string) {
	log.l.Info(msg)
}

func (log *Log) Warn(msg string) {
	log.l.Warn(msg)
}

func (log *Log) Error(msg string, args ...any) {
	log.l.Error(msg, utils.Map(args, intoZapField)...)
}

func (log *Log) Fatal(msg string) {
	log.l.Fatal(msg)
}
