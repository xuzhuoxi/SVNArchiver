// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
	"github.com/xuzhuoxi/infra-go/filex"
)

func HandleDateArch(ctx *env.ArchDateContext) {
	if nil == ctx {
		return
	}
	clearExportDir(ctx.ArchPath)
	rsQuery, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		fmt.Println("HandleDateArch on QueryLog Error:", err)
		return
	}
	reversion, err := rsQuery.GetDateRevision(ctx.Date)
	if nil != err {
		fmt.Println("HandleDateArch on GetDateRevision Error:", err)
		return
	}
	fmt.Println(fmt.Sprintf(`Handle "arch date[%s] reversion[%d]" Command:`, ctx.DateString(), reversion))
	svn.Export(ctx.TargetPath, reversion, ctx.ArchPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, reversion, ctx.TargetPath))
}

func HandleRevArch(ctx *env.ArchRevContext) {
	if nil == ctx {
		return
	}
	fmt.Println(fmt.Sprintf(`Handle "arch reversion[%d]" Command:`, ctx.Reversion))
	clearExportDir(ctx.ArchPath)
	svn.Export(ctx.TargetPath, ctx.Reversion, ctx.ArchPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, ctx.Reversion, ctx.TargetPath))
}

func clearExportDir(dir string) {
	if !filex.IsDir(dir) {
		return
	}
	filex.RemoveAll(dir)
}
