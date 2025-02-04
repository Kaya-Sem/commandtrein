package cmd

import (
	"fmt"

	"github.com/Kaya-Sem/commandtrein/api"

	table "github.com/Kaya-Sem/commandtrein/tables"
	teaTable "github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(IssuesCommand())
}

func IssuesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "issues",
		Short: "Show transit disturbances",
		// Remove the argument restriction
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			handleIssues()
		},
	}
}

func handleIssues() {
	// Fetch issues using the API
	issues := api.GetIssues()

	columns := []teaTable.Column{
		{Title: "", Width: 70},
	}

	rows := make([]teaTable.Row, len(issues))

	for i, issue := range issues {
		rows[i] = teaTable.Row{
			issue.Title,
		}
	}

	table.RenderTable(columns, rows, issues)
}

func printIssue(i *api.Issue) {
	if i.Type == "planned" {

		fmt.Printf(" \033[2m[gepland]\033[0m ")
	}

	fmt.Printf("%s", i.Title)
	fmt.Println()

}
