// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
)

func HandleVersion(ctx *env.VersionContext) {
	if nil == ctx {
		return
	}
	fmt.Println(`Handle "svnversion" Command:`)
	rs, err := svnversion.QueryVersion(ctx.TargetPath)
	if nil != err {
		fmt.Println("QueryVersion Error:", err)
		return
	}
	fmt.Println(rs)
}
