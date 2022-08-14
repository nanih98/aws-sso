package configuration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nanih98/gologger"
	"github.com/pelletier/go-toml/v2"
)

var log = gologger.New(os.Stdout, "", log.Ldate|log.Ltime)

var key string = ""

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

func ConfigGenerator(account string, aws_access_key string, aws_secret_key string, aws_session_token string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// file := os.Remove(dirname+"/.aws/credentials")
	// log.Warn("Removing old credentials file in ", dirname+"/.aws/credentials")
	// if file != nil {
	//     log.Fatal(file)
	// }

	f, err := os.OpenFile(dirname+"/.aws/credentials", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	f.Truncate(0)
	f.Seek(0, 0)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	key = account
	resp := Profile{
		Credentials{
			Region:             "eu-west-1",
			AWSAccessKey:       aws_access_key,
			AWSSecretAccessKey: aws_secret_key,
			AWSSessionToken:    aws_session_token,
		},
	}

	data, _ := json.Marshal(resp)
	b := new(bytes.Buffer)
	convert(strings.NewReader(string(data)), b)
	fmt.Printf(b.String())
	f.Write([]byte(strings.ReplaceAll(b.String(), "'", "")))
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
