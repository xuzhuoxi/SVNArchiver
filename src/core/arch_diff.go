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
	"os"
	"strconv"
)

var (
	titleDataDiffArch = `"HandleDateDiffArch"`
	titleRevDiffArch  = `"HandleRevDiffArch"`
)

func HandleDateDiffArch(ctx *env.ArchDateDiffContext) {
	if nil == ctx {
		return
	}

	Logger.Infoln(titleDataDiffArch, ":")

	logResult, logRevN, logRevM, err := getRev(ctx)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s ["getEvn"] Error[%s]`, titleDataDiffArch, err))
		return
	}

	diffResult, fixRevN, fixRevM, err := queryDiff(ctx.TargetPath, logRevN.Reversion, logRevM.Reversion)
	if err != nil {
		Logger.Warnln(fmt.Sprintf(`%s ["queryDiff diff[%s:%s]"] Error[%s]`,
			titleDataDiffArch, ctx.DateStartString(), ctx.DateTargetString(), err))
		return
	}

	Logger.Infoln(fmt.Sprintf(`%s Start: diff[%s:%s] -target[%s]`, titleDataDiffArch, ctx.DateStartString(), ctx.DateTargetString(), ctx.TargetPath))
	archPath := getArchDiffPathD(ctx.ArchPath, logResult, fixRevN, fixRevM)
	handleArchDiff2(ctx.TargetPath, diffResult, fixRevN, fixRevM, archPath, titleDataDiffArch)
	Logger.Infoln(fmt.Sprintf(`%s Finish: diff[%d:%d] file=[%s]`, titleDataDiffArch, fixRevN, fixRevM, archPath))
}

func HandleRevDiffArch(ctx *env.ArchRevDiffContext) {
	if nil == ctx {
		return
	}

	diffResult, fixRevN, fixRevM, err := queryDiff(ctx.TargetPath, ctx.RevStart, ctx.RevTarget)
	if err != nil {
		Logger.Warnln(fmt.Sprintf(`%s ["queryDiff diff[%s:%s]"] Error[%s]`,
			titleRevDiffArch, ctx.RevStartString(), ctx.RevTargetString(), err))
		return
	}

	Logger.Infoln(fmt.Sprintf(`%s Start: diff[%s:%s] -target[%s]`, titleRevDiffArch, ctx.RevStartString(), ctx.RevTargetString(), ctx.TargetPath))
	archPath := getArchDiffPathR(ctx.ArchPath, fixRevN, fixRevM)
	handleArchDiff2(ctx.TargetPath, diffResult, fixRevN, fixRevM, archPath, titleRevDiffArch)
	Logger.Infoln(fmt.Sprintf(`%s Finish: diff[%d:%d] file=[%s]`, titleRevDiffArch, fixRevN, fixRevM, archPath))
}

func getRev(ctx *env.ArchDateDiffContext) (logResult *model.LogResult, revN, revM model.LogRev, err error) {
	logResult, err = svn.QueryLog(ctx.TargetPath)
	if nil != err {
		return
	}

	if ctx.ExistStart {
		rev, err := logResult.GetDateRevision(ctx.DateStart)
		if nil != err {
			return logResult, model.LogRev{}, model.LogRev{}, err
		}
		revN = rev
	}
	if ctx.ExistTarget {
		rev, err := logResult.GetDateRevision(ctx.DateTarget)
		if nil != err {
			return logResult, model.LogRev{}, model.LogRev{}, err
		}
		revM = rev
	}
	return logResult, revN, revM, nil
}

func queryDiff(targetPath string, revN, revM int) (l *model.DiffResult, fixRevN, fixRevM int, err error) {
	if revN == 0 {
		return svn.QueryDiffFromPrev(targetPath, revM)
	}
	if revM == 0 {
		return svn.QueryDiffToLast(targetPath, revN)
	}

	l, err = svn.QueryDiffBetween(targetPath, revN, revM)
	fixRevN, fixRevM = revN, revM
	return
}

func getArchDiffPathR(archPath string, fixRevN, fixRevM int) string {
	archPath = env.ReplaceWildcards(archPath, env.WildcardR0, strconv.Itoa(fixRevN))
	archPath = env.ReplaceWildcards(archPath, env.WildcardR1, strconv.Itoa(fixRevM))
	return archPath
}

func getArchDiffPathD(archPath string, logResult *model.LogResult, fixRevN, fixRevM int) string {
	fixLogRevN, _ := logResult.GetLogEntry(fixRevN)
	fixLogRevM, _ := logResult.GetLogEntry(fixRevM)
	archPath = env.ReplaceWildcards(archPath, env.WildcardD0, fixLogRevN.GetDateString())
	archPath = env.ReplaceWildcards(archPath, env.WildcardD1, fixLogRevM.GetDateString())
	archPath = env.ReplaceWildcards(archPath, env.WildcardR0, fixLogRevN.GetReversionString())
	archPath = env.ReplaceWildcards(archPath, env.WildcardR1, fixLogRevM.GetReversionString())
	return archPath
}

// 效率低
// 这个方法的逻辑如下
// 1. 取差异列表
// 2. 遍历差异列表中目录，创建目录
// 3. 遍历差异列表中文件，使用"svn export"命令导出
func handleArchDiff(targetPath string, diffResult *model.DiffResult, revN, revM int, archPath string, errTitle string) {
	baseLen := len(targetPath)
	tempDir := genNextTempDir()
	for _, v := range diffResult.Paths.Paths {
		if v.IsDeleted() || v.IsFile() {
			continue
		}
		relativePath := v.XmlValue[baseLen:]
		archPath := filex.Combine(tempDir, relativePath)
		if filex.IsDir(archPath) {
			continue
		}
		os.MkdirAll(archPath, os.ModePerm)
	}
	for _, v := range diffResult.Paths.Paths {
		if v.IsDeleted() || v.IsDir() {
			continue
		}
		relativePath := v.XmlValue[baseLen:]
		archPath := filex.Combine(tempDir, relativePath)
		err := svn.Export(v.XmlValue, revM, archPath)
		if nil != err {
			Logger.Warnln(fmt.Sprintf(`%s  ["svn exprot"] [-r%d %s] Error[%s]`, errTitle, revM, archPath, err))
			continue
		}
		Logger.Warnln(fmt.Sprintf(`%s  ["svn exprot"] [-r%d %s] succ.`, errTitle, revM, archPath))
	}
	err := lib.Archive(tempDir, archPath, true)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s  ["tar"] [%s] Error[%s]`, errTitle, tempDir, err))
		return
	}
	Logger.Infoln(fmt.Sprintf(`%s  ["tar"] [%s] succ.`, errTitle, tempDir))
}

