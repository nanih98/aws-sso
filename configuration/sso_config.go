package configuration

import (
	"encoding/json"
	"fmt"

	"github.com/nanih98/aws-sso/utils"
)

// PrettyStruct convert data to JSON
// func PrettyStruct(data interface{}) ([]byte, error) {
// 	val, err := json.MarshalIndent(data, "", "    ")
// 	if err != nil {
// 		return ,err
// 	}
// 	return val, err
// }

// Configuration blablabla
type Configuration struct {
	AppName   string `json:"appName"`
	StartURL  string `json:"startUrl"`
	AccountID string `json:"accountID"`
	RoleName  string `json:"roleName"`
	Region    string `json:"region"`
}

// GetSSOConfig get the user input data
func GetSSOConfig(appName string, startURL string, accountID string, roleName string, region string) {
	fmt.Println("Setting up configuration....")

	config := Configuration{
		AppName:   appName,
		StartURL:  startURL,
		AccountID: accountID,
		RoleName:  roleName,
		Region:    region,
	}

	resp, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", resp)
	}

	utils.WriteConfigFile(resp,appName)
	utils.GetConfigurations()
}
