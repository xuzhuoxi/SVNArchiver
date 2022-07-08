package svn

import (
	"encoding/xml"
	"os/exec"
)

type ListRoot struct {
	Name xml.Name      `xml:"lists"`
	List ListEntryList `xml:"list"`
}

type ListEntryList struct {
	Path    string      `xml:"path,attr"`
	Entries []ListEntry `xml:"entry"`
}

type ListEntry struct {
	Kind   string      `xml:"kind,attr"`
	Name   string      `xml:"name"`
	Size   int         `xml:"size"`
	Commit CommitEntry `xml:"commit"`
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.list.html
func QueryList(path string, recursive bool) (l *ListRoot, err error) {
	var cmd *exec.Cmd
	if recursive {
		cmd = exec.Command(CommandName, SubList, ArgXml, ArgRecursive, path)
	} else {
		cmd = exec.Command(CommandName, SubList, ArgXml, path)
	}
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := ListRoot{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
