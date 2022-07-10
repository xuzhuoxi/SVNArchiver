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

	fmt.Println(fmt.Sprintf(`Handle "arch date[%s]reversion[%d][%s]" Command:`,
		ctx.DateString(), reversion, ctx.TargetPath))
	archReversion(ctx.TargetPath, reversion, ctx.GetArchPath())
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, reversion, ctx.GetArchPath()))
}

func HandleRevArch(ctx *env.ArchRevContext) {
	if nil == ctx {
		return
	}
	fmt.Println(fmt.Sprintf(`Handle "arch reversion[%d][%s]" Command:`,
		ctx.Reversion, ctx.TargetPath))
	archReversion(ctx.TargetPath, ctx.Reversion, ctx.GetArchPath())
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
	tempDir := getNextTempDir()
	svn.Export(targetPath, reversion, tempDir)
}

func clearExportDir(dir string) {
	if !filex.IsDir(dir) {
		return
	}
	filex.RemoveAll(dir)
}
