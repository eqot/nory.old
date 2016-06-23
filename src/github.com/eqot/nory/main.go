package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/urfave/cli.v2"

	"./artifact"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "search artifact",
			Action: func(c *cli.Context) error {
				artifacts := artifact.GetInfo(c.Args().First())

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Group Id", "Artifact Id", "Version"})

				for _, artifact := range artifacts {
					table.Append([]string{artifact.G, artifact.A, artifact.LatestVersion})
				}
				table.Render()

				return nil
			},
		},

		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install artifact",
			Action: func(c *cli.Context) error {
				artifact := artifact.Find(c.Args().First())

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Group Id", "Artifact Id", "Version"})

				table.Append([]string{artifact.G, artifact.A, artifact.LatestVersion})

				table.Render()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
