// Create on 2022/7/8
// @author xuzhuoxi
package svn

import (
	"os/exec"
	"encoding/xml"
	"fmt"
	"sort"
)

const (
	StatusItemNormal      = "normal"
	StatusItemUnversioned = "unversioned"
)

type StatusResult struct {
	Name         xml.Name      `xml:"status"`
	StatusTarget *StatusTarget `xml:"target"`
}

func (r *StatusResult) HandleResult() {
	r.StatusTarget.EntryList.Filter()
	r.StatusTarget.EntryList.Sort()
}

type StatusTarget struct {
	Path      string          `xml:"path,attr"`
	EntryList StatusEntryList `xml:"entry"`
}

func (st *StatusTarget) String() string {
	return fmt.Sprintf("{Path=%s, List=%v}", st.Path, st.EntryList)
}

type StatusEntryList []StatusEntry

func (l StatusEntryList) Len() int {
	return len(l)
}

func (l StatusEntryList) Less(i, j int) bool {
	return l[j].WcStatus.Commit.Revision < l[i].WcStatus.Commit.Revision
}

func (l StatusEntryList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l StatusEntryList) Sort() {
	sort.Sort(l)
}

func (l StatusEntryList) Filter() {
	for index := l.Len() - 1; index >= 0; index -= 1 {
		if !l[index].IsNormalCommitted() {
			l = append(l[:index], l[index+1:]...)
		}
	}
}

type StatusEntry struct {
	Path     string          `xml:"path,attr"`
	WcStatus *StatusWcStatus `xml:"wc-status"`
}

func (se *StatusEntry) String() string {
	return fmt.Sprintf("{Path=%s, Status=%v}", se.Path, se.WcStatus)
}

func (se *StatusEntry) IsNormalCommitted() bool {
	if se.WcStatus == nil {
		return false
	}
	if se.WcStatus.Item != StatusItemNormal {
		return false
	}
	if se.WcStatus.Commit == nil {
		return false
	}
	return true
}

type StatusWcStatus struct {
	Item     string       `xml:"item,attr"`
	Revision int          `xml:"revision,attr"`
	Props    string       `xml:"props,attr"`
	Commit   *CommitEntry `xml:"commit"`
}

func (wc *StatusWcStatus) String() string {
	return fmt.Sprintf("{Item=%s, Revision=%d, Props=%s, Commit=%v}",
		wc.Item, wc.Revision, wc.Props, wc.Commit)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.status.html
// 功能：针对每个文件或目录，查询最新状态
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryStatus(path string) (l *StatusResult, err error) {
	cmd := exec.Command(CommandName, SubStatus, ArgVerbose, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &StatusResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	rs.HandleResult()
	return rs, nil
}
