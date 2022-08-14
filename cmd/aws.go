package cmd

import (
	sso "github.com/nanih98/aws-sso/aws"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/nanih98/gologger"
	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

var (
	profileName string
	startURL    string
	region      string
)

var log = gologger.New(os.Stdout, "", log.Ldate|log.Ltime)

func init() {
	//rootCmd.AddCommand(ssoConfig)
	rootCmd.AddCommand(ssoConfig)
	rootCmd.AddCommand(ssoStart)
	ssoStart.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoStart.MarkPersistentFlagRequired("profileName")
	ssoConfig.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoConfig.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoConfig.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoConfig.MarkPersistentFlagRequired("startURL")
	ssoConfig.MarkPersistentFlagRequired("region")
	ssoConfig.MarkPersistentFlagRequired("profileName")
}

var ssoConfig = &cobra.Command{
	Use:   "config",
	Short: "Setup your information regarding to your SSO",
	Long:  "Setup SSO configuration like SSO Start url, AWS region...",
	Run: func(cmd *cobra.Command, args []string) {
		configuration.GetSSOConfig(log, profileName, startURL, region)
	},
}

var ssoStart = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Long:  "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := utils.FileExists(log, profileName)
		startURL, region := utils.ReadFile(log, filePath)
		sso.Login(log, startURL, region)
	},
}
