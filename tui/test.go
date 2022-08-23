package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/nanih98/aws-sso/dto"
	"golang.design/x/clipboard"
)

func getProfiles(filepath string) []dto.Profile {
	b, err := ioutil.ReadFile(filepath)

	var items []dto.Profile

	if err != nil {
		fmt.Println(err)
	}

	data := regexp.MustCompile(`\[([^\[\]]*)\]`) // prints content inside brackets, without filtering

	profiles := data.FindAllString(string(b), -1)

	for _, profile := range profiles {
		profile = strings.Trim(profile, "[")
		profile = strings.Trim(profile, "]")
		items = append(items, dto.Profile{Key: profile})
	}

	return items // ["[Production Platform]","[Sandbox Platform]"]
}

func main() {
	profiles := getProfiles("/Users/danielcascalesromero/.aws/credentials")

	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🚀 {{ .Key | cyan }}",
		Inactive: "  {{ .Key | cyan }}",
		Selected: "🚀 {{ .Key | red | cyan }}",
		// Details: `
		// --------- Pepper ----------
		// {{ "Name:" | faint }}	{{ .Name }}
		// {{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
		// {{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		profile := profiles[index]
		name := strings.Replace(strings.ToLower(profile.Key), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select your AWS PROFILE: ",
		Items:     profiles,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("Profile %s copied to the clipboard, paste in your terminal to set the AWS_PROFILE env", profiles[i].Key)
	clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("export AWS_PROFILE=%s", profiles[i].Key)))
	//fmt.Printf("You choose number %d: %s\n", i+1, profiles[i].Key)
}
