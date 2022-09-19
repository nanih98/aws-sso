package sso

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/nanih98/aws-sso/dto"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
	"golang.design/x/clipboard"
)

func getProfiles(filepath string) []dto.Profile {
	b, err := os.ReadFile(filepath)

	var items []dto.Profile

	if err != nil {
		fmt.Println(err)
	}

	data := regexp.MustCompile(`\[([^\[\]]*)\]`)

	profiles := data.FindAllString(string(b), -1)

	for _, profile := range profiles {
		profile = strings.Trim(profile, "[")
		profile = strings.Trim(profile, "]")
		items = append(items, dto.Profile{Key: profile})
	}

	return items
}

func clipBoard(log *logger.CustomLogger, profile string) {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("export AWS_PROFILE='%s'", profile)))
	log.Info(fmt.Sprintf("Profile %s copied to the clipboard, paste the command in your terminal to set the AWS_PROFILE env", profile))
}

func Profile(log *logger.CustomLogger) {
	log.Info("Setting your AWS_PROFILE environment variable...")
	log.Info("Reading file .aws/credentials")
	credentialsPath := utils.GetUserHome(log) + "/.aws/credentials"

	profiles := getProfiles(credentialsPath)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🚀 {{ .Key | cyan }}",
		Inactive: "  {{ .Key | cyan }}",
		Selected: "🚀 {{ .Key | red | cyan }}",
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
		Size:      20,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Return the selected profile to the clipboard
	clipBoard(log, profiles[i].Key)
}
