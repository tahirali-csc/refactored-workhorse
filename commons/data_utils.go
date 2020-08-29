package commons

import "time"

func ParseDBTime(val interface{}) time.Time {
	return ParseDBTimeFromString(val.(string))
}

func ParseDBTimeFromString(val string) time.Time {
	time, _ := time.Parse("2006-01-02T15:04:05.999999999", val)
	return time
}
