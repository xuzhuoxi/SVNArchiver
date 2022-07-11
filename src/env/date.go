// Create on 2022/7/8
// @author xuzhuoxi
package env

import (
	"errors"
	"fmt"
	"time"
)

const (
	localInputLayout0 = "20060102T150405"
	localInputLayout1 = "20060102T1504"
	localInputLayout2 = "20060102T15"
	localInputLayout3 = "20060102"
)

const (
	LocalOutputLayout = "200601021504"
)

var (
	localLayouts = []string{localInputLayout0, localInputLayout1, localInputLayout2, localInputLayout3}
	datetimeZero time.Time
)

func init() {
	datetime, err := time.ParseInLocation(localInputLayout0, "19700101T000000", time.UTC)
	if nil != err {
		panic(err)
	}
	datetimeZero = datetime
}

func ParseInputDatetime(str string) (date time.Time, err error) {
	for index := range localLayouts {
		rs, e := time.ParseInLocation(localLayouts[index], str, time.Local)
		if nil == e {
			return rs, nil
		}
	}
	return datetimeZero, errors.New(fmt.Sprintf(`parsing time "%s" fail!`, str))
}

func ParseDatetimeUTC(layout string, str string) (date time.Time, err error) {
	return time.ParseInLocation(layout, str, time.UTC)
}

func ParseDatetimeLocal(layout string, str string) (date time.Time, err error) {
	return time.ParseInLocation(layout, str, time.Local)
}

func ParseSVNDatetime(str string) (date time.Time, err error) {
	return time.ParseInLocation(time.RFC3339Nano, str, time.UTC)
}
