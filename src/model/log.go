// Create on 2022/7/9
// @author xuzhuoxi
package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/lib"
	"strings"
	"time"
)

type LogResult struct {
	Name       xml.Name          `xml:"log"`
	LogEntries []*LogResultEntry `xml:"logentry"` // 默认按reversion升序存放
}

func (r *LogResult) LogSize() int {
	return len(r.LogEntries)
}

func (r *LogResult) String() string {
	return fmt.Sprintf("{Name:%v, Size=%d}", r.Name, r.LogSize())
}

// 返回变动版本号
func (r *LogResult) GetDateRevision(date time.Time) (nearRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%v) Not Found!", date))
	}

	if date.Before(r.LogEntries[0].GetDate()) {
		return 0, errors.New(fmt.Sprintf("Revision(%v) Not Found!", date))
	}
	size := r.LogSize()
	lastDate := r.LogEntries[size-1].GetDate()
	if date.Equal(lastDate) || date.After(lastDate) {
		return r.LogEntries[size-1].Revision, nil
	}
	for index := 0; index < size; index += 1 {
		entryDate := r.LogEntries[index].GetDate()
		if date.Equal(entryDate) {
			return r.LogEntries[index].Revision, nil
		}
		if date.Before(entryDate) {
			if index-1 >= 0 {
				return r.LogEntries[index-1].Revision, nil
			}
			break
		}
	}
	return 0, errors.New(fmt.Sprintf("Revision(%v) Not Found!", date))
}

// 返回变动版本号
func (r *LogResult) GetCommittedRevision(revision int) (committedRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	if revision < r.LogEntries[0].Revision {
		return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
	}
	size := r.LogSize()
	if revision >= r.LogEntries[size-1].Revision {
		return r.LogEntries[size-1].Revision, nil
	}
	for index := 0; index < size; index += 1 {
		if revision == r.LogEntries[index].Revision {
			return r.LogEntries[index].Revision, nil
		}
		if revision < r.LogEntries[index].Revision {
			if index-1 >= 0 {
				return r.LogEntries[index-1].Revision, nil
			}
			break
		}
	}
	return 0, errors.New(fmt.Sprintf("Revision(%d) Not Found!", revision))
}

// 返回上一个版本号
func (r *LogResult) GetPrevCommittedRevision(revision int) (prevRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Prev Revision(%d) Not Found!", revision))
	}
	if revision < r.LogEntries[0].Revision {
		return 0, errors.New(fmt.Sprintf("Prev Revision(%d) Not Found!", revision))
	}
	size := r.LogSize()
	if revision > r.LogEntries[size-1].Revision {
		return r.LogEntries[size-1].Revision, nil
	}
	for index := 0; index < size; index += 1 {
		if revision <= r.LogEntries[index].Revision {
			if index-1 >= 0 {
				return r.LogEntries[index-1].Revision, nil
			}
			break
		}
	}
	return 0, errors.New(fmt.Sprintf("Prev Revision(%d) Not Found!", revision))
}

// 返回下一个版本号
func (r *LogResult) GetNextCommittedRevision(revision int) (nextRevision int, err error) {
	ln := len(r.LogEntries)
	if ln == 0 {
		return 0, errors.New(fmt.Sprintf("Next Revision(%d) Not Found!", revision))
	}
	if revision < r.LogEntries[0].Revision {
		return r.LogEntries[0].Revision, nil
	}
	size := r.LogSize()
	if revision >= r.LogEntries[size-1].Revision {
		return 0, errors.New(fmt.Sprintf("Next Revision(%d) Not Found!", revision))
	}
	for index := 0; index < size; index += 1 {
		if revision == r.LogEntries[index].Revision {
			if index+1 < len(r.LogEntries) {
				return r.LogEntries[index+1].Revision, nil
			}
			break
		}
		if revision < r.LogEntries[index].Revision {
			return r.LogEntries[index].Revision, nil
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

func (r *LogResult) GetFirstLogEntry() (e *LogResultEntry, err error) {
	if len(r.LogEntries) == 0 {
		return nil, errors.New(fmt.Sprintf("GetFirstLogEntry Not Found!"))
	}
	return r.LogEntries[0], nil
}

func (r *LogResult) GetLastLogEntry() (e *LogResultEntry, err error) {
	if len(r.LogEntries) == 0 {
		return nil, errors.New(fmt.Sprintf("GetLastLogEntry Not Found!"))
	}
	return r.LogEntries[r.LogSize()-1], nil
}

type LogResultEntry struct {
	Revision int             `xml:"revision,attr"`
	Author   string          `xml:"author"`
	Date     string          `xml:"date"`
	Paths    *LogResultPaths `xml:"paths"`
	Msg      string          `xml:"msg"`

	actions string
	date    *time.Time
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

func (l *LogResultEntry) GetDate() time.Time {
	if nil != l.date {
		return *l.date
	}
	d, _ := lib.ParseDatetimeByRFC3339Nano(l.Date)
	l.date = &d
	return d
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
