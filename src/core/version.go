// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
	"fmt"
)

func HandleVersion(ctx *env.VersionContext) {
	if nil == ctx {
		return
	}
	rs, err := svnversion.QueryVersion(ctx.TargetPath)
	if nil != err {
		fmt.Println("QueryVersion Error:", err)
		return
	}
	fmt.Println(rs)
}
