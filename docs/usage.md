# Usage

1. Setup your account configuration

This will save a config file inside your home .aws-sso/myaccount.json

```bash
$ aws-sso config --startURL "https://myaccount.awsapps.com/start" --region us-east-1 --profileName="myaccount"
```
* **startURL:** The AWS SSO url (get it from the AWS CONSOLE where you have your configured SSO). *Is required*
* **region:** region of AWS of your SSO resource. The region where you created the SSO inside your AWS account. *Is required*
* **profileName:** the name of the account, for example, the name of the company or the platform. *Is required*

2. Start the application

This will open the browser where you are login with your external IDP provider and will start getting the aws credentials for each account your user have access. Will write all the access key, secret key and token inside your **.aws/credentials** file.

```bash
$ aws-sso start --profileName="myaccount"
```

* **profileName:** is the name of the account you put in the previous step. This command will read the configuration inside your **.aws-sso/** folder and will load the data (**myaccount.json**)
* **level:** logging level of the program. Default: **info** Options: debug, warning, info, error, trace.

1. Set the profile

Once your credentials are saved inside your **.aws/credentials** file, this command will read and promt each profile in an interactive terminal list. Then, the profile you selected, will be copied in your clipboard and then you will execute the command **export AWS_PROFILE=yourprofile**

> Note: in the interactive terminal, you can search for your profiles using SHIFT+/

5. Usage 

This will show this file in a terminal markdown render.

```bash
$ aws-sso usage
```

1. Version

Get your aws-sso version installed in your local.

```bash
$ aws-sso version
```

7. Help

For more help, run:

```bash
$ aws-sso help
```