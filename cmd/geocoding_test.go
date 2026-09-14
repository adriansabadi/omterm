package cmd

import (
	"testing"

	"github.com/adriansabadi/omterm/internal/api"
	"github.com/adriansabadi/omterm/internal/models"
)

// TestGeocoding calls client.geocoding, checking a valid return
// for a given city name
func TestGeocoding(t *testing.T) {
	client := api.Client{GeocodingBaseUrl: geocodingURL}
	request := models.GeocodingRequest{
		Name:     "Havana",
		Language: "en",
		Format:   "json",
		Count:    10,
	}
	response := models.GeocodingResponse{}

	response, err := client.Geocoding(request)
	if err != nil {
		t.Error(err)
	}

	if len(response.Results) == 0 {
		t.Errorf("No results found")
	}

	if response.Results[0].Country != "Cuba" {
		t.Errorf("Expected first location country name: %s, recieved: %s", "Cuba", response.Results[0].Country)
	}
}
