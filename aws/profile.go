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

func Profile(log *logger.CustomLogger) {
	log.Warn("This command will read all the profiles inside your .aws/credentials file. After reading it, the intention is to set the profile as an AWS_PORFILE variable, but it is not possible from a child process. For more information: ")
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

	log.Info(fmt.Sprintf("Execute $ export AWS_PROFILE='%s' or see this following README...", profiles[i].Key))
}
