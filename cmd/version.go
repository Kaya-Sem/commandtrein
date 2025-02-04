package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "2.2.7"

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version of commandtrein",
		Run: func(cmd *cobra.Command, args []string) {
			handleVersion()
		},
	}
}

func handleVersion() {
	fmt.Printf("commandtrein %s\n", Version)
}
