package configuration

import (
	"github.com/nanih98/aws-sso/dto"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteProfileToFile(t *testing.T) {
	type args struct {
		profile  dto.Profile
		filepath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test WriteProfileToFile should create a non existing file",
			args: args{
				profile: dto.Profile{
					Creds: dto.Credentials{
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
			err := WriteProfileToFile(tt.args.profile, "/tmp")
			assert.NoError(t, err)
			os.Remove(tt.args.filepath)
		})
	}
}
