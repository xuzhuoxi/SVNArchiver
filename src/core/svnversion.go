// Package core
// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func HandleLocalVersion(ctx *env.QueryVersionContext) {
	if nil == ctx {
		return
	}
	Logger.Println(`HandleLocalVersion with command["svnversion"]:`)
	rs, err := svn.QueryLocalVersion(ctx.TargetPath)
	if nil != err {
		Logger.Warnln("QueryLocalVersion Error:", err)
		return
	}
	Logger.Println(rs)
}
