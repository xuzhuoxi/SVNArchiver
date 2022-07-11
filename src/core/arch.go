// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/lib"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
	"strconv"
	"time"
)

var (
	titleDataArch = `"HandleDateArch"`
	titleRevArch  = `"HandleRevArch"`
)

func HandleDateArch(ctx *env.ArchDateContext) {
	if nil == ctx {
		return
	}

	Logger.Infoln(titleDataArch, ":")

	logResult, logRev, err := queryReversion(ctx.TargetPath, ctx.Date)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s ["queryReversion"] Error[%s]`, titleDataArch, err))
		return
	}

	Logger.Infoln(fmt.Sprintf(`%s Start: -d[%s] -r[%d] -target[%s]`, titleDataArch, ctx.DateString(), logRev.Reversion, ctx.TargetPath))
	archPath := getArchPathD(ctx.GetArchPath(), logResult, logRev.Reversion)
	archReversion(ctx.TargetPath, logRev.Reversion, archPath, titleDataArch)
	Logger.Infoln(fmt.Sprintf(`%s Finish: file=[%s]`, titleDataArch, archPath))
}

func HandleRevArch(ctx *env.ArchRevContext) {
	if nil == ctx {
		return
	}
	Logger.Infoln(titleRevArch, ":")

	logResult, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s ["svn log"] Error[%s]`, titleRevArch, err))
		return
	}
	logRev, err := logResult.GetCommittedRevision(ctx.Reversion)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s ["logResult.GetCommittedRevision"] Error[%s]`, titleRevArch, err))
		return
	}

	Logger.Infoln(fmt.Sprintf(`%s Start: -r[%d] -target[%s]`, titleRevArch, ctx.Reversion, ctx.TargetPath))
	archPath := getArchPathR(ctx.GetArchPath(), logRev.Reversion)
	archReversion(ctx.TargetPath, logRev.Reversion, archPath, titleRevArch)
	Logger.Infoln(fmt.Sprintf(`%s Finish: file=[%s]`, titleRevArch, archPath))
}

func queryReversion(targetPath string, date time.Time) (logResult *model.LogResult, logRev model.LogRev, err error) {
	logResult, err = svn.QueryLog(targetPath)
	if nil != err {
		return nil, model.LogRev{}, errors.New(fmt.Sprintf("QueryLog [%s]", err))
	}
	logRev, err = logResult.GetDateRevision(date)
	if nil != err {
		return nil, model.LogRev{}, errors.New(fmt.Sprintf("GetDateRevision [%s]", err))
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

func archReversion(targetPath string, reversion int, archPath string, errTitle string) {
	tempDir := getNextTempDir()
	err := svn.Export(targetPath, reversion, tempDir)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s \t["svn exprot"] [-r %d %s] Error[%s]`, errTitle, reversion, tempDir, err))
		return
	}
	Logger.Infoln(fmt.Sprintf(`%s \t["svn exprot"] [-r %d %s] succ.`, errTitle, reversion, tempDir))
	err = lib.Archive(tempDir, archPath, true)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s \t["tar"] [%s] Error[%s]`, errTitle, tempDir, err))
		return
	}
	Logger.Infoln(fmt.Sprintf(`%s \t["tar"] [%s] succ.`, errTitle, tempDir))
}
