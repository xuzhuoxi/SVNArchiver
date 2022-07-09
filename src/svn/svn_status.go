// Create on 2022/7/8
// @author xuzhuoxi
package svn

import (
	"os/exec"
	"encoding/xml"
	"fmt"
)

const (
	StatusItemNormal      = "normal"
	StatusItemUnversioned = "unversioned"
)

type StatusResult struct {
	Name         xml.Name            `xml:"status"`
	StatusTarget *StatusResultTarget `xml:"target"`
}

func (o *StatusResult) HandleResult() {
	o.StatusTarget.EntryList.Filter()
	o.StatusTarget.EntryList.Sort()
}

type StatusResultTarget struct {
	Path      string                `xml:"path,attr"`
	EntryList StatusResultEntryList `xml:"entry"`
}

func (o *StatusResultTarget) String() string {
	return fmt.Sprintf("{Path=%s, List=%v}", o.Path, o.EntryList)
}

type StatusResultEntryList []*StatusResultEntry

func (o StatusResultEntryList) Len() int {
	return len(o)
}

func (o StatusResultEntryList) Less(i, j int) bool {
	return o[j].WcStatus.Commit.Revision < o[i].WcStatus.Commit.Revision
}

func (o StatusResultEntryList) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o StatusResultEntryList) Sort() {
	//sort.Sort(o)
}

func (o StatusResultEntryList) Filter() {
	for index := o.Len() - 1; index >= 0; index -= 1 {
		if !o[index].IsNormalCommitted() {
			o = append(o[:index], o[index+1:]...)
		}
	}
}

type StatusResultEntry struct {
	Path     string          `xml:"path,attr"`
	WcStatus *StatusWcStatus `xml:"wc-status"`
}

func (o *StatusResultEntry) String() string {
	return fmt.Sprintf("{Path=%s, Status=%v}", o.Path, o.WcStatus)
}

func (o *StatusResultEntry) IsNormalCommitted() bool {
	if o.WcStatus == nil {
		return false
	}
	if o.WcStatus.Item != StatusItemNormal {
		return false
	}
	if o.WcStatus.Commit == nil {
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

func (o *StatusWcStatus) String() string {
	return fmt.Sprintf("{Item=%s, Revision=%d, Props=%s, Commit=%v}",
		o.Item, o.Revision, o.Props, o.Commit)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.status.html
// 功能：针对每个文件或目录，查询最新状态
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryStatus(path string) (l *StatusResult, err error) {
	cmd := exec.Command(MainCmd, SubCmdStatus, ArgVerbose, ArgXml, path)
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
