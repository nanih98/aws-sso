package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-sso",
	Short: "Get your AWS SSO credentials easily.",
	Long:  `Get your AWS SSO credentials easily from your terminal of all your accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute blablabla
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("😢 %s\n", err.Error())
		os.Exit(1)
	}
}
