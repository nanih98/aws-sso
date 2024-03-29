package configuration

import (
	"os"
	"testing"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/file_manager"
	"github.com/nanih98/aws-sso/logger"
	"github.com/stretchr/testify/assert"
)

//func TestWriteProfileToFile(t *testing.T) {
//	type args struct {
//		profile  dto.Profile
//		filepath string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "Test WriteProfileToFile should create a non existing file",
//			args: args{
//				profile: dto.Profile{
//					Key: "sandbox-lifullconnect",
//					Creds: dto.Credentials{
//						Region:             "eu-west-1",
//						AWSAccessKey:       "accessKeyTest",
//						AWSSecretAccessKey: "secretAccessKeyTest",
//						AWSSessionToken:    "sessionTokenTest",
//					},
//				},
//				filepath: "/tmp/.aws/credentials",
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			os.Mkdir("/tmp/.aws", os.ModePerm)
//			err := WriteProfileToFile(tt.args.profile, "/tmp")
//			assert.NoError(t, err)
//			os.Remove(tt.args.filepath)
//		})
//	}
//}

func TestReplaceProfileInFile(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "credentials")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(`[test-lifull]
aws_access_key_id = 3123123
aws_secret_access_key = 31232131
region = eu-west-1

[test2-lifull]
aws_access_key_id = fed23eqweasdasd
aws_secret_access_key = eqwe234wedwd12
region = eu-west-1`,
	)

	if err != nil {
		panic(err)
	}
	type args struct {
		filename    string
		profileName string
		profile     dto.Profile
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test ReplaceProfileInFile should replace the credentials",
			args: args{
				filename:    f.Name(),
				profileName: "test-lifull",
				profile: dto.Profile{
					Creds: dto.Credentials{
						Region:             "eu-west-1",
						AWSAccessKey:       "accesskeyfake",
						AWSSecretAccessKey: "secretaccesskeyfake",
						AWSSessionToken:    "sessiontokenfake",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = ReplaceProfileInFile(tt.args.filename, tt.args.profileName, tt.args.profile)
			assert.NoError(t, err)
		})
	}
}

func TestWriteProfilesToFile(t *testing.T) {
	type args struct {
		profiles []dto.Profile
		dirname  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test for write multiple profiles",
			args: args{
				profiles: []dto.Profile{
					{
						Key: "sandbox-lifull",
						Creds: dto.Credentials{
							Region:             "eu-west-1",
							AWSAccessKey:       "acceskeyfake",
							AWSSecretAccessKey: "secretkeyfake",
							AWSSessionToken:    "sessiontokenfake",
						},
					},
					{
						Key: "sandbox-lifull2",
						Creds: dto.Credentials{
							Region:             "eu-west-2",
							AWSAccessKey:       "acceskeyfake2",
							AWSSecretAccessKey: "secretkeyfake2",
							AWSSessionToken:    "sessiontokenfake2",
						},
					},
				},
				dirname: "/tmp/.aws/credentials",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileManager := file_manager.NewFileProcessor(&logger.CustomLogger{})
			err := fileManager.WriteProfilesToFile(tt.args.profiles, tt.args.dirname)
			assert.NoError(t, err)
		})
	}
}
