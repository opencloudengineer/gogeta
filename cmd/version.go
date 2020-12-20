package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version   string = "Dev"
	GitCommit string
)

func AppVersion() *cobra.Command {
	var command = &cobra.Command{
		Use:          "version",
		Short:        "Print Version",
		Example:      `gogeta version`,
		Aliases:      []string{"v"},
		SilenceUsage: false,
	}
	command.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("Version :", Version)
		fmt.Println("Git Commit :", GitCommit)
	}
	return command
}
