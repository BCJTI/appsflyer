package appsflyer

import (
	"time"
)

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

// ParseDateTimeFormat returns
func ParseDateTimeFormat(v string) (time.Time, error) {
	return time.Parse(DateTimeLayout, v)
}
