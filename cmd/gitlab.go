package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"log"
)

func Gitlab() *cobra.Command {

	example := `gogeta gitlab gitlab.com/gitlab-org/gitlab-runner`

	var command = &cobra.Command{
		Use:   "gitlab",
		Short: `Fetch Gitlab Release(s)`,
		Long: `To Fetch a Gitlab Release
Repository URL is a required argument for the gitlab command`,
		Example:      example,
		SilenceUsage: true,
		Aliases:      []string{"gl", "lab"},
	}

	command.Flags().StringP("match", "m", "", `download release matching a specific pattern.
if no pattern is passed, then all releases are fetched.`)

	command.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {

			cmd.Help()
			return nil
		}

		/*
			match := ""
			if command.Flags().Changed("match") {
				match, _ = command.Flags().GetString("match")
			}
			splitURL := strings.Split(args[0], "/")
			projectName := splitURL[len(splitURL)-1]
			userOrOrgName := splitURL[len(splitURL)-2]

		*/

		git, err := gitlab.NewClient("")
		searchOptions := &gitlab.ListProjectsOptions{Search: gitlab.String("yamllint")}
		projects, _, err := git.Projects.ListProjects(searchOptions)

		fmt.Println(projects)

		if err != nil {
			fmt.Println(err)
		}

		opt := &gitlab.ListReleasesOptions{
			PerPage: 1,
		}
		releases, _, err := git.Releases.ListReleases(250833, opt)
		if err != nil {
			log.Fatal(err)
		}

		// this loop should run only once
		// as we are querying 1 release per page
		for _, release := range releases {

			for _, v := range release.Assets.Links {

				fmt.Println(v.URL)
			}

			for _, v := range release.Assets.Sources {

				fmt.Println(v.URL)
			}

		}

		return nil
	}
	return command
}
