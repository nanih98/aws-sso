package main

import (
	"github.com/nanih98/aws-sso/cmd"
	//"github.com/pkg/browser"
	"github.com/nanih98/gologger"
)

func main() {
	log := gologger.New(os.Stdout, "", log.Ldate|log.Ltime)
	log.Info("Starting aws-sso...")
	cmd.Execute()
}
