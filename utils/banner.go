package utils

import "fmt"

const (
	colorReset = "\033[0m"
	colorCyan  = "\033[36m"
)

func PrintBanner(version string) {
	fmt.Printf(`%s
		 █████╗ ██╗    ██╗███████╗    ███████╗███████╗ ██████╗ 
		██╔══██╗██║    ██║██╔════╝    ██╔════╝██╔════╝██╔═══██╗		Status: Alpha
		███████║██║ █╗ ██║███████╗    ███████╗███████╗██║   ██║		Version: %s
		██╔══██║██║███╗██║╚════██║    ╚════██║╚════██║██║   ██║		Author: github.com/nanih98
		██║  ██║╚███╔███╔╝███████║    ███████║███████║╚██████╔╝		License: Apache License 2.0
		╚═╝  ╚═╝ ╚══╝╚══╝ ╚══════╝    ╚══════╝╚══════╝ ╚═════╝ 
	%s`, colorCyan, version, colorReset)
}
