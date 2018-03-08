package appsflyer

import (
	"time"
	"regexp"
	"strings"
	"strconv"
	"github.com/asaskevich/govalidator"
	"github.com/metakeule/fmtdate"
)

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

func ParseDateTimeFormat(v string) (time.Time, error) {
	return time.Parse(DateTimeLayout, v)
}

func StringToInterface(str string) (value interface{}){

	// create regex for data validation
	isDate := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	// string as default
	value = strings.Trim(str," ")

	// convert data type
	if govalidator.IsInt(str) {
		value, _ = strconv.ParseInt(str, 10, 64, )
	} else if govalidator.IsFloat(str) {
		value, _ = strconv.ParseFloat(str,64)
	} else if isDate.MatchString(str) {
		value, _ = fmtdate.ParseDate(str)
	}

	return

}
