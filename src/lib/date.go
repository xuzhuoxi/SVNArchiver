// Create on 2022/7/8
// @author xuzhuoxi
package lib

import (
	"errors"
	"fmt"
	"time"
)

const (
	layout0 = "20060102T150405"
	layout1 = "20060102T1504"
	layout2 = "20060102T15"
	layout3 = "20060102"
)

var (
	layouts      = []string{layout0, layout1, layout2, layout3}
	datetimeZero time.Time
)

func init() {
	datetime, err := ParseDatetimeByLayout("19710101T000000", layout0)
	if nil != err {
		panic(err)
	}
	datetimeZero = datetime
}

func ParseDatetime(str string) (date time.Time, err error) {
	for index := range layouts {
		rs, e := ParseDatetimeByLayout(str, layouts[index])
		if nil == e {
			return rs, nil
		}
	}
	return datetimeZero, errors.New(fmt.Sprintf(`parsing time "%s" fail!`, str))
}

func ParseDatetimeByLayout(str string, layout string) (date time.Time, err error) {
	return time.ParseInLocation(layout, str, time.Local)
}

func ParseDatetimeByRFC3339Nano(str string) (date time.Time, err error) {
	return time.ParseInLocation(time.RFC3339Nano, str, time.Local)
}
