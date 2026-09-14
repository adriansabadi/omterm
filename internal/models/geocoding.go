package models

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type GeocodingRequest struct {
	Name        string
	Count       int
	Format      string
	Language    string
	Apikey      string
	CountryCode string
}

type GeocodingResponse struct {
	Results []Location
}

type Location struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Elevation   float64  `json:"elevation"`
	Timezone    string   `json:"timezone"`
	FeatureCode string   `json:"feature_code"`
	CountryCode string   `json:"country_code"`
	Country     string   `json:"country"`
	CountryId   int      `json:"country_id"`
	Population  int      `json:"population"`
	Postcodes   []string `json:"postcodes"`
	Admin1      string   `json:"admin1"`
	Admin2      string   `json:"admin2"`
	Admin3      string   `json:"admin3"`
	Admin4      string   `json:"admin4"`
	Admin1Id    int      `json:"admin1_id"`
	Admin2Id    int      `json:"admin2_id"`
	Admin3Id    int      `json:"admin3_id"`
	Admin4Id    int      `json:"admin4_id"`
}

func (geocoding *GeocodingRequest) BuildUrl(baseUrl string, response *GeocodingResponse) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	u.Scheme = "https"
	v := url.Values{}
	if strings.TrimSpace(geocoding.Name) == "" {
		return "", errors.New("The Geocoding request cannot contain an empty name")
	} else {
		v.Add("name", geocoding.Name)
	}

	if geocoding.Count > 0 {
		v.Add("count", fmt.Sprint(geocoding.Count))
	}

	if strings.TrimSpace(geocoding.Format) != "" {
		v.Add("format", geocoding.Format)
	} else {
		v.Add("format", "json")
	}

	if strings.TrimSpace(geocoding.Language) != "" {
		v.Add("language", geocoding.Language)
	}

	if strings.TrimSpace(geocoding.Apikey) != "" {
		v.Add("apikey", geocoding.Apikey)
	}

	if strings.TrimSpace(geocoding.CountryCode) != "" {
		v.Add("countryCode", geocoding.CountryCode)
	}

	url := baseUrl + "?" + v.Encode()

	fmt.Printf("Geocoding request URL: %s \n", url)

	return url, err
}
