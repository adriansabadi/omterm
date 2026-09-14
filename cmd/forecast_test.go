package cmd

import (
	"testing"

	"github.com/adriansabadi/omterm/internal/api"
	"github.com/adriansabadi/omterm/internal/models"
	"github.com/adriansabadi/omterm/internal/output"
)

// TestForecastHourly calls client.Forecast, checking a valid return
// for hourly weather variables
func TestForecastHourly(t *testing.T) {
	client := api.Client{ForecastBaseUrl: forecastURL}
	hourlyParams := "temperature_2m"
	request := models.ForecastRequest{
		Latitude:     52.52,
		Longitude:    13.41,
		Hourly:       output.ParseForecastParamsHourly(hourlyParams),
		ForecastDays: 3,
	}
	response := models.ForecastResponse{}

	response, err := client.Forecast(request)
	if err != nil {
		t.Error(err)
	}

	const numberOfResults int = 72

	if len(response.Hourly.Time) != numberOfResults {
		t.Errorf("Expected quantity of hours: %d, recieved: %d", numberOfResults, len(response.Hourly.Time))
	}

	if len(response.Hourly.Temperature2m) != numberOfResults {
		t.Errorf("Expected quantity of temperatures: %d, recieved: %d", numberOfResults, len(response.Hourly.Temperature2m))
	}
}

// TestForecastDaily calls client.Forecast, checking a valid return
// for daily weather variables
func TestForecastDaily(t *testing.T) {
	client := api.Client{ForecastBaseUrl: forecastURL}
	dailyParams := "temperature_2m_max,temperature_2m_min"
	request := models.ForecastRequest{
		Latitude:     52.52,
		Longitude:    13.41,
		Daily:        output.ParseForecastParamsDaily(dailyParams),
		ForecastDays: 7,
	}
	response := models.ForecastResponse{}

	response, err := client.Forecast(request)
	if err != nil {
		t.Error(err)
	}

	const numberOfResults int = 7

	if len(response.Daily.Time) != numberOfResults {
		t.Errorf("Expected number of days: %d, recieved: %d", numberOfResults, len(response.Daily.Time))
	}

	if len(response.Daily.Temperature2mMin) != numberOfResults {
		t.Errorf("Expected number of min temperatures: %d, recieved: %d", numberOfResults, len(response.Daily.Temperature2mMin))
	}

	if len(response.Daily.Temperature2mMax) != numberOfResults {
		t.Errorf("Expected number of max temperatures: %d, recieved: %d", numberOfResults, len(response.Daily.Temperature2mMax))
	}
}

// TestForecastDaily calls client.Forecast, checking a valid return
// for daily weather variables
func TestForecastMinutely15(t *testing.T) {
	client := api.Client{ForecastBaseUrl: forecastURL}
	minutelyParams := "temperature_2m"
	request := models.ForecastRequest{
		Latitude:     52.52,
		Longitude:    13.41,
		Minutely15:   output.ParseForecastParamsMinutely15(minutelyParams),
		ForecastDays: 1,
	}
	response := models.ForecastResponse{}

	response, err := client.Forecast(request)
	if err != nil {
		t.Error(err)
	}

	const numberOfResults int = 96

	if len(response.Minutely15.Time) != numberOfResults {
		t.Errorf("Expected number of days: %d, recieved: %d", numberOfResults, len(response.Minutely15.Time))
	}

	if len(response.Minutely15.Temperature2m) != numberOfResults {
		t.Errorf("Expected number of temperatures: %d, recieved: %d", numberOfResults, len(response.Minutely15.Temperature2m))
	}
}
