package configuration

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestConfigGenerator(t *testing.T) {
	type args struct {
		account         string
		awsAccessKey    string
		awsSecretKey    string
		awsSessionToken string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigGenerator(tt.args.account, tt.args.awsAccessKey, tt.args.awsSecretKey, tt.args.awsSessionToken)
		})
	}
}

func TestProfile_MarshalJSON(t *testing.T) {
	type fields struct {
		Creds Credentials
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Profile{
				Creds: tt.fields.Creds,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteProfileToFile(t *testing.T) {
	type args struct {
		profile  Profile
		filepath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test WriteProfileToFile should create a non existing file",
			args: args{
				profile: Profile{
					Creds: Credentials{
						Region:             "eu-west-1",
						AWSAccessKey:       "accessKeyTest",
						AWSSecretAccessKey: "secretAccessKeyTest",
						AWSSessionToken:    "sessionTokenTest",
					},
				},
				filepath: "/tmp/.aws/credentials",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Mkdir("/tmp/.aws", os.ModePerm)
			WriteProfileToFile(tt.args.profile, "/tmp")
			_, err := os.Stat(tt.args.filepath)
			assert.NoError(t, err)
		})
	}
}
