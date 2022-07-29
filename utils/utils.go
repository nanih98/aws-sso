package utils

import (
	"fmt"
	"io/ioutil"
	"log"
)

// WriteConfigFile first initial config file
func WriteConfigFile(config []byte, appName string) {
	dir := "/tmp/aws-sso/"
	_ = ioutil.WriteFile(dir+appName+".json", config, 0644)
}

// GetConfigurations is a blablabla
func GetConfigurations() {
	files, err := ioutil.ReadDir("/tmp/aws-sso/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
