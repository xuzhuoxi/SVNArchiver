package svn

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
)

type LogRoot struct {
	Name       xml.Name   `xml:"log"`
	LogEntries []LogEntry `xml:"logentry"`
}

// 返回变动版本号
func (l *LogRoot) GetCommittedRevision(revision int) (nearRevision int, err error) {
	ln := len(l.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	for index := range l.LogEntries {
		if l.LogEntries[index].Revision <= revision {
			return l.LogEntries[index].Revision, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
}

// 返回变动版本号
func (l *LogRoot) GetCommittedRevisionPrevious(revision int) (previousRevision int, err error) {
	ln := len(l.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	for index := range l.LogEntries {
		if l.LogEntries[index].Revision <= revision {
			if index < len(l.LogEntries)-1 {
				return l.LogEntries[index+1].Revision, nil
			}
		}
	}
	return 0, errors.New(fmt.Sprintf("Previous Revision(%d) Not Found!", revision))
}

func (l *LogRoot) GetLogEntry(commitRevision int) (e LogEntry, err error) {
	for index := range l.LogEntries {
		if l.LogEntries[index].Revision == commitRevision {
			return l.LogEntries[index], nil
		}
	}
	err = errors.New(fmt.Sprintf("LogEntry(commitRevision=%d) Not Found!", commitRevision))
	return
}

func (l *LogRoot) GetCommittedLogEntry(revision int) (e LogEntry, err error) {
	commitRevision, err := l.GetCommittedRevision(revision)
	if nil != err {
		return
	}
	return l.GetLogEntry(commitRevision)
}

type LogEntry struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
	Msg      string `xml:"msg"`
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryLog(path string) (l *LogRoot, err error) {
	cmd := exec.Command(CommandName, SubLog, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := LogRoot{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
