package holidays

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const apiBaseURL = "https://date.nager.at/api/v3"

func (c *client) getHolidays(countryCode string, year int) ([]Holiday, error) {
	cacheKey := countryCode + "-" + strconv.Itoa(year)

	c.mu.RLock()

	if holidays, found := c.cache[cacheKey]; found {
		if ttl, ttlFound := c.cacheTTL[cacheKey]; ttlFound && time.Now().Before(ttl) {
			c.mu.RUnlock()
			return holidays, nil
		}
	}

	c.mu.RUnlock()

	holidays, err := c.fetchHolidays(countryCode, year)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	c.cache[cacheKey] = holidays
	c.cacheTTL[cacheKey] = time.Now().Add(24 * time.Hour)
	c.mu.Unlock()

	return holidays, nil
}

func (c *client) fetchHolidays(countryCode string, year int) ([]Holiday, error) {
	url := c.buildEndpoint("/PublicHolidays/", strconv.Itoa(year), countryCode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var apiHolidays []NagerHoliday
	if err := getJSON(ctx, c.httpClient, url, &apiHolidays); err != nil {
		return nil, err
	}

	holidays := mapNagerToHolidays(apiHolidays, "")

	return holidays, nil
}

func (c *client) fetchTodayHolidayCountries() ([]Holiday, error) {
	url := c.buildEndpoint("/NextPublicHolidaysWorldwide")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var apiHolidays []NagerHoliday
	if err := getJSON(ctx, c.httpClient, url, &apiHolidays); err != nil {
		return nil, err
	}

	today := time.Now().Format(defaultDateFormat)
	holidays := mapNagerToHolidays(apiHolidays, today)

	return holidays, nil
}

func (c *client) buildEndpoint(segments ...string) string {
	endpoint, err := url.JoinPath(apiBaseURL, segments...)
	if err != nil {
		return apiBaseURL
	}
	return endpoint
}

func getJSON[T any](ctx context.Context, c *http.Client, url string, out *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Nager: unexpected status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

func mapNagerToHolidays(apiHolidays []NagerHoliday, filterDate string) []Holiday {
	holidays := make([]Holiday, 0, len(apiHolidays))

	for _, holiday := range apiHolidays {
		date, err := time.Parse(defaultDateFormat, holiday.Date)
		if err != nil {
			continue
		}

		if filterDate != "" && date.Format(defaultDateFormat) != filterDate {
			continue
		}

		holidays = append(holidays, Holiday{
			Name:        holiday.Name,
			Date:        date,
			CountryCode: holiday.CountryCode,
		})
	}

	return holidays
}
