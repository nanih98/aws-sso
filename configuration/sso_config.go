package configuration

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nanih98/aws-sso/utils"
)

// Configuration blablabla
type Configuration struct {
	ProfileName string `json:"profileName"`
	StartURL    string `json:"startURL"`
	Region      string `json:"region"`
}

// GetSSOConfig get the user input data
func GetSSOConfig(profileName string, startURL string, region string) {
	log.Println("Setting up configuration...")

	config := Configuration{
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
