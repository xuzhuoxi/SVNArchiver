package svn

const (
	CommandName = "svn"
)

const (
	SubInfo   = "info"
	SubLog    = "log"
	SubList   = "list"
	SubDiff   = "diff"
	SubExport = "export"
)

const (
	ArgXml       = "--xml"
	ArgRecursive = "--recursive"
	ArgSummarize = "--summarize"
)

type SvnCommit struct {
	Revision string `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}
