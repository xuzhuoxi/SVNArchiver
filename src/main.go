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
	tryHandleArchXml(cmdFlags)
	tryHandleArchContext(cmdFlags)

	core.ClearTempDir()
}

func tryHandleLog(cmdFlags *env.CmdFlags) {
	if ctx := cmdFlags.GetLogContext(); nil != ctx {
		core.HandleLog(ctx)
	}
}

func tryHandleArchXml(cmdFlags *env.CmdFlags) {
	archXml, err := cmdFlags.GetArchXml()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("Handle Arch Xml Error[%s]!", err))
		return
	}
	if nil == archXml {
		return
	}
	tasks := archXml.GetTasks()
	if len(tasks) == 0 {
		return
	}
	for index := range tasks {
		tryHandleArchTask(tasks[index])
		core.Logger.Println()
	}
}

func tryHandleArchContext(cmdFlags *env.CmdFlags) {
	if task, exist := cmdFlags.GetArchTask(); exist {
		tryHandleArchTask(task)
	}
}

func tryHandleArchTask(archTask env.ArchTask) {
	tryHandleRevArch(archTask)
	tryHandleDateArch(archTask)
	tryHandleRevDiffArch(archTask)
	tryHandleDateDiffArch(archTask)
}

func tryHandleRevArch(archTask env.ArchTask) {
	ctx, err := archTask.GetRevArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("RevArch Error[%s]!", err))
		return
	}
	core.HandleRevArch(ctx)
}

func tryHandleDateArch(archTask env.ArchTask) {
	ctx, err := archTask.GetDateArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("DateArch Error[%s]!", err))
		return
	}
	core.HandleDateArch(ctx)
}

func tryHandleRevDiffArch(archTask env.ArchTask) {
	ctx, err := archTask.GetRevDiffArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("RevDiffArch Error[%s]!", err))
		return
	}
	core.HandleRevDiffArch(ctx)
}

func tryHandleDateDiffArch(archTask env.ArchTask) {
	ctx, err := archTask.GetDateDiffArchContext()
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("DateDiffArch Error[%s]!", err))
		return
	}
	core.HandleDateDiffArch(ctx)
}
