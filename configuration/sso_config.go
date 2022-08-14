package configuration

import (
	"encoding/json"
	"fmt"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/utils"
	"github.com/nanih98/gologger"
)

// GetSSOConfig get the user input data
func GetSSOConfig(log gologger.CustomLogger, profileName string, startURL string, region string) {
	log.Info("Setting up configuration...")

	config := dto.Configuration{
		ProfileName: profileName,
		StartURL:    startURL,
		Region:      region,
	}

	resp, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	utils.WriteConfigFile(resp, profileName)
}
