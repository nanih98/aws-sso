package cmd

import (
	"fmt"
	"runtime"

	"github.com/nanih98/aws-sso/file_manager"

	sso "github.com/nanih98/aws-sso/aws"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

var (
	version   string
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

func InitSsoCommand(profileName *string, startURL *string, region *string, log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Setup your information regarding to your SSO",
		Long:  "Setup SSO configuration like SSO Start url, AWS region...",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			configuration.GetSSOConfig(log, *profileName, *startURL, *region, file_manager.NewFileProcessor(log))
		},
	}
}

func StartCommand(profileName *string, log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Long:  "Start the application will start the program to fetch all the credentials",
		Run: func(cmd *cobra.Command, args []string) {
			utils.PrintBanner(version)
			log.LogLevel(*level)
			fileProcessor := file_manager.NewFileProcessor(log)
			filePath := fileProcessor.FileExists(*profileName)
			startURL, region := fileProcessor.ReadFile(filePath)
			sso.Login(startURL, region, sso.NewLogin(log, fileProcessor), *profileName)
		},
	}
}

func Switcher(profileName *string, log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "switch",
		Short: "Select what credentials you want to use",
		Long:  "Select what credentials you want to use. Will be used in .aws/credentials file as a symlink",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			fileProcessor := file_manager.NewFileProcessor(log)
			fileProcessor.CredentialsFile(*profileName)
		},
	}
}

func SetProfile(log *logger.CustomLogger, level *string) *cobra.Command {
	return &cobra.Command{
		Use:   "profile",
		Short: "Set your aws profile",
		Long:  "This script will read your .aws/credentials file and will set the AWS_PROFILE env",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			sso.Profile(log)
		},
	}
}

func GetCLIVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "aws-sso version you are using",
		Long:  "Get the cli aws-sso version installed",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("aws-sso: %v with go version %s %s/%s", version, goversion, goos, goarch)
		},
	}
}

func Usage() *cobra.Command {
	return &cobra.Command{
		Use:   "usage",
		Short: "Usage will print the README.md of the project",
		Long:  "Usage will print the readme of the project. You need internet connection because it will download the README from the github repository",
		Run: func(cmd *cobra.Command, args []string) {
			var url string = "https://raw.githubusercontent.com/nanih98/aws-sso/main/docs/usage.md"
			utils.RenderREADME(url)
		},
	}
}
