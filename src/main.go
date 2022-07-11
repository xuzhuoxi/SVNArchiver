package main

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/core"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
)

func main() {
	core.InitLogger()
	cmdFlags, err := env.ParseFlags()
	if nil != err {
		core.Logger.Errorln("ParseFlogs Error[%s]!", err)
		return
	}

	core.ClearTempDir()

	tryHandleLog(cmdFlags)
	tryHandleRevArch(cmdFlags)
	tryHandleDateArch(cmdFlags)
	tryHandleRevDiffArch(cmdFlags)
	tryHandleDateDiffArch(cmdFlags)

	core.ClearTempDir()
}

func tryHandleLog(cmdFlags *env.CmdFlags) {
	if ctx := cmdFlags.GetLogContext(); nil != ctx {
		core.HandleLog(ctx)
	}
}

func tryHandleRevArch(cmdFlags *env.CmdFlags) {
	ctx, err := cmdFlags.GetRevArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("RevArch Error[%s]!", err))
		return
	}
	core.HandleRevArch(ctx)
}

func tryHandleDateArch(cmdFlags *env.CmdFlags) {
	ctx, err := cmdFlags.GetDateArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("DateArch Error[%s]!", err))
		return
	}
	core.HandleDateArch(ctx)
}

func tryHandleRevDiffArch(cmdFlags *env.CmdFlags) {
	ctx, err := cmdFlags.GetRevDiffArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("RevDiffArch Error[%s]!", err))
		return
	}
	core.HandleRevDiffArch(ctx)
}

func tryHandleDateDiffArch(cmdFlags *env.CmdFlags) {
	ctx, err := cmdFlags.GetDateDiffArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("DateDiffArch Error[%s]!", err))
		return
	}
	core.HandleDateDiffArch(ctx)
}
