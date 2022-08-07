package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger function
func Logger() *logrus.Entry {
	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	// log.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors: false,
	// 	FullTimestamp: true,
	// })

	contextLogger := log.WithFields(logrus.Fields{
		"app": "aws-sso",
	})

	return contextLogger
}
