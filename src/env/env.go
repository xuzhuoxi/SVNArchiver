// Create on 2022/7/8
// @author xuzhuoxi
package env

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"strings"
)

const DateCutIndex = 22

var (
	WildcardD  = []string{"{d}", "{D}"}   // 从0开始的分割顺序数。
	WildcardD0 = []string{"{d0}", "{D0}"} // 从1开始的分割顺序数。
	WildcardD1 = []string{"{d1}", "{D1}"} // 从1开始的分割顺序数。

	WildcardR  = []string{"{r}", "{R}"}   // 从0开始的分割顺序数。
	WildcardR0 = []string{"{r0}", "{R0}"} // 从1开始的分割顺序数。
	WildcardR1 = []string{"{r1}", "{R1}"} // 从1开始的分割顺序数。
)

func ToPrintDate(date string) string {
	index := mathx.MinInt(len(date), DateCutIndex)
	return date[:index]
}

func ReplaceWildcards(path string, wildcards []string, rep string) string {
	for _, wildcard := range wildcards {
		path = strings.ReplaceAll(path, wildcard, rep)
	}
	return path
}
