package cmd

import (
	"github.com/nanih98/aws-sso/logger"
)

var (
	profileName string
	startURL    string
	region      string
	level       string
)

func init() {
	log := logger.Logger()

	ssoInit := InitSsoCommand(&profileName, &startURL, &region, &log, &level)
	start := StartCommand(&profileName, &log, &level)
	profile := SetProfile(&log, &level)
	version := GetCLIVersion()
	usage := Usage()
	switcher := Switch(&profileName, &log, &level)

	rootCmd.AddCommand(ssoInit)
	rootCmd.AddCommand(start)
	rootCmd.AddCommand(profile)
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(usage)
	rootCmd.AddCommand(switcher)

	//Debug
	switcher.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	switcher.MarkPersistentFlagRequired("profileName")

	ssoInit.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")
	start.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")

	start.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	start.MarkPersistentFlagRequired("profileName")

	ssoInit.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoInit.PersistentFlags().StringVar(&region, "region", "eu-west-1", "AWS region")
	ssoInit.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoInit.MarkPersistentFlagRequired("startURL")
	ssoInit.MarkPersistentFlagRequired("region")
}
