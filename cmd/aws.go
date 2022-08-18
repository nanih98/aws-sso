package cmd

import (
	"github.com/nanih98/aws-sso/logger"
)

var (
	profileName string
	startURL    string
	region      string
	level       string
	filter 		string
)

func init() {
	log := logger.Logger()

	ssoInit := InitSsoCommand(&profileName, &startURL, &region, &log, &level)
	start := StartCommand(&profileName, &log, &level)
	profile := SetProfile(&log, &level, &filter)

	rootCmd.AddCommand(ssoInit)
	rootCmd.AddCommand(start)
	rootCmd.AddCommand(profile)

	//Debug
	ssoInit.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")
	start.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")

	start.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	start.MarkPersistentFlagRequired("profileName")

	ssoInit.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoInit.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoInit.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoInit.MarkPersistentFlagRequired("startURL")
	ssoInit.MarkPersistentFlagRequired("region")

	profile.PersistentFlags().StringVar(&filter, "filter", "", "Filter the profilename inside .aws/credentials file")
}
