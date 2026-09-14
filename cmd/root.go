package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "omterm",
	Short: "omterm is a cli tool for consulting the weather using open-meteo API",
	Long:  "omterm is a cli tool for consulting the weather using open-meteo API, allowing actions like saving data",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error ocurred while executing omterm: '%s' \n", err)
		os.Exit(1)
	}
}
