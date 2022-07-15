// Create on 2022/7/14
// @author xuzhuoxi
package env

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strings"
)

type ArchTask struct {
	TaskId       string // 任务标识
	TargetPath   string // 归档处理的svn目录，可以是svn仓库的非根目录。
	ArchPath     string // 归档文件保存路径，支持通配符。
	ArchOverride bool   // 归档文件存在时是否覆盖。

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
	return &ArchRevDiffContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath, Override: t.ArchOverride,
		RevStart: t.RevDiffN, RevTarget: t.RevDiffM}, nil
}

func (t *ArchTask) GetRevArchContext() (ctx *ArchRevContext, err error) {
	if !t.isRevArchCommand() {
		return nil, nil
	}
	return &ArchRevContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath, Override: t.ArchOverride,
		Reversion: t.Reversion}, nil
}

func (t *ArchTask) GetDateDiffArchContext() (ctx *ArchDateDiffContext, err error) {
	if !t.isDateDiffArchCommand() {
		return nil, nil
	}
	ctx = &ArchDateDiffContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath, Override: t.ArchOverride}
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
	ctx = &ArchDateContext{TargetPath: t.TargetPath, ArchPath: t.ArchPath, Override: t.ArchOverride, DateStr: t.Date}
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

type ArchXmlArchNode struct {
	Override string `xml:"override,attr"`
	XmlValue string `xml:",innerxml"`
}

func (o *ArchXmlArchNode) IsOverride() bool {
	return o.Override == "1" || strings.ToLower(o.Override) == "true"
}

func (o *ArchXmlArchNode) IsUnknown() bool {
	return o.Override == ""
}

type ArchXmlTask struct {
	Id     string           `xml:"id,attr"` // 归档任务Id
	Env    string           `xml:"env"`     // 归档任务的环境路径
	Target string           `xml:"target"`  // 归档任务的目标svn目录，支持绝对路径与相对路径(相对于env环境路径的实际运行值)。
	Arch   *ArchXmlArchNode `xml:"arch"`    // 归档文件的输出路径，支持绝对路径与相对路径(相对于env环境路径的实际运行值)。

	R  int    `xml:"r,attr"`  // 完整归档时使用, 用于指定具体版本号，并使用该版本号(或向前最近的版本号)进行归档。
	R0 int    `xml:"r0,attr"` // 差异归档时使用, 用于指定起始版本号。
	R1 int    `xml:"r1,attr"` // 差异归档时使用, 用于指定结束版本号。
	D  string `xml:"d,attr"`  // 完整归档时使用, 用于指定一个时间点，并使用该时间点上的版本号(或向前最近的版本号)进行归档。
	D0 string `xml:"d0,attr"` // 差异归档时使用, 用于指定起始时间。
	D1 string `xml:"d1,attr"` // 差异归档时使用, 用于指定结束时间。
}

type ArchXmlTasks struct {
	ArchOverride bool           `xml:"arch-override,attr"` // 归档文件存在时，是否覆盖
	Tasks        []*ArchXmlTask `xml:"task"`               // 归档任务列表
}

type ArchXmlLog struct {
	FileType string `xml:"file,attr"` // 归档信息文件的格式，支持json和xml
	CodeType string `xml:"code,attr"` // 归档文件提取的特征码类型，支持md5和sha1
	XmlValue string `xml:",innerxml"` // 归档信息文件的保存路径，支持绝对路径与相对路径(相对于main-env运行值)
}

type ArchXml struct {
	MainEnv string        `xml:"main-env"` // 主环境路径，可选，没有时使用执行文件所在目录
	Tasks   *ArchXmlTasks `xml:"tasks"`    // 归档任务配置， 必要
	Log     *ArchXmlLog   `xml:"log"`      // 归档信息记录配置，可选
}

func (o *ArchXml) Init() {
	o.initMainEnv()
	o.initTasks()
	o.initLog()
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
		if filex.IsExist(task.Arch.XmlValue) || filex.IsAbsFormat(task.Arch.XmlValue) {
			xmlTasks[index].Arch.XmlValue = filex.FormatPath(task.Arch.XmlValue)
		} else {
			xmlTasks[index].Arch.XmlValue = filex.Combine(xmlTasks[index].Env, task.Arch.XmlValue)
		}
	}
}

func (o *ArchXml) initLog() {
	if nil == o.Log {
		return
	}
	if filex.IsAbsFormat(o.Log.XmlValue) {
		return
	}
	o.Log.XmlValue = filex.Combine(o.MainEnv, o.Log.XmlValue)
	o.Log.FileType = strings.ToLower(o.Log.FileType)
	o.Log.CodeType = strings.ToLower(o.Log.CodeType)
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
		override := o.Tasks.ArchOverride
		if !xmlTask.Arch.IsUnknown() {
			override = xmlTask.Arch.IsOverride()
		}
		task := ArchTask{TaskId: xmlTask.Id,
			TargetPath: xmlTask.Target, ArchPath: xmlTask.Arch.XmlValue, ArchOverride: override,
			Reversion: xmlTask.R, RevDiffN: xmlTask.R0, RevDiffM: xmlTask.R1,
			Date: xmlTask.D, DateDiffN: xmlTask.D0, DateDiffM: xmlTask.D1}
		rs[index] = task
	}
	return rs
}

func (o *ArchXml) LogEnabled() bool {
	return o.Log != nil
}
