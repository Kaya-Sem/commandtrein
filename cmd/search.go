package cmd

import (
	"fmt"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd())
}

func searchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Search for train stations",
		Run: func(cmd *cobra.Command, args []string) {

			stationsJSON, err := api.GetSNCBStationsJSON()

			stations, err := api.ParseStations(stationsJSON)
			if err != nil {
				panic(err)
			}

			for _, station := range stations {
				fmt.Printf("%s\n", station.Name)
			}

		},
	}
}
