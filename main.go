package main

import (
	"github.com/nanih98/aws-sso/cmd"
	//"github.com/pkg/browser"
	"github.com/nanih98/aws-sso/logger"
)

func main() {
	log := logger.Logger()
	log.Info("Starting aws-sso...")
	cmd.Execute()
}
