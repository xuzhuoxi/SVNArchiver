package svn

import (
	"encoding/xml"
	"fmt"
	"os/exec"
)

// svn diff 命令主要针对文件

const (
	DiffItemAdded    = "added"
	DiffItemModified = "modified"
	DiffItemDeleted  = "deleted"
)

type DiffRoot struct {
	Name  xml.Name  `xml:"diff"`
	Entry InfoEntry `xml:"entry"`
}

type DiffPaths struct {
	Paths []DiffPath `xml:"path"`
}

type DiffPath struct {
	Item  string `xml:"item,attr"`
	Props string `xml:"props,attr"`
	Kind  string `xml:"kind,attr"`
	Path  string `xml:"path"`
}

func (p *DiffPath) IsAdded() bool {
	return p.Item == DiffItemAdded
}

func (p *DiffPath) IsModified() bool {
	return p.Item == DiffItemModified
}

func (p *DiffPath) IsDeleted() bool {
	return p.Item == DiffItemDeleted
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffSumPrevious(path string, targetRevision int) (l *DiffRoot, err error) {
	log, err := QueryLog(path)
	if nil != err {
		return nil, err
	}
	prev, err := log.GetCommittedRevisionPrevious(targetRevision)
	if nil != err {
		return nil, err
	}

	return QueryDiffSumRevision(path, targetRevision, prev)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
func QueryDiffSumRevision(path string, targetRevision int, baseRevision int) (l *DiffRoot, err error) {
	vStr := fmt.Sprintf("-r%d:%d", baseRevision, targetRevision)
	cmd := exec.Command(CommandName, SubDiff, vStr, ArgXml, ArgSummarize, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := DiffRoot{}
	err = xml.Unmarshal(out, &rs)
	if nil != err {
		return nil, err
	}
	return &rs, nil
}
