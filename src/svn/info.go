package svn

import (
	"encoding/xml"
	"os/exec"
)

type InfoRoot struct {
	Name  xml.Name  `xml:"info"`
	Entry InfoEntry `xml:"entry"`
}

type InfoEntry struct {
	Kind        string         `xml:"kind,attr"`
	Path        string         `xml:"path,attr"`
	Revision    string         `xml:"revision,attr"`
	Url         string         `xml:"url"`
	RelativeUrl string         `xml:"relative-url"`
	Repository  InfoRepository `xml:"repository"`
	WcInfo      InfoWcInfo     `xml:"wc-info"`
	Commit      SvnCommit      `xml:"commit"`
}

type InfoRepository struct {
	Root string `xml:"root"`
	UUID string `xml:"uuid"`
}

type InfoWcInfo struct {
	WcRootAbsPath string `xml:"wcroot-abspath"`
	Schedule      string `xml:"schedule"`
	Depth         string `xml:"depth"`
	TextUpdated   string `xml:"text-updated"`
	CheckSum      string `xml:"checksum"`
}

func (r *InfoRoot) IsDir() bool {
	return r.Entry.WcInfo.CheckSum == ""
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.info.html
func QueryInfo(path string) (l *InfoRoot, err error) {
	cmd := exec.Command(CommandName, SubInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := InfoRoot{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
