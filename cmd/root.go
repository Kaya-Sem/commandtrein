package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	departure_query bool   = true
	arrival_query   bool   = false
	time_query      string = ""
	isDeparture     bool
	Simple          bool
)

var rootCmd = &cobra.Command{
	Use:   "commandtrein [shortcut|station1] [station2]",
	Short: "commandtrein helps you find train connections in Belgium",
	Long: `commandtrein is a CLI tool for checking train schedules and connections in Belgium.
You can use it with station names directly or configure shortcuts for frequent routes.`,

	TraverseChildren: true,
	Args:             cobra.MaximumNArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Ensure XOR logic for --departure and --arrival
		if departure_query && arrival_query {
			return errors.New("only one of --departure (-d) or --arrival (-a) can be used, not both")
		} else if arrival_query {
			isDeparture = false
		} else {
			isDeparture = true
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			cmd.Help()
		case 1:
			// Check if argument is a shortcut
			if shortcut, exists := GetShortcut(args[0]); exists {

				// FIX: hardcoded true is replacement for the arrival/departure choice. Currently not working with the API
				handleConnection(shortcut.Station1, shortcut.Station2, time_query, true)
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
			// FIX:
			handleConnection(args[0], args[1], time_query, true)
		}
	},
}

func init() {
	//rootCmd.PersistentFlags().BoolVarP(&departure_query, "departure", "d", false, "Use the departure time")
	//rootCmd.PersistentFlags().BoolVarP(&arrival_query, "arrival", "a", false, "Use the arrival time")
	rootCmd.PersistentFlags().BoolVarP(&Simple, "simple", "s", false, "use simple version")
	rootCmd.PersistentFlags().StringVarP(&time_query, "time", "t", "", "Specify the time in hhmm format")
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
