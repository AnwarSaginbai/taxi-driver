package logger

import (
	"log/slog"
	"os"
)

type Log struct {
	Original *slog.Logger
}

func New() Log {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return Log{Original: logger}
}

func (l *Log) Error(msg string, args ...any) {
	l.Original.Error(msg, args...)
}

func (l *Log) WithError(err error, msg string, args ...any) {
	passArgs := make([]any, len(args)+2)
	passArgs[0] = "error"
	passArgs[1] = err

	for i, arg := range args {
		passArgs[i+2] = arg
	}
	l.Original.Error(msg, passArgs)
}

func (l *Log) Info(msg string, args ...any) {
	l.Original.Info(msg, args...)
}
