package cmd

import (
	"github.com/nanih98/aws-sso/configuration"
	"github.com/nanih98/aws-sso/logger"
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
	ssoInit.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoInit.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoInit.PersistentFlags().StringVar(&profileName, "profileName", "", "Profile name")
	ssoInit.MarkPersistentFlagRequired("startURL")
	ssoInit.MarkPersistentFlagRequired("region")
	ssoInit.MarkPersistentFlagRequired("profileName")

}

// var ssoConfig = &cobra.Command{
// 	Use:   "start",
// 	Short: "Setup configuration",
// 	Long:  "Setup SSO configuration like SSO_START_URL, AWS_REGION, ROLE_NAME, ACCOUNT_ID....",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		sso.Login(startURL, region)
// 	},
// }

var ssoInit = &cobra.Command{
	Use:   "init",
	Short: "Setup your information regarding to your SSO",
	Long:  "Setup SSO configuration like SSO Start url, AWS region...",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting the app")
		configuration.GetSSOConfig(profileName, startURL, region)
	},
}

// var requester = &cobra.Command{
// 	Use:   "request",
// 	Short: "Add domain to the request",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		queryGoogle(domain)
// 	},
// }

// func printhelloWorld() {
// 	fmt.Println("Hello World!")
// }

// func queryGoogle(domain string) {
// 	req, err := http.Get(domain)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer req.Body.Close()

// 	fmt.Println(req.StatusCode)
// }
