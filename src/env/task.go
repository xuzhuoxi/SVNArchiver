// Create on 2022/7/14
// @author xuzhuoxi
package env

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
)

type ArchTask struct {
	TargetPath string
	ArchPath   string

	Reversion int // 版本导出
	RevDiffN  int // 版本差异前
	RevDiffM  int // 版本差异后

	Date      string // 版本时间导出
	DateDiffN string // 版本时间差异前
	DateDiffM string // 版本时间差异后
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
	Env   string        `xml:"env"`
	Tasks *ArchXmlTasks `xml:"tasks"`
}

func (o *ArchXml) Init() {
	o.initMainEnv()
	o.initTasks()
}

func (o *ArchXml) initMainEnv() {
	runningRoot := osxu.GetRunningDir()
	if o.Env == "" {
		o.Env = runningRoot
		return
	}
	if filex.IsDir(o.Env) {
		o.Env = filex.FormatPath(o.Env)
		return
	}
	o.Env = filex.Combine(runningRoot, o.Env)
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
		return o.Env
	}
	if filex.IsDir(task.Env) {
		return filex.FormatPath(task.Env)
	}
	if filex.IsRelativeFormat(task.Env) {
		return filex.Combine(o.Env, task.Env)
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
