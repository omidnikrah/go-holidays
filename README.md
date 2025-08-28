# go-holidays

A Go library for checking public holidays using the [Nager API](https://date.nager.at/). This library provides a simple and efficient way to determine if a specific date is a holiday in any country and to get worldwide holiday information.

## Features

- ✅ Check if today is a holiday in any country
- ✅ Check if a specific date is a holiday in any country

## Installation

```bash
go get github.com/omidnikrah/go-holidays
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/omidnikrah/go-holidays"
)

func main() {
    // Check if today is a holiday in Moldova
    isHoliday := holidays.IsTodayHoliday("MD")
    fmt.Printf("Is today a holiday in Moldova? %t\n", isHoliday)
    
    // Check if a specific date is a holiday in the US
    date := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)
    isChristmasHoliday := holidays.IsHoliday("US", date)
    fmt.Printf("Is Christmas 2024 a holiday in the US? %t\n", isChristmasHoliday)
    
    // Get all countries that have holidays today
    todayHolidays, err := holidays.GetTodayHolidayCountries()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Countries with holidays today: %+v\n", todayHolidays)
}
```

## API Reference

### Functions

#### `IsTodayHoliday(countryCode string) bool`

Checks if today is a holiday in the specified country.

**Parameters:**
- `countryCode` (string): ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "DE")

**Returns:**
- `bool`: `true` if today is a holiday, `false` otherwise

**Example:**
```go
isHoliday := holidays.IsTodayHoliday("US")
```

#### `IsHoliday(countryCode string, date time.Time) bool`

Checks if a specific date is a holiday in the specified country.

**Parameters:**
- `countryCode` (string): ISO 3166-1 alpha-2 country code
- `date` (time.Time): The date to check

**Returns:**
- `bool`: `true` if the date is a holiday, `false` otherwise

**Example:**
```go
date := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)
isHoliday := holidays.IsHoliday("US", date)
```

#### `GetTodayHolidayCountries() ([]Holiday, error)`

Returns a list of all countries that have holidays today.

**Returns:**
- `[]Holiday`: Slice of holiday information
- `error`: Any error that occurred during the request

**Example:**
```go
holidays, err := holidays.GetTodayHolidayCountries()
if err != nil {
    log.Fatal(err)
}
for _, holiday := range holidays {
    fmt.Printf("%s: %s\n", holiday.CountryCode, holiday.Name)
}
```

## Country Codes

This library uses ISO 3166-1 alpha-2 country codes. Some common examples:

- `US` - United States
- `GB` - United Kingdom
- `DE` - Germany
- `FR` - France
- `CA` - Canada
- `AU` - Australia
- `JP` - Japan
- `IN` - India
- `BR` - Brazil
- `MD` - Moldova

For a complete list of supported country codes, refer to the [Nager API documentation](https://date.nager.at/).

## License

This project is open source and available under the [MIT License](LICENSE).
