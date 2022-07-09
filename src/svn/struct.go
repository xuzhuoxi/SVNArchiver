// Create on 2022/7/9
// @author xuzhuoxi
package svn

import "fmt"

type CommitEntry struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

func (c CommitEntry) String() string {
	return fmt.Sprintf("{Revision=%s, Author=%s, Date=%s}", c.Revision, c.Author, c.Date)
}
