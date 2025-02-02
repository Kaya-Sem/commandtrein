package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shortcutCmd())
}

func shortcutCmd() *cobra.Command {
	var simple bool

	cmd := &cobra.Command{
		Use:   "shortcut",
		Short: "Manage connection shortcuts",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all shortcuts",
		Run: func(cmd *cobra.Command, args []string) {
			if simple {
				printShortcutsSimple()
			} else {
				printShortcutsFormatted()
			}
		},
	}

	listCmd.Flags().BoolVarP(&simple, "simple", "s", false, "Print shortcuts in simple format")
	cmd.AddCommand(
		&cobra.Command{
			Use:   "add [name] [station1] [station2]",
			Short: "Add a new shortcut",
			Args:  cobra.ExactArgs(3),
			Run: func(cmd *cobra.Command, args []string) {
				if err := AddShortcut(args[0], args[1], args[2]); err != nil {
					fmt.Printf("Error adding shortcut: %v\n", err)
					return
				}
				fmt.Printf("Added shortcut '%s': %s → %s\n", args[0], args[1], args[2])
			},
		},
		listCmd,
	)

	return cmd
}

func printShortcutsFormatted() {
	fmt.Println("Configured shortcuts:")
	for name, shortcut := range config.Shortcuts {
		style := lipgloss.NewStyle().Italic(true)
		fmt.Printf("  %s: %s → %s\n", style.Render(name), shortcut.Station1, shortcut.Station2)
	}
}

func printShortcutsSimple() {
	for name := range config.Shortcuts {
		fmt.Println(name)
	}
}
