package configuration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nanih98/aws-sso/utils"
	"io"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var key = ""

// Credentials is a struct to declare .aws/credentials configuration file
type Credentials struct {
	Region             string `json:"region"`
	AWSAccessKey       string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
	AWSSessionToken    string `json:"aws_session_token"`
}

// Profile is a struct for each account in .aws/credentials configuration file
type Profile struct {
	Creds Credentials `json:"credentials"`
}

// MarshalJSON is an implementation of the function from the official pkg
func (s Profile) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		key: s.Creds,
	}
	return json.Marshal(data)
}

func ConfigGenerator(account string, awsAccessKey string, awsSecretKey string, awsSessionToken string) error {
	key = account
	resp := Profile{
		Credentials{
			Region:             "eu-west-1",
			AWSAccessKey:       awsAccessKey,
			AWSSecretAccessKey: awsSecretKey,
			AWSSessionToken:    awsSessionToken,
		},
	}

	err := WriteProfileToFile(resp, utils.GetHomeDir())
	if err != nil {
		return err
	}
	return nil
}

func WriteProfileToFile(profile Profile, dirname string) error {
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
