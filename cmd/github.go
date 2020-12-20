package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"strings"
)

func Github() *cobra.Command {

	example := `gogeta github github.com/GoogleContainerTools/skaffold
gogeta github /stedolan/jq -m linux64
gogeta github aquasecurity/trivy -m 64bit.deb
gogeta github koalaman/shellcheck -m linux.x86_64
gogeta github https://github.com/starship/starship -m linux-gnu`

	var command = &cobra.Command{
		Use:   "github",
		Short: `Fetch Github Release(s)`,
		Long: `To Fetch a Github Release
Repository URL is a required argument for the github command`,
		Example:      example,
		SilenceUsage: true,
		Aliases:      []string{"gh", "hub"},
	}

	command.Flags().StringP("match", "m", "", `download release matching a specific pattern.
if no pattern is passed, then all releases are fetched.`)

	command.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {

			cmd.Help()
			return nil
		}

		match := ""
		if command.Flags().Changed("match") {
			match, _ = command.Flags().GetString("match")
		}
		splitURL := strings.Split(args[0], "/")
		projectName := splitURL[len(splitURL)-1]
		userOrOrgName := splitURL[len(splitURL)-2]

		opt := &github.ListOptions{PerPage: 1}
		client := github.NewClient(nil)
		releasesInfo, _, err := client.Repositories.ListReleases(context.Background(), userOrOrgName, projectName, opt)

		if err != nil {
			fmt.Println(err)
		}

		releasesInfoJSON, err := json.Marshal(releasesInfo)
		if err != nil {
			fmt.Println(err)
		}

		query := fmt.Sprintf("#.assets.#(name%%\"*%s*\")#.browser_download_url", match)
		downloadURL := gjson.Get(string(releasesInfoJSON), query).Array()

		fmt.Println(downloadURL[0].Array())

		return nil
	}
	return command
}
