// Create on 2022/7/8
// @author xuzhuoxi
package lib

import "time"

const (
	DatetimeLayout = "20060102T150405"
)

var (
	DatetimeZero time.Time
)

func init() {
	datetime, err := ParseDatetime("19710101T000000")
	if nil != err {
		panic(err)
	}
	DatetimeZero = datetime
}

func ParseDatetime(str string) (date time.Time, err error) {
	return time.ParseInLocation(DatetimeLayout, str, time.Local)
}

func ParseDatetimeByRFC3339Nano(str string) (date time.Time, err error) {
	return time.ParseInLocation(time.RFC3339Nano, str, time.Local)
}
