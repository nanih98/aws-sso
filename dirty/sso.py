#!/usr/bin/env python3

# This script tries to login into the AWS SSO URL and retrieve all the access credentials you're
# allowed to have access to.


import boto3
import webbrowser
import socket
from pathlib import Path
from configparser import ConfigParser
from threading import Thread
from collections import defaultdict

AWS_PROFILE = ''
AWS_DEFAULT_REGION = 'eu-west-1'
AWS_CREDENTIAL_PATH = str(Path.home()) + '/.aws/credentials'
sso_start_url = 'XXXXXX'

client = boto3.client('sso-oidc', region_name=AWS_DEFAULT_REGION)
sso_client = boto3.client('sso', region_name=AWS_DEFAULT_REGION)


def read_config(path):
    config = ConfigParser()
    config.read(path)
    return config


def write_config(path, config):
    with open(path, "w") as destination:
        config.write(destination)


def device_registration(client_name, client_type):
    try:
        response_client_registration = client.register_client(
            clientName=client_name,
            clientType=client_type,
        )
        return response_client_registration['clientId'], response_client_registration['clientSecret']
    except Exception as e:
        return e


def get_auth_device(id, secret, start_url):
    try:
        response_device_authorization = client.start_device_authorization(
            clientId=id,
            clientSecret=secret,
            startUrl=start_url
        )
        return response_device_authorization['verificationUriComplete'], response_device_authorization['deviceCode'], response_device_authorization['userCode']
    except Exception as e:
        return e


def get_token(id, secret, device_code, user_code):
    try:
        response_token_creation = client.create_token(
            clientId=id,
            clientSecret=secret,
            grantType='urn:ietf:params:oauth:grant-type:device_code',  # review
            deviceCode=device_code,
            code=user_code
        )
        return response_token_creation['accessToken']
    except Exception as e:
        return e


def get_list_accounts(token):
    try:
        response_list_accounts = sso_client.list_accounts(
            # nextToken='string',
            maxResults=123,
            accessToken=token
        )
        return response_list_accounts['accountList']
    except Exception as e:
        return e


def get_roles_account(token, accountid):
    try:
        response_account_roles = sso_client.list_account_roles(
            # nextToken='string',
            maxResults=123,
            accessToken=token,
            accountId=accountid
        )
        return response_account_roles['roleList']
    except Exception as e:
        return e


def get_roles_credentials(rolename, accountid, token):
    try:
        response_role_credentials = sso_client.get_role_credentials(
            roleName=rolename,
            accountId=accountid,
            accessToken=token
        )
        return response_role_credentials['roleCredentials']
    except Exception as e:
        return e


def get_credentials(account_data, results):
    account_id = account_data['accountId']
    account_name = account_data['accountName']
    role_name_data = get_roles_account(token, account_id)
    temp_credentials = defaultdict(list)
    for role_data in role_name_data:
        role_name = role_data['roleName']
        account_id = role_data['accountId']
        temp_credentials[role_name].append(
            get_roles_credentials(role_name, account_id, token)
        )
        print(account_name, 'created')
    results[account_name] = temp_credentials


def update_aws_credentials(account_credentials):
    region = AWS_DEFAULT_REGION
    config = read_config(AWS_CREDENTIAL_PATH)
    for account_name, profiles in account_credentials.items():
        for profile, credentials in profiles.items():
            for credential in credentials:
                if profile == "AdministratorAccess":
                    full_name = account_name
                else:
                    full_name = account_name + "-" + profile
                if config.has_section(full_name):
                    config.remove_section(full_name)
                config.add_section(full_name)
                config.set(full_name, "region", region)
                config.set(full_name, "aws_access_key_id",
                           credential["accessKeyId"])
                config.set(full_name, "aws_secret_access_key ",
                           credential["secretAccessKey"])
                config.set(full_name, "aws_session_token",
                           credential["sessionToken"])
    Path(AWS_CREDENTIAL_PATH).parent.mkdir(parents=True, exist_ok=True)
    write_config(AWS_CREDENTIAL_PATH, config)


clientId, clientSecrets = device_registration(socket.gethostname(), 'public')
url, deviceCode, userCode = get_auth_device(
    clientId, clientSecrets, sso_start_url)

try:
    webbrowser.open(url)
except:
    print("Please manual login: %s \n" % (url))

input("After login, press Enter to continue...")

token = get_token(clientId, clientSecrets, deviceCode, userCode)

accounts_list = get_list_accounts(token)

threads = list()
all_credentials = dict()
for account in accounts_list:
    thread = Thread(target=get_credentials, args=(account, all_credentials))
    thread.start()
    threads.append(thread)

for thread in threads:
    thread.join()

update_aws_credentials(all_credentials)
