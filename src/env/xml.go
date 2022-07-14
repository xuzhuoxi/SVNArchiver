// Create on 2022/7/14
// @author xuzhuoxi
package env

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
)

type ArchTask struct {
	TargetPath string // 归档处理的svn目录，可以是svn仓库的非根目录。
	ArchPath   string // 归档文件保存路径，支持通配符。

	Reversion int // 完整归档时使用, 用于指定具体版本号，并使用该版本号(或向前最近的版本号)进行归档。
	RevDiffN  int // 差异归档时使用, 用于指定起始版本号。
	RevDiffM  int // 差异归档时使用, 用于指定结束版本号。

	Date      string // 完整归档时使用, 用于指定一个时间点，并使用该时间点上的版本号(或向前最近的版本号)进行归档。
	DateDiffN string // 差异归档时使用, 用于指定起始时间。
	DateDiffM string // 差异归档时使用, 用于指定结束时间。
}

func (t *ArchTask) GetRevDiffArchContext() (ctx *ArchRevDiffContext, err error) {
	if !t.isRevDiffArchCommand() {
		return nil, nil
	}
	return &ArchRevDiffContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath,
		RevStart: t.RevDiffN, RevTarget: t.RevDiffM}, nil
}

func (t *ArchTask) GetRevArchContext() (ctx *ArchRevContext, err error) {
	if !t.isRevArchCommand() {
		return nil, nil
	}
	return &ArchRevContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath,
		Reversion: t.Reversion}, nil
}

func (t *ArchTask) GetDateDiffArchContext() (ctx *ArchDateDiffContext, err error) {
	if !t.isDateDiffArchCommand() {
		return nil, nil
	}
	ctx = &ArchDateDiffContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath}
	if t.DateDiffN != "" {
		start, e := ParseInputDatetime(t.DateDiffN)
		if nil != e {
			return nil, e
		}
		ctx.DateStart, ctx.DateStartStr, ctx.ExistStart = start, t.DateDiffN, true
	}
	if t.DateDiffM != "" {
		target, e := ParseInputDatetime(t.DateDiffM)
		if nil != e {
			return nil, e
		}
		ctx.DateTarget, ctx.DateTargetStr, ctx.ExistTarget = target, t.DateDiffM, true
	}
	return
}

func (t *ArchTask) GetDateArchContext() (ctx *ArchDateContext, err error) {
	if !t.isDateArchCommand() {
		return nil, nil
	}
	ctx = &ArchDateContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath, DateStr: t.Date}
	date, e := ParseInputDatetime(t.Date)
	if nil != e {
		return nil, e
	}
	ctx.Date = date
	return
}

func (t *ArchTask) isRevArchCommand() bool {
	return t.ArchPath != "" && t.hasRevParams()
}

func (t *ArchTask) isRevDiffArchCommand() bool {
	return t.ArchPath != "" && t.hasRevDiffParams()
}

func (t *ArchTask) hasRevParams() bool {
	return t.Reversion > 0
}

func (t *ArchTask) hasRevDiffParams() bool {
	return t.RevDiffN > 0 || t.RevDiffM > 0
}

func (t *ArchTask) isDateArchCommand() bool {
	return t.ArchPath != "" && t.hasDateParams()
}

func (t *ArchTask) isDateDiffArchCommand() bool {
	return t.ArchPath != "" && t.hasDateDiffParams()
}

func (t *ArchTask) hasDateParams() bool {
	return t.Date != ""
}

func (t *ArchTask) hasDateDiffParams() bool {
	return t.DateDiffN != "" || t.DateDiffM != ""
}

type ArchXmlTask struct {
	Env    string `xml:"env"`
	Target string `xml:"target"`
	Arch   string `xml:"arch"`

	R  int    `xml:"r,attr"`
	R0 int    `xml:"r0,attr"`
	R1 int    `xml:"r1,attr"`
	D  string `xml:"d,attr"`
	D0 string `xml:"d0,attr"`
	D1 string `xml:"d1,attr"`
}

type ArchXmlTasks struct {
	Tasks []*ArchXmlTask `xml:"task"`
}

type ArchXml struct {
	MainEnv string        `xml:"main-env"`
	Tasks   *ArchXmlTasks `xml:"tasks"`
}

func (o *ArchXml) Init() {
	o.initMainEnv()
	o.initTasks()
}

func (o *ArchXml) initMainEnv() {
	runningRoot := osxu.GetRunningDir()
	if o.MainEnv == "" {
		o.MainEnv = runningRoot
		return
	}
	if filex.IsDir(o.MainEnv) {
		o.MainEnv = filex.FormatPath(o.MainEnv)
		return
	}
	o.MainEnv = filex.Combine(runningRoot, o.MainEnv)
}

func (o *ArchXml) initTasks() {
	xmlTasks := o.Tasks.Tasks
	for index, task := range xmlTasks {
		xmlTasks[index].Env = o.getTaskEnv(task)
		if filex.IsExist(task.Target) {
			xmlTasks[index].Target = filex.FormatPath(task.Target)
		} else {
			xmlTasks[index].Target = filex.Combine(xmlTasks[index].Env, task.Target)
		}
		if filex.IsExist(task.Arch) || filex.IsAbsFormat(task.Arch) {
			xmlTasks[index].Arch = filex.FormatPath(task.Arch)
		} else {
			xmlTasks[index].Arch = filex.Combine(xmlTasks[index].Env, task.Arch)
		}
	}
}

func (o *ArchXml) getTaskEnv(task *ArchXmlTask) string {
	if task.Env == "" {
		return o.MainEnv
	}
	if filex.IsDir(task.Env) {
		return filex.FormatPath(task.Env)
	}
	if filex.IsRelativeFormat(task.Env) {
		return filex.Combine(o.MainEnv, task.Env)
	}
	return filex.FormatPath(task.Env)
}

func (o *ArchXml) GetTasks() []ArchTask {
	xmlTasks := o.Tasks.Tasks
	rs := make([]ArchTask, len(xmlTasks))
	for index, xmlTask := range xmlTasks {
		task := ArchTask{TargetPath: xmlTask.Target, ArchPath: xmlTask.Arch,
			Reversion: xmlTask.R, RevDiffN: xmlTask.R0, RevDiffM: xmlTask.R1,
			Date: xmlTask.D, DateDiffN: xmlTask.D0, DateDiffM: xmlTask.D1}
		rs[index] = task
	}
	return rs
}
