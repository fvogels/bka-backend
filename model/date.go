package model

import (
	"bass-backend/util"
	"errors"
	"fmt"
	"time"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (date Date) ToYYYYMMSS() string {
	return fmt.Sprintf("%04d%02d%02d", date.Year, date.Month, date.Day)
}

func ParseYYYYMMSS(str string) (Date, error) {
	ExpectedStringLength := 8
	YearStartIndex := 0
	YearEndIndex := 4
	MonthStartIndex := 4
	MonthEndIndex := 6
	DayStartIndex := 6
	DayEndIndex := 8

	if len(str) != ExpectedStringLength {
		return Date{}, fmt.Errorf("%w: %s does not follow YYYYMMSS format", ErrInvalidString, str)
	}

	yearString := str[YearStartIndex:YearEndIndex]
	monthString := str[MonthStartIndex:MonthEndIndex]
	dayString := str[DayStartIndex:DayEndIndex]

	year, yearErr := util.ParseInt(yearString)
	month, monthErr := util.ParseInt(monthString)
	day, dayErr := util.ParseInt(dayString)
	combinedErr := errors.Join(yearErr, monthErr, dayErr)

	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, combinedErr
}

func DateFromTime(time time.Time) Date {
	return Date{
		Year:  time.Year(),
		Month: int(time.Month()),
		Day:   time.Day(),
	}
}
