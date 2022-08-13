package cmd

import "github.com/nanih98/aws-sso/logger"

var (
	profileName string
	startURL    string
	region      string
)

func init() {
	log := logger.Logger()
	ssoInit := InitSsoCommand(profileName, startURL, region, &log)
	start := StartCommand(profileName, &log)

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
