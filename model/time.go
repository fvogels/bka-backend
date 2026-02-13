package model

import (
	"bass-backend/util"
	"errors"
	"fmt"
	"time"
)

type Time struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

func (time Time) ToHHMMSS() string {
	return fmt.Sprintf("%02d%02d%02d", time.Hour, time.Minute, time.Second)
}

func ParseHHMMSS(str string) (Time, error) {
	ExpectedStringLength := 6
	HourStartIndex := 0
	HourEndIndex := 2
	MinuteStartIndex := 2
	MinuteEndIndex := 4
	SecondStartIndex := 4
	SecondEndIndex := 6

	if len(str) != ExpectedStringLength {
		return Time{}, ErrInvalidString
	}

	hourString := str[HourStartIndex:HourEndIndex]
	minuteString := str[MinuteStartIndex:MinuteEndIndex]
	secondString := str[SecondStartIndex:SecondEndIndex]

	hour, hoursErr := util.ParseInt(hourString)
	minute, minutesErr := util.ParseInt(minuteString)
	second, secondsErr := util.ParseInt(secondString)
	combinedErr := errors.Join(hoursErr, minutesErr, secondsErr)

	return Time{
		Hour:   hour,
		Minute: minute,
		Second: second,
	}, combinedErr
}

func TimeFromTime(time time.Time) Time {
	return Time{
		Hour:   time.Hour(),
		Minute: time.Minute(),
		Second: time.Second(),
	}
}
