package configuration

import (
	"encoding/json"
	"github.com/nanih98/aws-sso/file_manager"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/logger"
)

// GetSSOConfig get the user input data
func GetSSOConfig(log *logger.CustomLogger, profileName string, startURL string, region string, processor *file_manager.FileProcessor) {
	log.Info("Setting up configuration...")

	config := dto.Configuration{
		ProfileName: profileName,
		StartURL:    startURL,
		Region:      region,
	}

	resp, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	processor.WriteConfigFile(resp, profileName)
}
