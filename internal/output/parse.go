package output

import (
	"strings"

	"github.com/adriansabadi/omterm/internal/models"
)

type Direction struct {
	Name         string
	Abbreviation string
}

func ParseForecastParamsHourly(params string) (hourly []models.ForecastParamsHourly) {
	var p []string = strings.Split(params, ",")
	for i := 0; i < len(p); i++ {
		f := models.ForecastParamsHourly(p[i])
		hourly = append(hourly, f)
	}
	return
}

func ParseForecastParamsDaily(params string) (daily []models.ForecastParamsDaily) {
	var p []string = strings.Split(params, ",")
	for i := 0; i < len(p); i++ {
		f := models.ForecastParamsDaily(p[i])
		daily = append(daily, f)
	}
	return
}

func ParseForecastParamsCurrent(params string) (current []models.ForecastParamsCurrent) {
	var p []string = strings.Split(params, ",")
	for i := 0; i < len(p); i++ {
		f := models.ForecastParamsCurrent(p[i])
		current = append(current, f)
	}
	return
}

func ParseForecastParamsMinutely15(params string) (minutely15 []models.ForecastParamsMinutely15) {
	var p []string = strings.Split(params, ",")
	for i := 0; i < len(p); i++ {
		f := models.ForecastParamsMinutely15(p[i])
		minutely15 = append(minutely15, f)
	}
	return
}

func ParseForecastParamsCellSelection(s string) models.ForecastParamsCellSelection {
	return models.ForecastParamsCellSelection(s)
}

func ParseForecastParamsTemperatureUnit(s string) models.ForecastParamsTemperatureUnit {
	return models.ForecastParamsTemperatureUnit(s)
}

func ParseForecastParamsWindSpeedUnit(s string) models.ForecastParamsWindSpeedUnit {
	return models.ForecastParamsWindSpeedUnit(s)
}

func ParseForecastParamsPrecipitationUnit(s string) models.ForecastParamsPrecipitationUnit {
	return models.ForecastParamsPrecipitationUnit(s)
}

func ParseForecastParamsTimeformat(s string) models.ForecastParamsTimeformat {
	return models.ForecastParamsTimeformat(s)
}

func ParseForecastParamsModels(s string) (forecastModels []models.ForecastParamsModels) {
	var p []string = strings.Split(s, ",")
	for i := 0; i < len(p); i++ {
		f := models.ForecastParamsModels(p[i])
		forecastModels = append(forecastModels, f)
	}
	return
}

func ParseWindDirection(degrees float32) Direction {
	directions := []Direction{
		{Name: "North", Abbreviation: "N"},
		{Name: "NorthEast", Abbreviation: "NE"},
		{Name: "East", Abbreviation: "E"},
		{Name: "SouthEast", Abbreviation: "SE"},
		{Name: "South", Abbreviation: "S"},
		{Name: "SouthWest", Abbreviation: "SW"},
		{Name: "West", Abbreviation: "W"},
		{Name: "NorthWest", Abbreviation: "NW"},
	}

	degrees = degrees / 360.0
	if degrees < 0 {
		degrees += 360
	}
	index := int((degrees+22.5)/45) % 8
	return directions[index]
}

func ParseWeatherCode(code int) string {
	codes := make(map[int]string)
	codes[0] = "Clear sky"
	codes[1] = "Mainly clear"
	codes[2] = "Partly cloudy"
	codes[3] = "Overcast"
	codes[45] = "Fog"
	codes[48] = "Depositing rime fog"
	codes[51] = "Drizzle: light intensity"
	codes[53] = "Drizzle: moderate intensity"
	codes[55] = "Drizzle: dense intensity"
	codes[56] = "Freezing Drizzle: light intensity"
	codes[57] = "Freezing Drizzle: dense intensity"
	codes[61] = "Rain: light intensity"
	codes[63] = "Rain: moderate intensity"
	codes[65] = "Rain: heavy intensity"
	codes[66] = "Freezing Rain: light intensity"
	codes[67] = "Freezing Rain: heavy intensity"
	codes[71] = "Snow fall: slight intensity"
	codes[73] = "Snow fall: moderate intensity"
	codes[75] = "Snow fall: heavy intensity"
	codes[77] = "Snow grains"
	codes[80] = "Rain showers: slight"
	codes[81] = "Rain showers: moderate"
	codes[82] = "Rain showers: violent"
	codes[85] = "Snow showers: slight"
	codes[86] = "Snow showers: heavy"
	codes[95] = "Thunderstorm: slight or moderate"
	codes[96] = "Thunderstorm with slight hail"
	codes[99] = "Thunderstorm with heavy hail"
	return codes[code]
}
