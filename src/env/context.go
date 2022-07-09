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

type ArchContext struct {
	TargetPath string
	ArchPath   string
	StartVer   int
	TargetVer  int
}

type ArchDateContext struct {
	TargetPath string
	ArchPath   string
	StartDate  time.Time
	TargetDate time.Time
}

func (ctx ArchDateContext) GetArchContext() *ArchContext {
	return nil
}
