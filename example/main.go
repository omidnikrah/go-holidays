package main

import (
	"fmt"

	"github.com/omidnikrah/go-holidays"
)

func main() {
	isholiday := holidays.IsTodayHoliday("MD")
	todayHolidays, _ := holidays.GetTodayHolidayCountries()
	fmt.Println("Countries having holiday today:", todayHolidays)
	fmt.Println("Is today a holiday in the MD?", isholiday)
}
