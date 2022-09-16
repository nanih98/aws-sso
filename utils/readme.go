package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/charmbracelet/glamour"
)

var url string = "https://raw.githubusercontent.com/nanih98/aws-sso/main/README.md"

func RenderREADME() {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	out, err := glamour.Render(string(data), "dark")
	fmt.Print(out)
}
