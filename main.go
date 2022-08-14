package main

import (
	"github.com/nanih98/aws-sso/cmd"
	"github.com/nanih98/aws-sso/logger"
)

func main() {
	logger := logger.Logger()
	logger.Info("Starting aws-sso...")
	cmd.Execute()
}
