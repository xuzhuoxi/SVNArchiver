// Create on 2022/7/10
// @author xuzhuoxi
package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
)

var (
	ctxOutLog   *env.OutLogContext
	saveFuncMap map[string]func(archLog *env.ArchLog, path string)
)

func init() {
	saveFuncMap = make(map[string]func(archLog *env.ArchLog, path string))
	saveFuncMap["xml"] = archLogToXml
	saveFuncMap["json"] = archLogToJson
}

func TryHandleQueryLog(cmdFlags *env.CmdFlags) {
	if ctx := cmdFlags.GetLogContext(); nil != ctx {
		HandleSvnLog(ctx)
	}
}

func TryHandleArchXml(cmdFlags *env.CmdFlags) {
	archXml, err := cmdFlags.GetArchXml()
	if nil != err {
		Logger.Warnln(fmt.Sprintf("Handle Arch Xml Error[%s]!", err))
		return
	}
	if nil == archXml {
		return
	}
	tasks := archXml.GetTasks()
	if len(tasks) == 0 {
		return
	}
	if archXml.LogEnabled() {
		ctxOutLog = env.NewOutLogContext(archXml.Log)
	}
	for index := range tasks {
		tryHandleArchTask(tasks[index])
		Logger.Println()
	}
	if archXml.LogEnabled() {
		saveLog()
	}
}

func TryHandleArchContext(cmdFlags *env.CmdFlags) {
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

func tryHandleRevArch(archTask env.ArchTask) bool {
	ctx, err := archTask.GetRevArchContext()
	if nil != err {
		Logger.Warnln(fmt.Sprintf("RevArch Error[%s]!", err))
		return false
	}
	archPath := HandleRevArch(ctx)
	saveArchInfo(archTask.TaskId, archPath)
	return true
}

func tryHandleDateArch(archTask env.ArchTask) bool {
	ctx, err := archTask.GetDateArchContext()
	if nil != err {
		Logger.Warnln(fmt.Sprintf("DateArch Error[%s]!", err))
		return false
	}
	archPath := HandleDateArch(ctx)
	saveArchInfo(archTask.TaskId, archPath)
	return true
}

func tryHandleRevDiffArch(archTask env.ArchTask) bool {
	ctx, err := archTask.GetRevDiffArchContext()
	if nil != err {
		Logger.Warnln(fmt.Sprintf("RevDiffArch Error[%s]!", err))
		return false
	}
	archPath := HandleRevDiffArch(ctx)
	saveArchInfo(archTask.TaskId, archPath)
	return true
}

func tryHandleDateDiffArch(archTask env.ArchTask) bool {
	ctx, err := archTask.GetDateDiffArchContext()
	if nil != err {
		Logger.Warnln(fmt.Sprintf("DateDiffArch Error[%s]!", err))
		return false
	}
	archPath := HandleDateDiffArch(ctx)
	saveArchInfo(archTask.TaskId, archPath)
	return true
}

func saveArchInfo(id string, path string) {
	if "" == path || nil == ctxOutLog {
		return
	}
	code := GetCode(path, ctxOutLog.ArchXmlLog.CodeType)
	ctxOutLog.AppendLog(id, code, path)
}

func saveLog() {
	if nil == ctxOutLog {
		return
	}
	if f, ok := saveFuncMap[ctxOutLog.ArchXmlLog.FileType]; ok {
		f(ctxOutLog.ArchLog, ctxOutLog.ArchXmlLog.XmlValue)
	}
}

func archLogToXml(archLog *env.ArchLog, path string) {
	bs, err := xml.Marshal(archLog)
	if nil != err {
		Logger.Warnln(err)
	}
	filex.WriteFile(path, bs, os.ModePerm)
}

func archLogToJson(archLog *env.ArchLog, path string) {
	bs, err := json.Marshal(archLog)
	if nil != err {
		Logger.Warnln(err)
	}
	filex.WriteFile(path, bs, os.ModePerm)
}
