package svn

import (
	"fmt"
)

const (
	CommandName = "svn"
)

// 针对文件的子命令
const (
	SubInfo = "info"
	SubLog  = "log"
	SubDiff = "diff"
)

const (
	SubList   = "list"
	SubExport = "export"
	SubStatus = "status"
)

const (
	ArgXml       = "--xml"
	ArgRecursive = "--recursive"
	ArgSummarize = "--summarize"
	ArgVerbose   = "--verbose"
	ArgQuiet     = "--quiet"
)

type CommitEntry struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

func (c CommitEntry) String() string {
	return fmt.Sprintf("{Revision=%s, Author=%s, Date=%s}", c.Revision, c.Author, c.Date)
}
