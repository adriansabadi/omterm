package cmd

import (
	"errors"

	"github.com/adriansabadi/omterm/internal/api"
	"github.com/adriansabadi/omterm/internal/models"
	"github.com/adriansabadi/omterm/internal/output"
	"github.com/spf13/cobra"
)

var hourlyParams, dailyParams, currentParams, minutely15Params, temperatureUnit, windSpeedUnit, precipitationUnit, timeformat, timezone, startDate, endDate, cellSelection, apiKey, weatherModels string
var forecastDays, pastDays, forecastHours, pastHours int
var elevation, tilt, azimuth float32

const forecastURL string = "https://api.open-meteo.com/v1/forecast?"

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the weather forecast of a location",
	Long:  "Get the weather forecast of a location using Open-Meteo free weather API",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.Client{GeocodingBaseUrl: geocodingURL, ForecastBaseUrl: forecastURL}
		geocodingRequest := models.GeocodingRequest{Name: args[0], Language: "en", Count: 1}
		geocodingResponse := models.GeocodingResponse{}
		geocodingResponse, err := client.Geocoding(geocodingRequest)
		if err != nil {
			return err
		}
		locations := geocodingResponse.Results
		if len(locations) == 0 {
			return errors.New("No location was found")
		}
		var location models.Location = locations[0]

		if forecastDays < 0 {
			return errors.New("--forecastDays cannot be less than 0")
		}

		if pastDays < 0 {
			return errors.New("--pastDays cannot be less than 0")
		}

		if (startDate != "" || endDate != "") && (forecastDays > 0 || pastDays > 0) {
			return errors.New("--forecastDays, --pastDays are mutually exclusive with --startDate,--endDate")
		}

		if forecastHours > 0 && forecastDays == 0 {
			return errors.New("--forecastHours needs --forecastDays greater than 0")
		}

		if pastHours > 0 && pastDays == 0 {
			return errors.New("--pastHours needs --pastDays greater than 0")
		}

		forecastRequest := models.ForecastRequest{
			Latitude:          location.Latitude,
			Longitude:         location.Longitude,
			Hourly:            output.ParseForecastParamsHourly(hourlyParams),
			Daily:             output.ParseForecastParamsDaily(dailyParams),
			Current:           output.ParseForecastParamsCurrent(currentParams),
			Minutely15:        output.ParseForecastParamsMinutely15(minutely15Params),
			Elevation:         elevation,
			TemperatureUnit:   models.ForecastParamsTemperatureUnit(temperatureUnit),
			WindSpeedUnit:     models.ForecastParamsWindSpeedUnit(windSpeedUnit),
			PrecipitationUnit: models.ForecastParamsPrecipitationUnit(precipitationUnit),
			Timeformat:        models.ForecastParamsTimeformat(timeformat),
			Timezone:          timezone,
			PastDays:          pastDays,
			ForecastDays:      forecastDays,
			PastHours:         pastHours,
			ForecastHours:     forecastHours,
			StartDate:         startDate,
			EndDate:           endDate,
			Tilt:              tilt,
			Azimuth:           azimuth,
			CellSelection:     models.ForecastParamsCellSelection(cellSelection),
			Apikey:            apiKey,
			Models:            output.ParseForecastParamsModels(weatherModels),
		}
		forecastResponse, err := client.Forecast(forecastRequest)
		if err != nil {
			return err
		}
		output.PrintHourly(forecastResponse)
		output.PrintDaily(forecastResponse)
		output.PrintCurrent(forecastResponse)
		output.PrintMinutely15(forecastResponse)
		return nil
	},
}

func init() {
	forecastCmd.Flags().StringVar(&hourlyParams, "hourly", "", "Defines hourly weather variables to request")
	forecastCmd.Flags().StringVar(&dailyParams, "daily", "", "Defines daily weather variables to request")
	forecastCmd.Flags().StringVar(&currentParams, "current", "", "Defines current weather variables to request")
	forecastCmd.Flags().StringVar(&minutely15Params, "minutely15", "", "Defines minutely (15min) weather variables to request")
	forecastCmd.Flags().Float32Var(&elevation, "elevation", 0, "Defines the elevation used for statistical downscaling")
	forecastCmd.Flags().StringVar(&temperatureUnit, "temperature_unit", "", "Defines the temperature unit")
	forecastCmd.Flags().StringVar(&windSpeedUnit, "wind_speed_unit", "", "Defines the wind speed unit")
	forecastCmd.Flags().StringVar(&precipitationUnit, "precipitation_unit", "", "Defines the precipitation unit")
	forecastCmd.Flags().StringVar(&timeformat, "timeformat", "", "Defines the time format")
	forecastCmd.Flags().StringVar(&timezone, "timezone", "", "Defines the timezone")
	forecastCmd.Flags().IntVar(&pastDays, "pastDays", 0, "Defines the number of days of archived forecasts")
	forecastCmd.Flags().IntVar(&forecastDays, "forecastDays", 7, "Defines the number of forecast days")
	forecastCmd.Flags().IntVar(&pastHours, "past_hours", 0, "Adjust the archived forecast time range for hourly weather variables")
	forecastCmd.Flags().IntVar(&forecastHours, "forecast_hours", 0, "Adjust the forecast time range for hourly weather variables")
	forecastCmd.Flags().StringVar(&startDate, "start_date", "", "Defines the time interval start date")
	forecastCmd.Flags().StringVar(&endDate, "end_date", "", "Defines the time interval end date")
	forecastCmd.Flags().Float32Var(&tilt, "tilt", 0, "Defines the panel tilt solar radiation variable")
	forecastCmd.Flags().Float32Var(&azimuth, "azimuth", 0, "Defines the azimuth solar radiation variable")
	forecastCmd.Flags().StringVar(&cellSelection, "cell_selection", "", "Set a preference how grid-cells are selected.")
	forecastCmd.Flags().StringVar(&apiKey, "apikey", "", "Defines the Open-Meteo api key to use in the requests")
	forecastCmd.Flags().StringVar(&weatherModels, "models", "", "Select one or more weather models to use")
	rootCmd.AddCommand(forecastCmd)
}
