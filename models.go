package holidays

import "time"

type Holiday struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	CountryCode string    `json:"country_code"`
}

type NagerHoliday struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

const defaultDateFormat = "2006-01-02"
