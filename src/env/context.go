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
	Override   bool
}

type ArchRevDiffContext struct {
	TargetPath string
	ArchPath   string
	RevStart   int
	RevTarget  int
	Override   bool
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

type ArchDateContext struct {
	TargetPath string
	ArchPath   string
	Date       time.Time
	DateStr    string
	Override   bool
}

func (c *ArchDateContext) DateString() string {
	return ToPrintDate(c.DateStr)
}

type ArchDateDiffContext struct {
	TargetPath string
	ArchPath   string
	Override   bool

	DateStart    time.Time
	DateStartStr string
	ExistStart   bool

	DateTarget    time.Time
	DateTargetStr string
	ExistTarget   bool
}

func (c *ArchDateDiffContext) DateStartString() string {
	return ToPrintDate(c.DateStartStr)
}

func (c *ArchDateDiffContext) DateTargetString() string {
	return ToPrintDate(c.DateTargetStr)
}
