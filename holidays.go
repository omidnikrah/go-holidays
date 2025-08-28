package holidays

import (
	"strings"
	"time"
)

func IsTodayHoliday(countryCode string) bool {
	return IsHoliday(countryCode, time.Now())
}

func GetTodayHolidayCountries() ([]Holiday, error) {
	return getClient().fetchTodayHolidayCountries()
}

func IsHoliday(countryCode string, date time.Time) bool {
	cc := strings.ToUpper(strings.TrimSpace(countryCode))
	client := getClient()

	holidays, err := client.getHolidays(cc, date.Year())
	if err != nil {
		return false
	}

	dateStr := date.Format(defaultDateFormat)

	for _, holiday := range holidays {
		if holiday.Date.Format(defaultDateFormat) == dateStr {
			return true
		}
	}

	return false
}
