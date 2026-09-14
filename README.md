# omterm - Open-Meteo Go CLI

omterm (Open Meteo Terminal) is a CLI tool designed with the objective of using [Open-Meteo](https://open-meteo.com) API services. For the moment supports the [geocoding API](https://open-meteo.com/en/docs/geocoding-api) and the [weather forecast API](https://open-meteo.com/en/docs) with some options (I plan to add most of the available options in the future).

## Features

* Check current, hourly, daily and minutely (15 min) weather forecast for 1 to 16 days.
* Can display the results of the requests in the terminal.
* Search for cities using geocoding API

## Installation

### From source

```bash
git clone https://github.com/adriansabadi/omterm.git
cd omterm
go build
```

## Built with

* [Go](https://go.dev/) - Main language
* [Cobra](https://github.com/spf13/cobra) - CLI framework
* [Open-Meteo]([htt/](https://open-meteo.com)) - Free weather API

## Requirements

Requires Go 1.26.5 or later.

## Usage Examples

### PowerShell Terminal

#### Tool help

```powershell
.\omterm --help
```

#### Command help

```powershell
.\omterm forecast --help
```

#### Current weather

```powershell
.\omterm forecast Havana --current temperature_2m
```

### Hourly weather for 3 days

```powershell
.\omterm forecast Havana --hourly temperature_2m,apparent_temperature,precipitation_probability --forecastDays 3
```

### Daily weather for 7 days

```powershell
.\omterm forecast Havana --daily temperature_2m_max,temperature_2m_min --forecastDays 7
```

### Minutely weather for 1 day

```powershell
.\omterm forecast Havana --minutely15 temperature_2m --forecastDays 1
```

## Acknowledgments

This tool uses [Open-Meteo](https://open-meteo.com) API services, if you exceed 10'000 requests per day, please contact Open-Meteo. They reserve the right to block applications and IP addresses that misuse their service.

For commercial use of Open-Meteo APIs, please contact them.

More info about [Open-Meteo terms and privacy](https://github.com/open-meteo/open-meteo#terms--privacy)
