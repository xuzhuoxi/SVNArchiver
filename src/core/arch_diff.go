// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
)

func HandleDateDiffArch(ctx *env.ArchDateDiffContext) {
	if nil == ctx {
		return
	}

	_, revN, revM, err := getRev(ctx)
	if nil != err {
		fmt.Println(fmt.Sprintf(`Handle "arch date diff" query reversion error:[%s]`, err))
		return
	}

	diffResult, fixRevN, fixRevM, err := queryDiff(ctx.TargetPath, revN, revM)
	if err != nil {
		fmt.Println(fmt.Sprintf(`Handle "arch date diff[%s:%s]" Command:`, ctx.DateStartString(), ctx.DateTargetString()))
		return
	}

	clearExportDir(ctx.GetArchPath())

	fmt.Println(fmt.Sprintf(`Handle "arch date diff[%s:%s]" Command:`, ctx.DateStartString(), ctx.DateTargetString()))
	handleArchDiff(ctx.TargetPath, diffResult, fixRevN, fixRevM, ctx.GetArchPath())
	fmt.Println(fmt.Sprintf(`Export date diff[%s:%s] to:[%s]`, ctx.DateStartString(), ctx.DateTargetString(), ctx.GetArchPath()))
}

func HandleRevDiffArch(ctx *env.ArchRevDiffContext) {
	if nil == ctx {
		return
	}

	diffResult, fixRevN, fixRevM, err := queryDiff(ctx.TargetPath, ctx.RevStart, ctx.RevTarget)
	if err != nil {
		fmt.Println(fmt.Sprintf(`Handle "arch reversion diff[%s:%s]" Command:`, ctx.RevStartString(), ctx.RevTargetString()))
		return
	}

	clearExportDir(ctx.GetArchPath())

	fmt.Println(fmt.Sprintf(`Handle "arch reversion diff[%s:%s]" Command:`, ctx.RevStartString(), ctx.RevTargetString()))
	handleArchDiff(ctx.TargetPath, diffResult, fixRevN, fixRevM, ctx.GetArchPath())
	fmt.Println(fmt.Sprintf(`Export reversion diff[%s:%s] to:[%s]`, ctx.RevStartString(), ctx.RevTargetString(), ctx.GetArchPath()))
}

func getRev(ctx *env.ArchDateDiffContext) (logResult *model.LogResult, revN, revM int, err error) {
	logResult, err = svn.QueryLog(ctx.TargetPath)
	if nil != err {
		return
	}

	if ctx.ExistStart {
		rev, err := logResult.GetDateRevision(ctx.DateStart)
		if nil != err {
			return logResult, 0, 0, err
		}
		revN = rev
	}
	if ctx.ExistTarget {
		rev, err := logResult.GetDateRevision(ctx.DateTarget)
		if nil != err {
			return logResult, 0, 0, err
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

func handleArchDiff(targetPath string, diffResult *model.DiffResult, revN, revM int, archDir string) {
	baseLen := len(targetPath)
	for _, v := range diffResult.Paths.Paths {
		if v.IsDeleted() || v.IsFile() {
			continue
		}
		relativePath := v.XmlValue[baseLen:]
		archPath := filex.Combine(archDir, relativePath)
		if filex.IsDir(archPath) {
			continue
		}
		os.MkdirAll(archPath, os.ModePerm)
	}
	fmt.Println(fmt.Sprintf(`"$Base"=%s`, archDir))
	for _, v := range diffResult.Paths.Paths {
		if v.IsDeleted() || v.IsDir() {
			continue
		}
		relativePath := v.XmlValue[baseLen:]
		archPath := filex.Combine(archDir, relativePath)
		fmt.Println(fmt.Sprintf("Export: %s", filex.Combine("$Base", relativePath)))
		svn.Export(v.XmlValue, revM, archPath)
	}
}
