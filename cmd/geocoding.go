package cmd

import (
	"github.com/adriansabadi/omterm/internal/api"
	"github.com/adriansabadi/omterm/internal/models"
	"github.com/adriansabadi/omterm/internal/output"
	"github.com/spf13/cobra"
)

var language, countryCode string
var count int
var format = "json"

const geocodingURL string = "https://geocoding-api.open-meteo.com/v1/search"

var geocodingCmd = &cobra.Command{
	Use:   "geocoding",
	Short: "Search locations globally",
	Long:  "Search locations globally in any language using Open-Meteo Geocoding API",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.Client{GeocodingBaseUrl: geocodingURL}
		request := models.GeocodingRequest{
			Name:     args[0],
			Language: language,
			Format:   format,
			Count:    count,
			Apikey:   apiKey,
		}
		response := models.GeocodingResponse{}
		response, err := client.Geocoding(request)
		if err != nil {
			return err
		}
		output.PrintGeocoding(response)
		return nil
	},
}

func init() {
	geocodingCmd.Flags().IntVar(&count, "count", 10, "Defines the number of search results to return")
	geocodingCmd.Flags().StringVar(&format, "format", "json", "Defines the format of the response")
	geocodingCmd.Flags().StringVar(&language, "language", "en", "Return translated results, if available, otherwise return english or the native location name")
	geocodingCmd.Flags().StringVar(&apiKey, "apikey", "", "Defines the Open-Meteo api key to use in the requests")
	geocodingCmd.Flags().StringVar(&countryCode, "countryCode", "", "Defines the country code of the search")
	rootCmd.AddCommand(geocodingCmd)
}
