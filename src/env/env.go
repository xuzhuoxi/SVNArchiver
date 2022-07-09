// Create on 2022/7/8
// @author xuzhuoxi
package env

import "github.com/xuzhuoxi/infra-go/mathx"

const DateCutIndex = 22

func ToPrintDate(date string) string {
	index := mathx.MinInt(len(date), DateCutIndex)
	return date[:index]
}
