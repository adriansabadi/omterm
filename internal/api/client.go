package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/adriansabadi/omterm/internal/models"
)

type Client struct {
	http             http.Client
	GeocodingBaseUrl string
	ForecastBaseUrl  string
}

func (c *Client) doRequest(url string) ([]byte, error) {
	r, err := c.http.Get(url)

	if r.StatusCode != 200 {
		errorResponse := models.ErrorForecastResponse{}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &errorResponse)
		return nil, errors.New(errorResponse.Reason)
	}
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) Geocoding(request models.GeocodingRequest) (models.GeocodingResponse, error) {
	response := models.GeocodingResponse{}
	if c.GeocodingBaseUrl == "" {
		return response, errors.New("HTTP client must have a geocoding base URL")
	}
	url, err := request.BuildUrl(c.GeocodingBaseUrl)
	if err != nil {
		return response, err
	}

	b, err := c.doRequest(url)

	err = json.Unmarshal(b, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) Forecast(request models.ForecastRequest) (models.ForecastResponse, error) {
	response := models.ForecastResponse{}
	if c.ForecastBaseUrl == "" {
		return response, errors.New("HTTP client must have a forecast base URL")
	}
	url, err := request.BuildUrl(c.ForecastBaseUrl)
	if err != nil {
		return response, err
	}

	b, err := c.doRequest(url)

	err = json.Unmarshal(b, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
