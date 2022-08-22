package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// const (
// 	colorReset = "\033[0m"
// 	colorCyan  = "\033[36m"
// )

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-sso",
	Short: "Get your AWS SSO credentials easily.",
	Long:  `Get your AWS SSO credentials easily from your terminal of all your accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts the cli
func Execute() {
	// fmt.Println(string(colorCyan), `
	// 	 █████╗ ██╗    ██╗███████╗    ███████╗███████╗ ██████╗ 
	// 	██╔══██╗██║    ██║██╔════╝    ██╔════╝██╔════╝██╔═══██╗		Status: Alpha
	// 	███████║██║ █╗ ██║███████╗    ███████╗███████╗██║   ██║		Version: 
	// 	██╔══██║██║███╗██║╚════██║    ╚════██║╚════██║██║   ██║		Author: github.com/nanih98
	// 	██║  ██║╚███╔███╔╝███████║    ███████║███████║╚██████╔╝		License: Apache License 2.0
	// 	╚═╝  ╚═╝ ╚══╝╚══╝ ╚══════╝    ╚══════╝╚══════╝ ╚═════╝ 
	// `, string(colorReset))
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("😢 %s\n", err.Error())
		os.Exit(1)
	}
}
