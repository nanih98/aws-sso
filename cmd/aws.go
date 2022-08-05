package cmd

import (
	sso "github.com/nanih98/aws-sso/aws"
	"github.com/spf13/cobra"
)

var (
	appName  string
	startURL string
	region   string
)

func init() {
	rootCmd.AddCommand(ssoConfig)
	ssoConfig.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoConfig.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoConfig.MarkPersistentFlagRequired("startURL")
	ssoConfig.MarkPersistentFlagRequired("region")
}

var ssoConfig = &cobra.Command{
	Use:   "config",
	Short: "Setup configuration",
	Long:  "Setup SSO configuration like SSO_START_URL, AWS_REGION, ROLE_NAME, ACCOUNT_ID....",
	Run: func(cmd *cobra.Command, args []string) {
		sso.Login(startURL, region)
	},
}

var ssoInit = &cobra.Command{
	Use:   "init",
	Short: "Setup your information regarding to your SSO",
	Long:  "Setup SSO configuration like SSO Start url, AWS region...",
	Run: func(cmd *cobra.Command, args []string) {
		sso.Login(startURL, region)
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
