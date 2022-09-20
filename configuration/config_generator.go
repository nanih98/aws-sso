package configuration

import (
	"fmt"
	"os"
	"strings"

	"github.com/nanih98/aws-sso/dto"
)

func ConfigGenerator(account string, awsAccessKey string, awsSecretKey string, awsSessionToken string) (dto.Profile, error) {
	resp := dto.Profile{
		Key: account,
		Creds: dto.Credentials{
			Region:             "eu-west-1",
			AWSAccessKey:       awsAccessKey,
			AWSSecretAccessKey: awsSecretKey,
			AWSSessionToken:    awsSessionToken,
		},
	}
	return resp, nil
}

//func WriteProfileToFile(profile dto.Profile, dirname string) error {
//	f, err := os.OpenFile(dirname+"/.aws/credentials", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
//	if err != nil {
//		return err
//	}
//
//	defer f.Close()
//
//	data, _ := json.Marshal(profile)
//	b := new(bytes.Buffer)
//	convert(strings.NewReader(string(data)), b)
//	fmt.Printf(b.String())
//	f.Write([]byte(strings.ReplaceAll(b.String(), "'", "")))
//	return nil
//}

func ReplaceProfileInFile(filename, profileName string, profile dto.Profile) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, fmt.Sprintf("[%s]", profileName)) {
			lines[i] = fmt.Sprintf("[%s]", profileName)
			lines[i+1] = fmt.Sprintf("aws_access_key_id = %s", profile.Creds.AWSAccessKey)
			lines[i+2] = fmt.Sprintf("aws_secret_access_key = %s", profile.Creds.AWSSecretAccessKey)
			lines[i+3] = fmt.Sprintf("region = %s", profile.Creds.Region)
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
