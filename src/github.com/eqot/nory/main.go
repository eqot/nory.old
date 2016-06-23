package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"./artifact"
)

func main() {
	artifacts := artifact.GetInfo("rxjava")

	fmt.Println(">> " + artifacts[0].Id)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group Id", "Artifact Id", "Version"})

	for _, artifact := range artifacts {
		table.Append([]string{artifact.G, artifact.A, artifact.LatestVersion})
	}
	table.Render()
}
