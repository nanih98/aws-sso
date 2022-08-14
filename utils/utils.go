package utils

import (
	"io/ioutil"
	"os"

	"github.com/nanih98/gologger"
)

var log = gologger.Logger()

// WriteConfigFile first initial config file
func WriteConfigFile(config []byte, profileName string) {
	directory := UserDirectory()
	fileName := directory + profileName + ".json"
	log.Info("Saving profile configuration for " + profileName)
	_ = ioutil.WriteFile(fileName, config, 0644)
	log.Info("Configuration saved in " + fileName)
}

// UserDirectory is a function to check if the directory to store the config exists
func UserDirectory() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := dirname + "/.aws-sso/"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Warn("Directory " + configPath + " don't exists. Creating a new one...")
		err = os.Mkdir(configPath, 0700)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("Directory " + configPath + " exists")
	}
	return configPath
}

// GetConfigurations is a blablabla
// func GetConfigurations() {
// 	files, err := ioutil.ReadDir("/tmp/aws-sso/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		fmt.Println(file.Name(), file.IsDir())
// 	}
// }
