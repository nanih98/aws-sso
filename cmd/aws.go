package cmd

import (
	sso "github.com/nanih98/aws-sso/aws"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

var (
	profileName string
	startURL    string
	region      string
)

var log = logger.Logger()

func init() {
	//rootCmd.AddCommand(ssoConfig)
	rootCmd.AddCommand(ssoInit)
	rootCmd.AddCommand(start)
	start.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	start.MarkPersistentFlagRequired("profileName")
	ssoInit.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoInit.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoInit.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoInit.MarkPersistentFlagRequired("startURL")
	ssoInit.MarkPersistentFlagRequired("region")
	ssoInit.MarkPersistentFlagRequired("profileName")
}

var ssoInit = &cobra.Command{
	Use:   "init",
	Short: "Setup your information regarding to your SSO",
	Long:  "Setup SSO configuration like SSO Start url, AWS region...",
	Run: func(cmd *cobra.Command, args []string) {
		configuration.GetSSOConfig(log, profileName, startURL, region)
	},
}

var start = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Long:  "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := utils.FileExists(log, profileName)
		startURL, region := utils.ReadFile(log, filePath)
		sso.Login(startURL, region)
	},
}
