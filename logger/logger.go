package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	Log *logrus.Entry
}

func (c *CustomLogger) Info(msg string) {
	c.Log.Info(msg)
}

func (c *CustomLogger) Warn(msg string) {
	c.Log.Warn(msg)
}

func (c *CustomLogger) Fatal(msg error) {
	c.Log.Fatal(msg)
}

func (c *CustomLogger) LogLevel(level *string) {
	switch *level {
	case "debug":
		c.Log.Level = logrus.DebugLevel
	default:
		c.Log.Level = logrus.InfoLevel
		fmt.Printf("Level default INFO")
	}
}

// Logger function
func Logger() CustomLogger {
	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		//Level:     logrus.DebugLevel,
	}

	contextLogger := log.WithFields(logrus.Fields{
		"app": "aws-sso",
	})

	return CustomLogger{Log: contextLogger}
}
