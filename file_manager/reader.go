package file_manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nanih98/aws-sso/dto"
)

func (p *FileProcessor) ReadFile(filePath string) (string, string) {
	var data dto.Configuration

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		p.log.Fatal(err)
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		p.log.Fatal(err)
	}

	return data.StartURL, data.Region
}

// FileExists checks if blablabluuu
func (p *FileProcessor) FileExists(profileName string) string {
	filePath := profileName + ".json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		p.log.Fatal(err)
		return ""
	}
	configPath := dirname + "/.aws-sso/" + filePath
	fileExist := checkFileExists(configPath)

	if !fileExist {
		p.log.Fatal(fmt.Errorf("profile don't exists"))
		return ""
	}

	p.log.Info("Profile exists")
	return configPath
}

func (p *FileProcessor) CredentialsExists(profileName string) (string, string) {
	filePath := "credentials." + profileName
	dirname, err := os.UserHomeDir()

	if err != nil {
		p.log.Fatal(err)
		return "", ""
	}

	credentialsPath := dirname + "/.aws/credentials"

	configPath := dirname + "/.aws/" + filePath

	fileExist := checkFileExists(configPath)

	os.Remove(configPath)

	if !fileExist {
		p.log.Fatal(fmt.Errorf("Credentials file don't exist"))
		return "", ""
	}

	p.log.Info("Crentials file exist")

	return configPath, credentialsPath
}

func (p *FileProcessor) SetCredentials(filePath string, credentialsPath string) {
	os.Symlink(filePath, credentialsPath)
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
