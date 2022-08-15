package sso

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	"github.com/nanih98/aws-sso/logger"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/nanih98/aws-sso/configuration"
	"github.com/pkg/browser"
)

type AWSLogin struct {
	cfg           aws.Config
	ssooidcClient *ssooidc.Client
	register      *ssooidc.RegisterClientOutput
	deviceAuth    *ssooidc.StartDeviceAuthorizationOutput
	token         *ssooidc.CreateTokenOutput
	ssoClient     *sso.Client
	log           *logger.CustomLogger
}

func NewLogin(log *logger.CustomLogger) *AWSLogin {
	return &AWSLogin{
		cfg:           aws.Config{},
		ssooidcClient: nil,
		register:      nil,
		deviceAuth:    nil,
		token:         nil,
		ssoClient:     nil,
		log:           log,
	}
}

// Login function blablabla
func Login(startURL string, region string, awsSso *AWSLogin) {
	var err error
	awsSso.log.Info("Starting the program....")
	os.Setenv("AWS_REGION", region)
	// load default aws config
	awsSso.cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		awsSso.log.Fatal(err)
	}

	awsSso.SetupSsoOidcClient(startURL)

	// trigger OIDC login. open browser to login. close tab once login is done. press enter to continue
	err = awsSso.TriggerLogin()
	if err != nil {
		awsSso.log.Fatal(err)
	}

	// generate sso token
	err = awsSso.GenerateToken()
	if err != nil {
		awsSso.log.Fatal(err)
	}

	// create sso client
	ssoClient := sso.NewFromConfig(awsSso.cfg)
	// list accounts [ONLY provided for better example coverage]

	awsSso.log.Info("Fetching list of all accounts for user")
	accountPaginator := sso.NewListAccountsPaginator(ssoClient, &sso.ListAccountsInput{
		AccessToken: awsSso.token.AccessToken,
	})

	for accountPaginator.HasMorePages() {
		listAccountsOutput, err := accountPaginator.NextPage(context.TODO())
		if err != nil {
			awsSso.log.Fatal(err)
		}

		for _, accountInfo := range listAccountsOutput.AccountList {
			rolePaginator := awsSso.GetRolePaginator(accountInfo)

			for rolePaginator.HasMorePages() {
				listAccountRolesOutput, err := rolePaginator.NextPage(context.TODO())
				if err != nil {
					awsSso.log.Fatal(err)
				}

				awsSso.FetchRoleCredentials(listAccountRolesOutput, accountInfo)
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

func (a *AWSLogin) GetRolePaginator(accountInfo types.AccountInfo) *sso.ListAccountRolesPaginator {
	a.log.Info("-------------------------------------------------------")
	a.log.Info(fmt.Sprintf("\nAccount ID: %v Name: %v Email: %v\n", aws.ToString(accountInfo.AccountId), aws.ToString(accountInfo.AccountName), aws.ToString(accountInfo.EmailAddress)))
	a.log.Info(fmt.Sprintf("\n\nFetching roles of account %v for user\n", aws.ToString(accountInfo.AccountId)))

	// list roles for a given account [ONLY provided for better example coverage]
	rolePaginator := sso.NewListAccountRolesPaginator(a.ssoClient, &sso.ListAccountRolesInput{
		AccessToken: a.token.AccessToken,
		AccountId:   accountInfo.AccountId,
	})
	return rolePaginator
}

func (a *AWSLogin) FetchRoleCredentials(listAccountRolesOutput *sso.ListAccountRolesOutput, accountInfo types.AccountInfo) {
	for _, roleInfo := range listAccountRolesOutput.RoleList {
		a.log.Info(fmt.Sprintf("Account ID: %v Role Name: %v\n", aws.ToString(roleInfo.AccountId), aws.ToString(roleInfo.RoleName)))
		a.log.Info("Fetching credentials....")
		credentials, err := a.ssoClient.GetRoleCredentials(context.TODO(), &sso.GetRoleCredentialsInput{
			AccessToken: a.token.AccessToken,
			AccountId:   roleInfo.AccountId,
			RoleName:    roleInfo.RoleName,
		})
		if err != nil {
			a.log.Fatal(err)
		}

		printLoggingStatus(credentials, a.log)
		err = configuration.ConfigGenerator(
			aws.ToString(accountInfo.AccountName),
			aws.ToString(credentials.RoleCredentials.AccessKeyId),
			aws.ToString(credentials.RoleCredentials.SecretAccessKey),
			aws.ToString(credentials.RoleCredentials.SessionToken))
		if err != nil {
			a.log.Fatal(err)
		}
	}
}

func (a *AWSLogin) TriggerLogin() error {
	url := aws.ToString(a.deviceAuth.VerificationUriComplete)
	a.log.Info(fmt.Sprintf("If browser is not opened automatically, please open link:\n%v\n", url))
	err := browser.OpenURL(url)
	if err != nil {
		a.log.Fatal(err)
	}
	a.log.Info("Press ENTER key once login is done")
	_ = bufio.NewScanner(os.Stdin).Scan()
	return err
}

func (a *AWSLogin) SetupSsoOidcClient(startURL string) {
	// create sso oidc client to trigger login flow
	a.ssooidcClient = ssooidc.NewFromConfig(a.cfg)

	// register your client which is triggering the login flow
	register, err := a.ssooidcClient.RegisterClient(context.TODO(), &ssooidc.RegisterClientInput{
		ClientName: aws.String("aws-sso"),
		ClientType: aws.String("public"),
		Scopes:     []string{"sso-portal:*"},
	})
	if err != nil {
		a.log.Fatal(err)
	}

	a.register = register

	// authorize your device using the client registration response
	deviceAuth, err := a.ssooidcClient.StartDeviceAuthorization(context.TODO(), &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     register.ClientId,
		ClientSecret: register.ClientSecret,
		StartUrl:     aws.String(startURL),
	})

	if err != nil {
		a.log.Fatal(err)
	}

	a.deviceAuth = deviceAuth
}

func (a *AWSLogin) GenerateToken() error {
	token, err := a.ssooidcClient.CreateToken(context.TODO(), &ssooidc.CreateTokenInput{
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

func printLoggingStatus(credentials *sso.GetRoleCredentialsOutput, customLogger *logger.CustomLogger) {
	customLogger.Info("Writing file....")
	customLogger.Info(fmt.Sprintf("\n\nPrinting credentials"))
	customLogger.Info(fmt.Sprintf("Access key id: %s", aws.ToString(credentials.RoleCredentials.AccessKeyId)))
	customLogger.Info(fmt.Sprintf("Secret access key: %s", aws.ToString(credentials.RoleCredentials.SecretAccessKey)))
	customLogger.Info(fmt.Sprintf("Expiration: %d", aws.ToInt64(&credentials.RoleCredentials.Expiration)))
	customLogger.Info(fmt.Sprintf("Session token: %s", aws.ToString(credentials.RoleCredentials.SessionToken)))
}
