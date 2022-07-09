package svn

import (
	"encoding/xml"
	"os/exec"
)

type InfoResult struct {
	Name  xml.Name         `xml:"info"`
	Entry *InfoResultEntry `xml:"entry"`
}

type InfoResultEntry struct {
	Kind        string                `xml:"kind,attr"`
	Path        string                `xml:"path,attr"`
	Revision    string                `xml:"revision,attr"`
	Url         string                `xml:"url"`
	RelativeUrl string                `xml:"relative-url"`
	Repository  *InfoResultRepository `xml:"repository"`
	WcInfo      *InfoResultWcInfo     `xml:"wc-info"`
	Commit      *CommitEntry          `xml:"commit"`
}

type InfoResultRepository struct {
	Root string `xml:"root"`
	UUID string `xml:"uuid"`
}

type InfoResultWcInfo struct {
	WcRootAbsPath string `xml:"wcroot-abspath"`
	Schedule      string `xml:"schedule"`
	Depth         string `xml:"depth"`
	TextUpdated   string `xml:"text-updated"`
	CheckSum      string `xml:"checksum"`
}

func (r *InfoResult) IsDir() bool {
	return r.Entry.WcInfo.CheckSum == ""
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.info.html
func QueryInfo(path string) (l *InfoResult, err error) {
	cmd := exec.Command(MainCmd, SubCmdInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := InfoResult{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
