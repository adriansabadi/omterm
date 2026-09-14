package output

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/adriansabadi/omterm/internal/models"
)

func PrintGeocoding(geocoding models.GeocodingResponse) {
	w := new(tabwriter.Writer)
	results := geocoding.Results

	if len(geocoding.Results) > 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Geocoding Results\n")
		fmt.Println("------------------------------------------------------------")

		w.Init(os.Stdout, 12, 8, 0, '\t', 0)
		fmt.Fprintln(w,
			"Id", "\t",
			"Nombre", "\t",
			"Latitude", "\t",
			"Longitude", "\t",
			"Country", "\t",
		)
		for i := 0; i < len(geocoding.Results); i++ {
			fmt.Fprintln(w,
				results[i].Id, "\t",
				results[i].Name, "\t",
				results[i].Latitude, "\t",
				results[i].Longitude, "\t",
				results[i].Country, "\t",
			)
		}
		fmt.Fprintln(w)
		w.Flush()
	}
}

func PrintHourly(forecast models.ForecastResponse) {
	w := new(tabwriter.Writer)

	if len(forecast.Hourly.Time) > 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Hourly Weather Variables\n")
		fmt.Println("------------------------------------------------------------")

		fields := reflect.TypeOf(forecast.Hourly)
		values := reflect.ValueOf(forecast.Hourly)
		units := reflect.ValueOf(forecast.HourlyUnits)

		for i := 0; i < fields.NumField(); i++ {
			field := fields.Field(i)
			value := values.Field(i)

			if field.Name != "Time" && !value.IsZero() {
				// Format in tab-separated columns with a tab stop of 8.
				w.Init(os.Stdout, 0, 8, 0, '\t', 0)
				if value.Kind() == reflect.Slice {
					unit := fmt.Sprintf("(%s)", units.FieldByName(field.Name))
					fmt.Fprintln(w, "No.", "\t", "Time", "\t", fields.Field(i).Name, unit, "\t")
					fmt.Fprintln(w, "---", "\t", "----------------", "\t", "-----------", "\t")
					for j := 0; j < value.Len(); j++ {
						fmt.Fprintln(w, j+1, "\t", forecast.Hourly.Time[j], "\t", value.Index(j).Interface(), "\t")
					}
					fmt.Fprintln(w)
				}
				w.Flush()
			}
		}
	}
}

func PrintDaily(forecast models.ForecastResponse) {
	w := new(tabwriter.Writer)

	if len(forecast.Daily.Time) > 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Daily Weather Variables\n")
		fmt.Println("------------------------------------------------------------")

		fields := reflect.TypeOf(forecast.Daily)
		values := reflect.ValueOf(forecast.Daily)
		units := reflect.ValueOf(forecast.DailyUnits)

		for i := 0; i < fields.NumField(); i++ {
			field := fields.Field(i)
			value := values.Field(i)

			if field.Name != "Time" && !value.IsZero() {
				// Format in tab-separated columns with a tab stop of 8.
				w.Init(os.Stdout, 0, 8, 0, '\t', 0)
				if value.Kind() == reflect.Slice {
					unit := fmt.Sprintf("(%s)", units.FieldByName(field.Name))
					fmt.Fprintln(w, "No.", "\t", "Time", "\t", fields.Field(i).Name, unit, "\t")
					fmt.Fprintln(w, "---", "\t", "----------------", "\t", "-----------", "\t")
					for j := 0; j < value.Len(); j++ {
						fmt.Fprintln(w, j+1, "\t", forecast.Daily.Time[j], "\t", value.Index(j).Interface(), "\t")
					}
					fmt.Fprintln(w)
				}
				w.Flush()
			}
		}
	}
}

func PrintMinutely15(forecast models.ForecastResponse) {
	w := new(tabwriter.Writer)

	if len(forecast.Minutely15.Time) > 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Minutely (15) Weather Variables\n")
		fmt.Println("------------------------------------------------------------")

		fields := reflect.TypeOf(forecast.Minutely15)
		values := reflect.ValueOf(forecast.Minutely15)
		units := reflect.ValueOf(forecast.Minutely15Units)

		for i := 0; i < fields.NumField(); i++ {
			field := fields.Field(i)
			value := values.Field(i)

			if field.Name != "Time" && !value.IsZero() {
				// Format in tab-separated columns with a tab stop of 8.
				w.Init(os.Stdout, 0, 8, 0, '\t', 0)
				if value.Kind() == reflect.Slice {
					unit := fmt.Sprintf("(%s)", units.FieldByName(field.Name))
					fmt.Fprintln(w, "No.", "\t", "Time", "\t", fields.Field(i).Name, unit, "\t")
					fmt.Fprintln(w, "---", "\t", "----------------", "\t", "-----------", "\t")
					for j := 0; j < value.Len(); j++ {
						fmt.Fprintln(w, j+1, "\t", forecast.Minutely15.Time[j], "\t", value.Index(j).Interface(), "\t")
					}
					fmt.Fprintln(w)
				}
				w.Flush()
			}
		}
	}
}

func PrintCurrent(forecast models.ForecastResponse) {
	w := new(tabwriter.Writer)

	if (forecast.Current != models.Current{}) {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Current Weather Variables\n")
		fmt.Println("------------------------------------------------------------")

		fields := reflect.TypeOf(forecast.Current)
		values := reflect.ValueOf(forecast.Current)
		units := reflect.ValueOf(forecast.CurrentUnits)

		for i := 0; i < fields.NumField(); i++ {
			field := fields.Field(i)
			value := values.Field(i)

			if field.Name != "Time" && !value.IsZero() {
				// Format in tab-separated columns with a tab stop of 18.
				w.Init(os.Stdout, 0, 24, 0, '\t', 0)
				if units.FieldByName(field.Name).Kind() != reflect.Invalid {
					fmt.Fprintln(w, field.Name, "\t", value.Interface(), "\t", units.FieldByName(field.Name), "\t")
				} else {
					fmt.Fprintln(w, field.Name, "\t", value.Interface(), "\t", "", "\t")
				}
				fmt.Fprintln(w)
				w.Flush()
			}
		}
	}
}
