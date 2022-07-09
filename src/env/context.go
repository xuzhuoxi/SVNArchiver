// Create on 2022/7/8
// @author xuzhuoxi
package env

import (
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

type ArchDateContext struct {
	TargetPath string
	ArchPath   string
	Date       time.Time
	DateStr    string
}

func (o *ArchDateContext) DateString() string {
	return ToPrintDate(o.DateStr)
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
