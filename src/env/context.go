// Create on 2022/7/8
// @author xuzhuoxi
package env

import (
	"strconv"
	"time"
)

type VersionContext struct {
	TargetPath string
}

type LogContext struct {
	TargetPath string
	LogSize    int
}

type ArchRevContext struct {
	TargetPath string
	ArchPath   string
	Reversion  int
}

func (o *ArchRevContext) GetArchPath() string {
	return o.ArchPath
}

type ArchRevDiffContext struct {
	TargetPath string
	ArchPath   string
	RevStart   int
	RevTarget  int
}

func (c *ArchRevDiffContext) ExitStart() bool {
	return c.RevStart != 0
}

func (c *ArchRevDiffContext) ExitTarget() bool {
	return c.RevTarget != 0
}

func (c *ArchRevDiffContext) ExitRange() bool {
	return c.RevTarget != 0 && c.RevStart != 0
}

func (c *ArchRevDiffContext) RevStartString() string {
	if 0 == c.RevStart {
		return ""
	}
	return strconv.Itoa(c.RevStart)
}

func (c *ArchRevDiffContext) RevTargetString() string {
	if 0 == c.RevTarget {
		return ""
	}
	return strconv.Itoa(c.RevTarget)
}

func (o *ArchRevDiffContext) GetArchPath() string {
	return o.ArchPath
}

type ArchDateContext struct {
	TargetPath string
	ArchPath   string
	Date       time.Time
	DateStr    string
}

func (o *ArchDateContext) DateString() string {
	return ToPrintDate(o.DateStr)
}

func (o *ArchDateContext) GetArchPath() string {
	return o.ArchPath
}

type ArchDateDiffContext struct {
	TargetPath string
	ArchPath   string

	DateStart    time.Time
	DateStartStr string
	ExistStart   bool

	DateTarget    time.Time
	DateTargetStr string
	ExistTarget   bool
}

func (o *ArchDateDiffContext) DateStartString() string {
	return ToPrintDate(o.DateStartStr)
}

func (o *ArchDateDiffContext) DateTargetString() string {
	return ToPrintDate(o.DateTargetStr)
}

func (o *ArchDateDiffContext) GetArchPath() string {
	return o.ArchPath
}
