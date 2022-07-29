package cmd

import (
	sso "github.com/nanih98/aws-sso/aws"
	"github.com/spf13/cobra"
)

var (
	appName   string
	startURL  string
	accountID string
	roleName  string
	region    string
)

func init() {
	rootCmd.AddCommand(ssoConfig)
	ssoConfig.PersistentFlags().StringVar(&startURL, "startURL", "", "Setup AWS SSO start URL")
	ssoConfig.PersistentFlags().StringVar(&accountID, "accountID", "", "Account id of AWS where SSO is configured")
	ssoConfig.PersistentFlags().StringVar(&roleName, "roleName", "", "RoleName to assume")
	ssoConfig.PersistentFlags().StringVar(&region, "region", "", "AWS region")
	ssoConfig.MarkPersistentFlagRequired("startURL")
	ssoConfig.MarkPersistentFlagRequired("accountID")
	ssoConfig.MarkPersistentFlagRequired("roleName")
	ssoConfig.MarkPersistentFlagRequired("region")
}

var ssoConfig = &cobra.Command{
	Use:   "config",
	Short: "Setup configuration",
	Long:  "Setup SSO configuration like SSO_START_URL, AWS_REGION, ROLE_NAME, ACCOUNT_ID....",
	Run: func(cmd *cobra.Command, args []string) {
		sso.Login(startURL, accountID, roleName, region)
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
