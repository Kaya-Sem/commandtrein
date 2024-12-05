package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commandtrein [shortcut|station1] [station2]",
	Short: "commandtrein helps you find train connections in Belgium",
	Long: `commandtrein is a CLI tool for checking train schedules and connections in Belgium.
You can use it with station names directly or configure shortcuts for frequent routes.`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			cmd.Help()
		case 1:
			// Check if argument is a shortcut
			if shortcut, exists := GetShortcut(args[0]); exists {
				handleConnection(shortcut.Station1, shortcut.Station2)
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
					fmt.Println("\nNo such command or shortcut found!")
					cmd.Help()
				}
			}
			os.Exit(1)
		case 2:
			handleConnection(args[0], args[1])
		}
	},
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
