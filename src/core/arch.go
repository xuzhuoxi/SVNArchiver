// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
	"github.com/xuzhuoxi/infra-go/filex"
	"time"
)

func HandleDateArch(ctx *env.ArchDateContext) {
	if nil == ctx {
		return
	}

	_, reversion, err := queryReversion(ctx.TargetPath, ctx.Date)
	if nil != err {
		return
	}

	clearExportDir(ctx.GetArchPath())

	fmt.Println(fmt.Sprintf(`Handle "arch date[%s] reversion[%d]" Command:`, ctx.DateString(), reversion))
	archReversion(ctx.GetArchPath(), reversion, ctx.ArchPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, reversion, ctx.GetArchPath()))
}

func HandleRevArch(ctx *env.ArchRevContext) {
	if nil == ctx {
		return
	}
	clearExportDir(ctx.GetArchPath())

	fmt.Println(fmt.Sprintf(`Handle "arch reversion[%d]" Command:`, ctx.Reversion))
	archReversion(ctx.GetArchPath(), ctx.Reversion, ctx.ArchPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, ctx.Reversion, ctx.GetArchPath()))
}

func queryReversion(targetPath string, date time.Time) (logResult *model.LogResult, reversion int, err error) {
	logResult, err = svn.QueryLog(targetPath)
	if nil != err {
		fmt.Println("Query Log Error:", err)
		return nil, 0, err
	}
	reversion, err = logResult.GetDateRevision(date)
	if nil != err {
		fmt.Println("GetDateRevision Error:", err)
		return nil, 0, err
	}
	return logResult, reversion, nil
}

func archReversion(targetPath string, reversion int, archPath string) {
	svn.Export(targetPath, reversion, archPath)
}

func clearExportDir(dir string) {
	if !filex.IsDir(dir) {
		return
	}
	filex.RemoveAll(dir)
}
