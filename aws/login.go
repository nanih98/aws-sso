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

type AWSLogin struct {
	cfg          aws.Config
	ssoidcClient *ssooidc.Client
	register     *ssooidc.RegisterClientOutput
	deviceAuth   *ssooidc.StartDeviceAuthorizationOutput
	token        *ssooidc.CreateTokenOutput
	ssoClient    *sso.Client
}

// Login function blablabla
func Login(startURL string, region string) {
	var err error
	awsLogin := AWSLogin{
		cfg:          aws.Config{},
		ssoidcClient: nil,
		register:     nil,
		deviceAuth:   nil,
		token:        nil,
		ssoClient:    nil,
	}

	log.Println("Starting the program....")
	os.Setenv("AWS_REGION", region)

	// load default aws config
	awsLogin.cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	awsLogin.SetupSsoOidcClient(startURL)

	// trigger OIDC login. open browser to login. close tab once login is done. press enter to continue
	err = triggerLogin(awsLogin.deviceAuth, err)
	if err != nil {
		panic(err)
	}

	// generate sso token
	err = awsLogin.GenerateToken()
	if err != nil {
		panic(err)
	}

	// create sso client
	ssoClient := sso.NewFromConfig(awsLogin.cfg)
	// list accounts [ONLY provided for better example coverage]

	fmt.Println("Fetching list of all accounts for user")
	accountPaginator := sso.NewListAccountsPaginator(ssoClient, &sso.ListAccountsInput{
		AccessToken: awsLogin.token.AccessToken,
	})

	for accountPaginator.HasMorePages() {
		listAccountsOutput, err := accountPaginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println(err)
		}

		for _, accountInfo := range listAccountsOutput.AccountList {
			fmt.Println("-------------------------------------------------------")
			fmt.Printf("\nAccount ID: %v Name: %v Email: %v\n", aws.ToString(accountInfo.AccountId), aws.ToString(accountInfo.AccountName), aws.ToString(accountInfo.EmailAddress))
			fmt.Printf("\n\nFetching roles of account %v for user\n", aws.ToString(accountInfo.AccountId))
			// list roles for a given account [ONLY provided for better example coverage]
			rolePaginator := sso.NewListAccountRolesPaginator(ssoClient, &sso.ListAccountRolesInput{
				AccessToken: awsLogin.token.AccessToken,
				AccountId:   accountInfo.AccountId,
			})

			for rolePaginator.HasMorePages() {
				listAccountRolesOutput, err := rolePaginator.NextPage(context.TODO())

				if err != nil {
					fmt.Println(err)
				}

				for _, roleInfo := range listAccountRolesOutput.RoleList {
					fmt.Printf("Account ID: %v Role Name: %v\n", aws.ToString(roleInfo.AccountId), aws.ToString(roleInfo.RoleName))
					fmt.Println("Fetching credentials....")
					credentials, err := ssoClient.GetRoleCredentials(context.TODO(), &sso.GetRoleCredentialsInput{
						AccessToken: awsLogin.token.AccessToken,
						AccountId:   roleInfo.AccountId,
						RoleName:    roleInfo.RoleName,
					})
					if err != nil {
						fmt.Println(err)
					}

					printLoggingStatus(credentials)
					configuration.ConfigGenerator(
						aws.ToString(accountInfo.AccountName),
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

func triggerLogin(deviceAuth *ssooidc.StartDeviceAuthorizationOutput, err error) error {
	url := aws.ToString(deviceAuth.VerificationUriComplete)
	fmt.Printf("If browser is not opened automatically, please open link:\n%v\n", url)
	err = browser.OpenURL(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Press ENTER key once login is done")
	_ = bufio.NewScanner(os.Stdin).Scan()
	return err
}

func (a *AWSLogin) SetupSsoOidcClient(startURL string) {
	// create sso oidc client to trigger login flow
	a.ssoidcClient = ssooidc.NewFromConfig(a.cfg)

	// register your client which is triggering the login flow
	register, err := a.ssoidcClient.RegisterClient(context.TODO(), &ssooidc.RegisterClientInput{
		ClientName: aws.String("aws-sso"),
		ClientType: aws.String("public"),
		Scopes:     []string{"sso-portal:*"},
	})
	if err != nil {
		fmt.Println(err)
	}

	a.register = register

	// authorize your device using the client registration response
	deviceAuth, err := a.ssoidcClient.StartDeviceAuthorization(context.TODO(), &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     register.ClientId,
		ClientSecret: register.ClientSecret,
		StartUrl:     aws.String(startURL),
	})

	if err != nil {
		fmt.Println(err)
	}

	a.deviceAuth = deviceAuth
}

func (a *AWSLogin) GenerateToken() error {
	token, err := a.ssoidcClient.CreateToken(context.TODO(), &ssooidc.CreateTokenInput{
		ClientId:     a.register.ClientId,
		ClientSecret: a.register.ClientSecret,
		DeviceCode:   a.deviceAuth.DeviceCode,
		GrantType:    aws.String("urn:ietf:params:oauth:grant-type:device_code"),
	})
	if err != nil {
		return err
	}

	a.token = token
	return nil
}

func printLoggingStatus(credentials *sso.GetRoleCredentialsOutput) {
	fmt.Println("Writing file....")
	fmt.Printf("\n\nPrinting credentials")
	fmt.Println("Access key id: ", aws.ToString(credentials.RoleCredentials.AccessKeyId))
	fmt.Println("Secret access key: ", aws.ToString(credentials.RoleCredentials.SecretAccessKey))
	fmt.Println("Expiration: ", aws.ToInt64(&credentials.RoleCredentials.Expiration))
	fmt.Println("Session token: ", aws.ToString(credentials.RoleCredentials.SessionToken))
}
