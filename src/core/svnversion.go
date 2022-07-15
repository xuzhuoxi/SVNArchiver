// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
)

func HandleVersion(ctx *env.QueryVersionContext) {
	if nil == ctx {
		return
	}
	Logger.Println(`Handle "svnversion" Command:`)
	rs, err := svnversion.QueryVersion(ctx.TargetPath)
	if nil != err {
		Logger.Warnln("QueryVersion Error:", err)
		return
	}
	Logger.Println(rs)
}
