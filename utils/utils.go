package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nanih98/aws-sso/logger"
)

// WriteConfigFile first initial config file
func WriteConfigFile(config []byte, profileName string, log *logger.CustomLogger) {
	directory, err := UserDirectory(log)
	if err != nil {
		log.Fatal(err)
	}
	fileName := directory + profileName + ".json"
	log.Info("Saving profile configuration for " + profileName)
	_ = ioutil.WriteFile(fileName, config, 0644)
	log.Info("Configuration saved in " + fileName)
}

// UserDirectory is a function to check if the directory to store the config exists
func UserDirectory(log *logger.CustomLogger) (string, error) {
	configPath := GetHomeDir(log) + "/.aws-sso/"
	if err := dirExists(configPath, log); err != nil {
		return "", fmt.Errorf("could not create the directory: %v", err)
	}
	return configPath, nil
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

// GetConfigurations is a blablabla
//func GetConfigurations() {
//	files, err := ioutil.ReadDir("/tmp/aws-sso/")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, file := range files {
//		fmt.Println(file.Name(), file.IsDir())
//	}
//}

func GetHomeDir(logger *logger.CustomLogger) string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}
	return dirname
}
