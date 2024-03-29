package file_manager

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/utils"
	"github.com/pelletier/go-toml/v2"
)

func (p *FileProcessor) WriteProfilesToFile(profiles []dto.Profile, dirname string) error {
	f, err := os.OpenFile(dirname, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	for i := range profiles {
		data, _ := json.Marshal(profiles[i])
		b := new(bytes.Buffer)
		err := convert(strings.NewReader(string(data)), b)
		if err != nil {
			p.log.Warn(err.Error())
		}

		_, err = f.Write([]byte(strings.ReplaceAll(b.String(), "'", "")))
		if err != nil {
			p.log.Warn(err.Error())
		}

		_, err = f.Write([]byte("\n"))
		if err != nil {
			p.log.Warn(err.Error())
		}
	}
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

// WriteConfigFile first initial config file
func (p *FileProcessor) WriteConfigFile(config []byte, profileName string) {
	directory, err := utils.UserDirectory(p.log)
	if err != nil {
		p.log.Fatal(err)
	}
	fileName := directory + profileName + ".json"
	p.log.Info("Saving profile configuration for " + profileName)
	_ = os.WriteFile(fileName, config, 0644)
	p.log.Info("Configuration saved in " + fileName)
}
