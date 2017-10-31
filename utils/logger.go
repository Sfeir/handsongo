package utils

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	// AppName is the application's name
	AppName = "handsongo"

	// LogStashFormatter is constant used to format logs as logstash format
	LogStashFormatter = "logstash"
	// TextFormatter is constant used to format logs as simple text format
	TextFormatter = "text"
)

// InitLog initializes the logrus logger
func InitLog(logLevel, formatter string) error {

	switch formatter {
	case LogStashFormatter:
		logrus.SetFormatter(&logrustash.LogstashFormatter{
			TimestampFormat: time.RFC3339,
			Type:            AppName,
		})

	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logLevel)

	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		return err
	}

	logrus.SetLevel(level)
	return nil
}
