package svn

import (
	"encoding/xml"
	"os/exec"
)

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

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.list.html
func QueryList(path string, recursive bool) (l *ListResult, err error) {
	var cmd *exec.Cmd
	if recursive {
		cmd = exec.Command(MainCmd, SubCmdList, ArgXml, ArgRecursive, path)
	} else {
		cmd = exec.Command(MainCmd, SubCmdList, ArgXml, path)
	}
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := ListResult{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
