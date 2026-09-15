package models

import (
	"net/url"
	"testing"
)

const forecastUrl string = "https://api.open-meteo.com/v1/forecast?"

var mockForecastRequest = ForecastRequest{
	Latitude:  52.52,
	Longitude: 13.41,
	Hourly: []ForecastParamsHourly{
		ForecastParamsHourlyTemperature2m,
		ForecastParamsHourlyApparentTemperature,
		ForecastParamsHourlyPrecipitationProbability,
		ForecastParamsHourlyWindSpeed10m,
		ForecastParamsHourlyUvIndex,
	},
	PastDays:     3,
	ForecastDays: 3,
}

func TestForecastBuildUrl(t *testing.T) {
	baseUrl := forecastUrl
	got, err := mockForecastRequest.BuildUrl(baseUrl)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	raw := baseUrl + "latitude=52.52&longitude=13.41&hourly=temperature_2m,apparent_temperature,precipitation_probability,wind_speed_10m,uv_index&past_days=3&forecast_days=3"
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

func TestParamsHourlyString(t *testing.T) {
	param := ForecastParamsHourlyApparentTemperature
	got := param.String()
	want := "apparent_temperature"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsDailyString(t *testing.T) {
	param := ForecastParamsDailyPrecipitationProbabilityMax
	got := param.String()
	want := "precipitation_probability_max"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsCurrentString(t *testing.T) {
	param := ForecastParamsCurrentTemperature2m
	got := param.String()
	want := "temperature_2m"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsMinutely15String(t *testing.T) {
	param := ForecastParamsMinutely15WindGusts10m
	got := param.String()
	want := "wind_gusts_10m"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsTemperatureUnitString(t *testing.T) {
	unit := Celsius
	got := unit.String()
	want := "celsius"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsWindSpeedUnitString(t *testing.T) {
	unit := Kmh
	got := unit.String()
	want := "kmh"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsPrecipitationUnitString(t *testing.T) {
	param := Inch
	got := param.String()
	want := "inch"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsTimeformatString(t *testing.T) {
	param := Iso8601
	got := param.String()
	want := "iso8601"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsCellSelectionString(t *testing.T) {
	param := Land
	got := param.String()
	want := "land"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}

func TestParamsModelsString(t *testing.T) {
	param := BomAccessGlobal
	got := param.String()
	want := "bom_access_global"
	if got != want {
		t.Errorf("got: %s, wanted: %s\n", got, want)
	}
}
