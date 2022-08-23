package dto

import "encoding/json"

var Key = ""

// Credentials is a struct to declare .aws/credentials configuration file
type Credentials struct {
	Region             string `json:"region"`
	AWSAccessKey       string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
	AWSSessionToken    string `json:"aws_session_token"`
}

// Profile is a struct for each account in .aws/credentials configuration file
type Profile struct {
	Key   string
	Creds Credentials `json:"credentials"`
}

// MarshalJSON is an implementation of the function from the official pkg
func (s Profile) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		s.Key: s.Creds,
	}
	return json.Marshal(data)
}