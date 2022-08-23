package cmd

import (
	"fmt"
	"os"

	"github.com/nanih98/aws-sso/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-sso",
	Short: "Get your AWS SSO credentials easily.",
	Long:  `Get your AWS SSO credentials easily from your terminal of all your accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintBanner(version)
		cmd.Help()
	},
}

// Execute starts the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("😢 %s\n", err.Error())
		os.Exit(1)
	}
}
