// Create on 2022/7/9
// @author xuzhuoxi
package model

import (
	"encoding/xml"
	"fmt"
)

// svn diff 命令主要针对文件

const (
	DiffItemAdded    = "added"
	DiffItemModified = "modified"
	DiffItemDeleted  = "deleted"
	DiffKindFile     = "file"
	DiffKindDir      = "dir"
)

type DiffResult struct {
	Name  xml.Name   `xml:"diff"`
	Paths *DiffPaths `xml:"paths"`
}

type DiffPaths struct {
	Paths []*DiffPath `xml:"path"`
}

func (o DiffPaths) String() string {
	return fmt.Sprintf("{DiffPaths.Len=%d}", len(o.Paths))
}

type DiffPath struct {
	Item     string `xml:"item,attr"`
	Props    string `xml:"props,attr"`
	Kind     string `xml:"kind,attr"`
	XmlValue string `xml:",innerxml"`
}

func (p *DiffPath) String() string {
	return fmt.Sprintf(`{item=%s, props=%s, kind=%s, value=%s}`, p.Item, p.Props, p.Kind, p.XmlValue)
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

func (p *DiffPath) IsFile() bool {
	return p.Kind == DiffKindFile
}

func (p *DiffPath) IsDir() bool {
	return p.Kind == DiffKindDir
}
