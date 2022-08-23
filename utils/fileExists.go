package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/logger"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

// FileExists checks if blablabluuu
func FileExists(log *logger.CustomLogger, profileName string) string {
	// lifullconnect-sso.json
	filePath := profileName + ".json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	configPath := dirname + "/.aws-sso/" + filePath
	fileExist := checkFileExists(configPath)

	if !fileExist {
		log.Fatal(fmt.Errorf("profile don't exists"))
		return ""
	}

	log.Info("Profile exists")
	return configPath
}

func ReadFile(log *logger.CustomLogger, filePath string) (string, string) {
	var data dto.Configuration

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.StartURL, data.Region
}
