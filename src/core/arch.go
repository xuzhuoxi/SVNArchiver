// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/lib"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
	"github.com/xuzhuoxi/infra-go/filex"
	"strconv"
	"time"
)

func HandleDateArch(ctx *env.ArchDateContext) {
	if nil == ctx {
		return
	}

	logResult, logRev, err := queryReversion(ctx.TargetPath, ctx.Date)
	if nil != err {
		return
	}

	fmt.Println(fmt.Sprintf(`Handle "arch date[%s]reversion[%d][%s]" Command:`,
		ctx.DateString(), logRev.Reversion, ctx.TargetPath))
	archPath := getArchPathD(ctx.GetArchPath(), logResult, logRev.Reversion)
	archReversion(ctx.TargetPath, logRev.Reversion, archPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, logRev.Reversion, archPath))
}

func HandleRevArch(ctx *env.ArchRevContext) {
	if nil == ctx {
		return
	}

	logResult, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		return
	}
	logRev, err := logResult.GetCommittedRevision(ctx.Reversion)
	if nil != err {
		return
	}

	fmt.Println(fmt.Sprintf(`Handle "arch reversion[%d][%s]" Command:`,
		ctx.Reversion, ctx.TargetPath))
	archPath := getArchPathR(ctx.GetArchPath(), logRev.Reversion)
	archReversion(ctx.TargetPath, logRev.Reversion, archPath)
	fmt.Println(fmt.Sprintf(`Export reversion[%d] to:[%s]`, ctx.Reversion, archPath))
}

func queryReversion(targetPath string, date time.Time) (logResult *model.LogResult, logRev model.LogRev, err error) {
	logResult, err = svn.QueryLog(targetPath)
	if nil != err {
		fmt.Println("Query Log Error:", err)
		return nil, model.LogRev{}, err
	}
	logRev, err = logResult.GetDateRevision(date)
	if nil != err {
		fmt.Println("GetDateRevision Error:", err)
		return nil, model.LogRev{}, err
	}
	return logResult, logRev, nil
}

func getArchPathR(archPath string, fixRev int) string {
	archPath = env.ReplaceWildcards(archPath, env.WildcardR, strconv.Itoa(fixRev))
	return archPath
}

func getArchPathD(archPath string, logResult *model.LogResult, fixRev int) string {
	fixLogRev, _ := logResult.GetLogEntry(fixRev)
	archPath = env.ReplaceWildcards(archPath, env.WildcardD, fixLogRev.GetDateString())
	archPath = env.ReplaceWildcards(archPath, env.WildcardR, fixLogRev.GetReversionString())
	return archPath
}

func archReversion(targetPath string, reversion int, archPath string) {
	tempDir := getNextTempDir()
	svn.Export(targetPath, reversion, tempDir)
	lib.Archive(tempDir, archPath, true)
}

func clearExportDir(dir string) {
	if !filex.IsDir(dir) {
		return
	}
	filex.RemoveAll(dir)
}
