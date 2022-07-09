// Create on 2022/7/9
// @author xuzhuoxi
package model

import "encoding/xml"

type ListResult struct {
	Name xml.Name             `xml:"lists"`
	List *ListResultEntryList `xml:"list"`
}

type ListResultEntryList struct {
	Path    string             `xml:"path,attr"`
	Entries []*ListResultEntry `xml:"entry"`
}

type ListResultEntry struct {
	Kind   string      `xml:"kind,attr"`
	Name   string      `xml:"name"`
	Size   int         `xml:"size"`
	Commit *CommitEntry `xml:"commit"`
}