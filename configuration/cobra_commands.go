package configuration

import (
	sso "github.com/nanih98/aws-sso/aws"
	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

func InitSsoCommand(profileName string, startURL string, region string) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Setup your information regarding to your SSO",
		Long:  "Setup SSO configuration like SSO Start url, AWS region...",
		Run: func(cmd *cobra.Command, args []string) {
			GetSSOConfig(log, profileName, startURL, region)
		},
	}
}

func StartCommand(profileName string) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Long:  "Start the application",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := utils.FileExists(log, profileName)
			startURL, region := utils.ReadFile(log, filePath)
			sso.Login(startURL, region)
		},
	}
}
