package configuration

import (
	"encoding/json"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
)

// GetSSOConfig get the user input data
func GetSSOConfig(log *logger.CustomLogger, profileName string, startURL string, region string) {
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

	utils.WriteConfigFile(resp, profileName, log)
}
