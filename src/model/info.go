// Create on 2022/7/9
// @author xuzhuoxi
package model

import "encoding/xml"

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
