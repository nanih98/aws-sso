package main

import (
	"fmt"

	"github.com/nanih98/aws-sso/cmd"
)

const (
	colorReset = "\033[0m"
	colorCyan  = "\033[36m"
)

func main() {
	fmt.Println(string(colorCyan), `
		 █████╗ ██╗    ██╗███████╗    ███████╗███████╗ ██████╗ 
		██╔══██╗██║    ██║██╔════╝    ██╔════╝██╔════╝██╔═══██╗		Status: Alpha
		███████║██║ █╗ ██║███████╗    ███████╗███████╗██║   ██║		Version: v0.0.1-alpha
		██╔══██║██║███╗██║╚════██║    ╚════██║╚════██║██║   ██║		Author: github.com/nanih98
		██║  ██║╚███╔███╔╝███████║    ███████║███████║╚██████╔╝		License: Apache License 2.0
		╚═╝  ╚═╝ ╚══╝╚══╝ ╚══════╝    ╚══════╝╚══════╝ ╚═════╝ 
	`, string(colorReset))
	cmd.Execute()
}