// Create on 2022/7/9
// @author xuzhuoxi
package model

import "encoding/xml"

// svn diff 命令主要针对文件

const (
	DiffItemAdded    = "added"
	DiffItemModified = "modified"
	DiffItemDeleted  = "deleted"
)

type DiffResult struct {
	Name  xml.Name  `xml:"diff"`
	Paths DiffPaths `xml:"paths"`
}

type DiffPaths []*DiffPath

type DiffPath struct {
	Item  string `xml:"item,attr"`
	Props string `xml:"props,attr"`
	Kind  string `xml:"kind,attr"`
	Value string `xml:"path"`
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