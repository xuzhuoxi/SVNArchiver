package svn

import (
	"encoding/xml"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"os/exec"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffToLast(path string, rev int) (l *model.DiffResult, revN, revM int, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, 0, 0, err
	}
	last, err := log.GetLastLogEntry()
	if nil != err {
		return nil, 0, 0, err
	}
	if last.Revision == rev {
		return nil, 0, 0, nil
	}

	l, err = QueryDiffBetween(path, rev, last.Revision)
	revN, revM = rev, last.Revision
	return
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffToNext(path string, rev int) (l *model.DiffResult, revN, revM int, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, 0, 0, err
	}
	next, err := log.GetNextCommittedRevision(rev)
	if nil != err {
		return nil, 0, 0, err
	}
	if next == rev {
		return nil, 0, 0, nil
	}
	l, err = QueryDiffBetween(path, rev, next)
	revN, revM = rev, next
	return
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffFromPrev(path string, rev int) (l *model.DiffResult, revN, revM int, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, 0, 0, err
	}
	prev, err := log.GetPrevCommittedRevision(rev)
	if nil != err {
		return nil, 0, 0, err
	}
	if prev == rev {
		return nil, 0, 0, nil
	}
	l, err = QueryDiffBetween(path, prev, rev)
	revN, revM = prev, rev
	return
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
