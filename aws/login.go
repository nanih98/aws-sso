package sso

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/pkg/browser"
)

// Login function blablabla
func Login(startURL string, region string) {
	log.Println("Starting the program....")
	os.Setenv("AWS_REGION", region)

	// load default aws config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	ssooidcClient, register, deviceAuth := setupSsoOidcClient(startURL, cfg, err)

	// trigger OIDC login. open browser to login. close tab once login is done. press enter to continue
	url := aws.ToString(deviceAuth.VerificationUriComplete)
	fmt.Printf("If browser is not opened automatically, please open link:\n%v\n", url)
	err = browser.OpenURL(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Press ENTER key once login is done")
	_ = bufio.NewScanner(os.Stdin).Scan()

	// generate sso token
	token := generateToken(err, ssooidcClient, register, deviceAuth)

	// create sso client
	ssoClient := sso.NewFromConfig(cfg)
	// list accounts [ONLY provided for better example coverage]
	fmt.Println("Fetching list of all accounts for user")

	accountPaginator := sso.NewListAccountsPaginator(ssoClient, &sso.ListAccountsInput{
		AccessToken: token.AccessToken,
	})

	for accountPaginator.HasMorePages() {
		x, err := accountPaginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
		}
		for _, y := range x.AccountList {
			fmt.Println("-------------------------------------------------------")
			fmt.Printf("\nAccount ID: %v Name: %v Email: %v\n", aws.ToString(y.AccountId), aws.ToString(y.AccountName), aws.ToString(y.EmailAddress))

			// list roles for a given account [ONLY provided for better example coverage]
			fmt.Printf("\n\nFetching roles of account %v for user\n", aws.ToString(y.AccountId))
			rolePaginator := sso.NewListAccountRolesPaginator(ssoClient, &sso.ListAccountRolesInput{
				AccessToken: token.AccessToken,
				AccountId:   y.AccountId,
			})

			for rolePaginator.HasMorePages() {
				z, err := rolePaginator.NextPage(context.TODO())

				if err != nil {
					fmt.Println(err)
				}

				for _, p := range z.RoleList {
					fmt.Printf("Account ID: %v Role Name: %v\n", aws.ToString(p.AccountId), aws.ToString(p.RoleName))
					fmt.Println("Fetching credentials....")
					credentials, err := ssoClient.GetRoleCredentials(context.TODO(), &sso.GetRoleCredentialsInput{
						AccessToken: token.AccessToken,
						AccountId:   p.AccountId,
						RoleName:    p.RoleName,
					})
					if err != nil {
						fmt.Println(err)
					}

					printLoggingStatus(credentials)
					configuration.ConfigGenerator(
						aws.ToString(y.AccountName),
						aws.ToString(credentials.RoleCredentials.AccessKeyId),
						aws.ToString(credentials.RoleCredentials.SecretAccessKey),
						aws.ToString(credentials.RoleCredentials.SessionToken))
				}
			}
		}
	}

	// fmt.Println("-------------------------------------------------------")
	// // exchange token received during oidc flow to fetch actual aws access keys
	// fmt.Printf("\n\nFetching credentails for role %v of account %v for user\n", roleName, accountID)
	// credentials, err := ssoClient.GetRoleCredentials(context.TODO(), &sso.GetRoleCredentialsInput{
	// 	AccessToken: token.AccessToken,
	// 	AccountId:   aws.String(accountID),
	// 	RoleName:    aws.String(roleName),
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // // printing access key to show how they are accessed
	// fmt.Printf("\n\nPriting aws access keysz")
	// fmt.Println("Access key id: ", aws.ToString(credentials.RoleCredentials.AccessKeyId))
	// fmt.Println("Secret access key: ", aws.ToString(credentials.RoleCredentials.SecretAccessKey))
	// fmt.Println("Expiration: ", aws.ToInt64(&credentials.RoleCredentials.Expiration))
	// fmt.Println("Session token: ", aws.ToString(credentials.RoleCredentials.SessionToken))
}

func setupSsoOidcClient(startURL string, cfg aws.Config, err error) (*ssooidc.Client, *ssooidc.RegisterClientOutput, *ssooidc.StartDeviceAuthorizationOutput) {
	// create sso oidc client to trigger login flow
	ssooidcClient := ssooidc.NewFromConfig(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// register your client which is triggering the login flow
	register, err := ssooidcClient.RegisterClient(context.TODO(), &ssooidc.RegisterClientInput{
		ClientName: aws.String("aws-sso"),
		ClientType: aws.String("public"),
		Scopes:     []string{"sso-portal:*"},
	})
	if err != nil {
		fmt.Println(err)
	}

	// authorize your device using the client registration response
	deviceAuth, err := ssooidcClient.StartDeviceAuthorization(context.TODO(), &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     register.ClientId,
		ClientSecret: register.ClientSecret,
		StartUrl:     aws.String(startURL),
	})

	if err != nil {
		fmt.Println(err)
	}
	return ssooidcClient, register, deviceAuth
}

func printLoggingStatus(credentials *sso.GetRoleCredentialsOutput) {
	fmt.Println("Writing file....")
	fmt.Printf("\n\nPrinting credentials")
	fmt.Println("Access key id: ", aws.ToString(credentials.RoleCredentials.AccessKeyId))
	fmt.Println("Secret access key: ", aws.ToString(credentials.RoleCredentials.SecretAccessKey))
	fmt.Println("Expiration: ", aws.ToInt64(&credentials.RoleCredentials.Expiration))
	fmt.Println("Session token: ", aws.ToString(credentials.RoleCredentials.SessionToken))
}

func generateToken(err error, ssooidcClient *ssooidc.Client, register *ssooidc.RegisterClientOutput, deviceAuth *ssooidc.StartDeviceAuthorizationOutput) *ssooidc.CreateTokenOutput {
	token, err := ssooidcClient.CreateToken(context.TODO(), &ssooidc.CreateTokenInput{
		ClientId:     register.ClientId,
		ClientSecret: register.ClientSecret,
		DeviceCode:   deviceAuth.DeviceCode,
		GrantType:    aws.String("urn:ietf:params:oauth:grant-type:device_code"),
	})

	if err != nil {
		fmt.Println(err)
	}
	return token
}
