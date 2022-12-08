package helper

import (
	"regexp"
	"time"
)

// ? convert string to time
func ParseTime(times string) time.Time {
	regexCompile, _ := regexp.Compile(`/`)
	times = regexCompile.ReplaceAllLiteralString(times, "-")

	layout := "1-2-06 15:04"
	resultTime, err := time.Parse(layout, times)
	PanicIFError(err)
	return resultTime
}

// ? convert time to string
func ParseTimeToString(times time.Time) string {
	layout := "1-2-06 15:04"
	timeString := times.Format(layout)
	regexCompile, _ := regexp.Compile(`-`)
	return regexCompile.ReplaceAllLiteralString(timeString, "/")
}
