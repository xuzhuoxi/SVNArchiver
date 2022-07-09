package svn

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
	"strings"
)

type LogResult struct {
	Name       xml.Name          `xml:"log"`
	LogEntries []*LogResultEntry `xml:"logentry"`
}

func (r *LogResult) LogSize() int {
	return len(r.LogEntries)
}

func (r *LogResult) String() string {
	return fmt.Sprintf("{Name:%v, Size=%d}", r.Name, r.LogSize())
}

// 返回变动版本号
func (r *LogResult) GetCommittedRevision(revision int) (nearRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	for index := range r.LogEntries {
		if r.LogEntries[index].Revision <= revision {
			return r.LogEntries[index].Revision, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
}

// 返回变动版本号
func (r *LogResult) GetPrevCommittedRevision(revision int) (previousRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	for index := range r.LogEntries {
		if r.LogEntries[index].Revision <= revision {
			if index < len(r.LogEntries)-1 {
				return r.LogEntries[index+1].Revision, nil
			}
		}
	}
	return 0, errors.New(fmt.Sprintf("Previous Revision(%d) Not Found!", revision))
}

func (r *LogResult) GetCommittedLogEntry(revision int) (e *LogResultEntry, err error) {
	committedRevision, err := r.GetCommittedRevision(revision)
	if nil != err {
		return
	}
	return r.GetLogEntry(committedRevision)
}

func (r *LogResult) GetLogEntry(committedRevision int) (e *LogResultEntry, err error) {
	for index := range r.LogEntries {
		if r.LogEntries[index].Revision == committedRevision {
			return r.LogEntries[index], nil
		}
	}
	err = errors.New(fmt.Sprintf("GetLogEntry(committedRevision=%d) Not Found!", committedRevision))
	return
}

type LogResultEntry struct {
	Revision int             `xml:"revision,attr"`
	Author   string          `xml:"author"`
	Date     string          `xml:"date"`
	Paths    *LogResultPaths `xml:"paths"`
	Msg      string          `xml:"msg"`

	actions string
}

func (l *LogResultEntry) GetActions() string {
	if l.actions != "" {
		return l.actions
	}
	if l.Paths.PathSize() == 0 {
		return ""
	}
	actions := ""
	for _, path := range l.Paths.Paths {
		if strings.Contains(actions, path.Action) {
			continue
		}
		actions = actions + path.Action
	}
	l.actions = actions
	return actions
}

func (l LogResultEntry) String() string {
	return fmt.Sprintf("{Revision:%d, Author:%s, Date:%s, Msg:%s}", l.Revision, l.Author, l.Date, l.Msg)
}

type LogResultPaths struct {
	Paths []*LogResultPath `xml:"path"`
}

func (o *LogResultPaths) PathSize() int {
	return len(o.Paths)
}

type LogResultPath struct {
	Action   string `xml:"action,attr"`
	PropMods bool   `xml:"prop-mods,attr"`
	TextMods bool   `xml:"text-mods,attr"`
	Kind     string `xml:"kind,attr"`
	XmlValue string `xml:",innerxml"`
}

func (o LogResultPath) String() string {
	return fmt.Sprintf("{Action=%s, Prop=%v, Text=%v, Kind=%s, Value=%s}",
		o.Action, o.PropMods, o.TextMods, o.Kind, o.XmlValue)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryLog(path string) (l *LogResult, err error) {
	rsVersion, err := svnversion.QueryVersion(path)
	if nil != err {
		return nil, err
	}
	verInfo := fmt.Sprintf("%d:%d", rsVersion.Min, rsVersion.Max)
	cmd := exec.Command(MainCmd, SubCmdLog, ArgVerbose, ArgRevision, verInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &LogResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	return rs, nil
}
