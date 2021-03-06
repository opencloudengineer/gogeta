package main

import (
	"github.com/opencloudengineer/gogeta/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	Version := cmd.AppVersion()
	Github := cmd.Github()
	var rootCmd = &cobra.Command{
		Use:   "gogeta",
		Short: "Go Get That App",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(Github)
	rootCmd.AddCommand(Version)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
