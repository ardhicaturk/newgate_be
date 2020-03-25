package logger

import (
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Log struct
type (
	Log struct {
		Logger *logrus.Logger
	}
)

// Logger function to create logger
func Logger() (*Log, error) {
	l := logrus.New()
	logf, err := rotatelogs.New(
		"store/log/access_log.%Y%m%d",

		// max : 7 days to keep
		rotatelogs.WithMaxAge(24*7*time.Hour),

		// rotate every day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil, err
	}

	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logf
	l.Level = logrus.DebugLevel

	return &Log{
		Logger: l,
	}, nil
}
