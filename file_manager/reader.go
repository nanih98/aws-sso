package file_manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/nanih98/aws-sso/dto"
)

func (p *FileProcessor) ReadFile(filePath string) (string, string) {
	var data dto.Configuration

	content, err := os.ReadFile(filePath)
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

// .aws/credentials.lifullconnect exists
func (p *FileProcessor) CredentialsFile(profileName string) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		p.log.Fatal(err)
	}

	// Constant variable .aws/credentials file
	credentials := dirname + "/.aws/credentials"

	// .aws/credentials.PROFILENAME
	credentialsPath := dirname + "/.aws/credentials." + profileName

	fileExist := checkFileExists(credentialsPath)

	if !fileExist {
		p.log.Fatal(fmt.Errorf("Credentials file %s don't exist", credentialsPath))
	}

	credentialsExists := checkFileExists(credentials)

	if credentialsExists {
		os.Remove(credentials)
	}

	os.Symlink(credentialsPath, credentials)

	p.log.Info(fmt.Sprintf("Using credentials %s", credentialsPath))
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
