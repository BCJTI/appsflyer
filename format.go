package appsflyer

import (
	"time"
	"strings"
	"strconv"
)

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

func ParseDateTimeFormat(v string) (time.Time, error) {
	return time.Parse(DateTimeLayout, v)
}

func StringToInterface(str string) (value interface{}){

	var err error

	if value, err = strconv.ParseInt(str, 10, 64); err == nil {
		return
	}

	if value, err = strconv.ParseFloat(str,64); err == nil {
		return
	}

	if len(str) > 1 {
		if value, err = strconv.ParseBool(str); err == nil {
			return
		}
	}

	if value, err = time.Parse("2006-01-02", str); err == nil {
		return
	}

	if value, err = time.Parse("2006-01-02 15:04:05", str); err == nil {
		return
	}

	if value, err = time.Parse("2006-01-02T15:04:05", str); err == nil {
		return
	}

	if value, err = time.Parse(time.RFC3339, str); err == nil {
		return
	}

	// string as default
	value = strings.Trim(str," ")

	return

}
