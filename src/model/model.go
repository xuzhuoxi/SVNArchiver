// Create on 2022/7/9
// @author xuzhuoxi
package model

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
)

type CommitEntry struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

func (o *CommitEntry) DateString() string {
	return env.ToPrintDate(o.Date)
}

func (o CommitEntry) String() string {
	return fmt.Sprintf("{Revision=%s, Author=%s, Date=%s}", o.Revision, o.Author, o.DateString())
}
