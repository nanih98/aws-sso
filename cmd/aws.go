package cmd

import (
	"github.com/nanih98/aws-sso/configuration"
)

var (
	profileName string
	startURL    string
	region      string
)

func init() {
	ssoInit := configuration.InitSsoCommand(profileName, startURL, region)
	start := configuration.StartCommand(profileName)

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
