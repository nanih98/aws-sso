package configuration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nanih98/aws-sso/dto"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func ConfigGenerator(account string, awsAccessKey string, awsSecretKey string, awsSessionToken string) error {
	dto.Key = account
	resp := dto.Profile{
		Creds: dto.Credentials{
			Region:             "eu-west-1",
			AWSAccessKey:       awsAccessKey,
			AWSSecretAccessKey: awsSecretKey,
			AWSSessionToken:    awsSessionToken,
		},
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	_, err = os.Stat(dirname + "/.aws/credentials")
	if err != nil {
		err = WriteProfileToFile(resp, dirname)
		if err != nil {
			return err
		}
	} else {
		ReplaceProfileInFile(dirname+"/.aws/credentials", account, resp)
	}

	return nil
}

func WriteProfileToFile(profile dto.Profile, dirname string) error {
	f, err := os.OpenFile(dirname+"/.aws/credentials", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	//f, err := os.OpenFile("/tmp/credentials", os.O_RDWR|os.O_WRONLY|os.O_CREATE, 0600)
	//os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	data, _ := json.Marshal(profile)
	b := new(bytes.Buffer)
	convert(strings.NewReader(string(data)), b)
	fmt.Printf(b.String())
	f.Write([]byte(strings.ReplaceAll(b.String(), "'", "")))
	return nil
}

func convert(r io.Reader, w io.Writer) error {
	var v interface{}

	d := json.NewDecoder(r)
	err := d.Decode(&v)
	if err != nil {
		return err
	}

	e := toml.NewEncoder(w)
	return e.Encode(v)
}

func ReplaceProfileInFile(filename, profileName string, profile dto.Profile) error {
	input, err := ioutil.ReadFile(filename)
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
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
