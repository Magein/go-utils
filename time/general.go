package times

import (
	"fmt"
	"strconv"
	"time"
)

// Unix Get the current timestamp
func Unix() int {
	now := time.Now()
	return int(now.Unix())
}

// Date Get the current date
func Date(parameter ...string) string {

	sp := "-"
	if len(parameter) > 0 {
		sp = parameter[0]
	}

	now := time.Now()

	return fmt.Sprintf("%d%s%s%s%s", now.Year(), sp, padTwoDigit(int(now.Month())), sp, padTwoDigit(now.Day()))
}

// Time Get the current time
func Time(parameter ...string) string {

	sp := ":"
	if len(parameter) > 0 {
		sp = parameter[0]
	}

	now := time.Now()

	return fmt.Sprintf("%s%s%s%s%s", padTwoDigit(now.Hour()), sp, padTwoDigit(now.Minute()), sp, padTwoDigit(now.Second()))
}

// Datetime Get the current date and time
func Datetime() string {
	return Date() + " " + Time()
}

// PadTwoDigit Convert to two digits
func padTwoDigit(number int) string {
	str := ""
	if number < 10 {
		str = "0"
	}
	return str + strconv.Itoa(number)
}
