package models

import (
	"net/url"
	"testing"
)

const geocodingUrl string = "https://geocoding-api.open-meteo.com/v1/search?"

var mockGeocodingRequest = GeocodingRequest{
	Name:     "Berlin",
	Count:    10,
	Format:   "json",
	Language: "en",
}

func TestGeocodingBuildUrl(t *testing.T) {
	baseUrl := geocodingUrl
	got, err := mockGeocodingRequest.BuildUrl(baseUrl)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	raw := baseUrl + "name=Berlin&count=10&language=en&format=json"
	u, err := url.Parse(raw)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	v := u.Query()

	want := baseUrl + v.Encode()

	if got != want {
		t.Errorf("got: %s,\n wanted: %s\n", got, want)
	}
}
