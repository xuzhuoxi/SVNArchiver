// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func HandleDateDiffArch(ctx *env.ArchDateDiffContext) {
	if nil == ctx {
		return
	}
	fmt.Println(fmt.Sprintf(`Handle "arch date diff[%s:%s]" Command:`, ctx.DateStartString(), ctx.DateTargetString()))
}

func HandleRevDiffArch(ctx *env.ArchRevDiffContext) {
	if nil == ctx {
		return
	}
	fmt.Println(fmt.Sprintf(`Handle "arch reversion diff[%d:%d]" Command:`, ctx.RevStart, ctx.RevTarget))
	rsQuery, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		fmt.Println("QueryList Error:", err)
		return
	}
	var start *model.LogResultEntry = nil
	var target *model.LogResultEntry = nil
	if ctx.ExitRange() {
		start, _ = rsQuery.GetLogEntry(ctx.RevStart)
		target, _ = rsQuery.GetLogEntry(ctx.RevTarget)
	} else {
		if ctx.ExitStart() {
			start, _ = rsQuery.GetLogEntry(ctx.RevStart)
			target, _ = rsQuery.GetLastLogEntry()
		}
		if ctx.ExitTarget() {
			start, _ = rsQuery.GetFirstLogEntry()
			target, _ = rsQuery.GetLogEntry(ctx.RevTarget)
		}
	}
	handleArchDiff(start, target)
}

func handleArchDiff(start *model.LogResultEntry, target *model.LogResultEntry) {
	fmt.Println("handleArchDiff:", start.Revision, target.Revision)
}