// 效率高
// 这个方法的逻辑如下
// 1. 导出目标版本号的全部
// 2. 根据差异列表移动文件
func handleArchDiff2(targetPath string, diffResult *model.DiffResult, revN, revM int, archPath string, errTitle string) {
	baseLen := len(targetPath)
	tempDir1 := getNextTempDir()
	err := svn.Export(targetPath, revM, tempDir1)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s  ["svn exprot"] [-r%d %s] Error[%s]`, errTitle, revM, tempDir1, err))
		return
	}
	Logger.Infoln(fmt.Sprintf(`%s  ["svn exprot"] [-r%d %s] succ.`, errTitle, revM, tempDir1))
	tempDir2 := genNextTempDir()
	Logger.Infoln(fmt.Sprintf("%s  [\"copy\"]:", errTitle))
	Logger.Println(fmt.Sprintf("\t%s", tempDir1))
	Logger.Println(fmt.Sprintf("\t  => %s", tempDir2))
	for _, v := range diffResult.Paths.Paths {
		if v.IsDeleted() || v.IsDir() {
			continue
		}
		relativePath := filex.FormatPath(v.XmlValue[baseLen+1:])
		srcPath := filex.Combine(tempDir1, relativePath)
		dstPath := filex.Combine(tempDir2, relativePath)
		filex.MoveAuto(srcPath, dstPath, os.ModePerm)
		Logger.Println(fmt.Sprintf("\t  -file: %s", relativePath))
	}
	err = lib.Archive(tempDir2, archPath, true)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`%s  ["tar"] [%s] Error[%s]`, errTitle, tempDir2, err))
		return
	}
	Logger.Infoln(fmt.Sprintf(`%s  ["tar"] [%s] succ.`, errTitle, tempDir2))
}
