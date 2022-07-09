package svn

import (
	"encoding/xml"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"os/exec"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffToLast(path string, rev int) (l *model.DiffResult, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, err
	}
	last, err := log.GetLastLogEntry()
	if nil != err {
		return nil, err
	}
	if last.Revision == rev {
		return nil, nil
	}
	return QueryDiffBetween(path, rev, last.Revision)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffToNext(path string, rev int) (l *model.DiffResult, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, err
	}
	next, err := log.GetNextCommittedRevision(rev)
	if nil != err {
		return nil, err
	}
	if next == rev {
		return nil, nil
	}
	return QueryDiffBetween(path, rev, next)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffFromPrev(path string, rev int) (l *model.DiffResult, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, err
	}
	prev, err := log.GetPrevCommittedRevision(rev)
	if nil != err {
		return nil, err
	}
	if prev == rev {
		return nil, nil
	}
	return QueryDiffBetween(path, prev, rev)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffBetween(path string, revN int, revM int) (l *model.DiffResult, err error) {
	vStr := fmt.Sprintf("-r%d:%d", revN, revM)
	cmd := exec.Command(MainCmd, SubCmdDiff, vStr, ArgXml, ArgSummarize, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &model.DiffResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	return rs, nil
}
