package cmd

import (
	"fmt"

	sso "github.com/nanih98/aws-sso/aws"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

var version string

func InitSsoCommand(profileName *string, startURL *string, region *string, log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Setup your information regarding to your SSO",
		Long:  "Setup SSO configuration like SSO Start url, AWS region...",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			configuration.GetSSOConfig(log, *profileName, *startURL, *region)
		},
	}
}

func StartCommand(profileName *string, log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Long:  "Start the application",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			filePath := utils.FileExists(log, *profileName)
			startURL, region := utils.ReadFile(log, filePath)
			sso.Login(startURL, region, *profileName, sso.NewLogin(log))
		},
	}
}

func SetProfile(log *logger.CustomLogger, level *string, filter *string) *cobra.Command {
	return &cobra.Command{
		Use:   "profile",
		Short: "Set your aws profile",
		Long:  "This script will read your .aws/credentials file and will set the AWS_PROFILE env",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			sso.Profile(log, *filter)
		},
	}
}

func GetCLIVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "aws-sso version you are using",
		Long:  "Get the cli aws-sso version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}
