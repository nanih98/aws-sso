package utils

import (
	"fmt"
	"os"

	"github.com/nanih98/aws-sso/logger"
)

// UserDirectory is a function to check if the directory to store the config exists
func UserDirectory(log *logger.CustomLogger) (string, error) {
	dirname := GetUserHome(log)

	configPath := dirname + "/.aws-sso/"
	if err := dirExists(configPath, log); err != nil {
		return "", fmt.Errorf("could not create the directory: %v", err)
	}
	return configPath, nil
}

// GetUserHome return the home of the user. Example: /Users/myuser or /home/myuser
func GetUserHome(log *logger.CustomLogger) string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return dirname
}

func dirExists(configPath string, log *logger.CustomLogger) error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Warn("Directory " + configPath + " don't exists. Creating a new one...")
		err = os.Mkdir(configPath, 0700)
		if err != nil {
			return err
		}
	}
	return nil
}
