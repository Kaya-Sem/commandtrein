package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	time_query string = ""
	arrival    bool   = false
	Simple     bool
)

var rootCmd = &cobra.Command{
	Use:   "commandtrein [shortcut|station1] [station2]",
	Short: "commandtrein helps you find train connections in Belgium",
	Long: `commandtrein is a CLI tool for checking train schedules and connections in Belgium.
You can use it with station names directly or configure shortcuts for frequent routes.`,

	TraverseChildren: true,
	Args:             cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			cmd.Help()
		case 1:
			// Check if argument is a shortcut
			if shortcut, exists := GetShortcut(args[0]); exists {
				handleConnection(shortcut.Station1, shortcut.Station2, time_query, !arrival)
			} else {
				// Check if it's a valid subcommand
				found := false
				for _, c := range cmd.Root().Commands() {
					if c.Name() == args[0] {
						found = true
						fmt.Printf("Please use: commandtrein %s --help for more information\n\n", args[0])
						break
					}
				}
				if !found {
					style := lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
					fmt.Printf("\n%s\n", style.Render("No such command or shortcut found"))
				}
			}
			os.Exit(1)
		case 2:
			handleConnection(args[0], args[1], time_query, !arrival)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&arrival, "arrival", "a", false, "Use arrival time instead of departure time")
	rootCmd.PersistentFlags().BoolVarP(&Simple, "simple", "s", false, "use simple version")
	rootCmd.PersistentFlags().StringVarP(&time_query, "time", "t", "", "Specify the time in hhmm format")

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
