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

	rootCmd.AddCommand(ssoInit)
	rootCmd.AddCommand(start)
	rootCmd.AddCommand(profile)
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(usage)

	//Debug
	ssoInit.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")
	start.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")

	start.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	err := start.MarkPersistentFlagRequired("profileName")
	if err != nil {
		log.Warn(err.Error())
	}

	ssoInit.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoInit.PersistentFlags().StringVar(&region, "region", "eu-west-1", "AWS region")
	ssoInit.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	err = ssoInit.MarkPersistentFlagRequired("startURL")
	if err != nil {
		log.Warn(err.Error())
	}

	err = ssoInit.MarkPersistentFlagRequired("region")
	if err != nil {
		log.Warn(err.Error())
	}
}
