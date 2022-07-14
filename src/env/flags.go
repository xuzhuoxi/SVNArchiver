package env

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"os"
	"strings"
)

type CmdFlags struct {
	EnvPath    string // 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	LogSize    int
	TaskPath   string
	TargetPath string
	ArchPath   string

	Reversion int // 版本导出
	RevDiffN  int // 版本差异前
	RevDiffM  int // 版本差异后

	Date      string // 版本时间导出
	DateDiffN string // 版本时间差异前
	DateDiffM string // 版本时间差异后
}

func (f *CmdFlags) GetArchXml() (ctx *ArchXml, err error) {
	if !f.isArchXmlCommand() {
		return nil, nil
	}
	os.ReadFile(f.TaskPath)
	bs, err := os.ReadFile(f.TaskPath)
	if nil != err {
		return nil, errors.New(fmt.Sprintf("Load Arch XML Fail[%s].", err))
	}
	ctx = &ArchXml{}
	err = xml.Unmarshal(bs, ctx)
	if nil != err {
		return nil, err
	}
	ctx.Init()
	return ctx, nil
}

func (f *CmdFlags) GetLogContext() (ctx *LogContext) {
	if !f.isLogCommand() {
		return nil
	}
	return &LogContext{TargetPath: f.TargetPath, LogSize: f.LogSize}
}

func (f *CmdFlags) GetArchTask() (ctx ArchTask, exist bool) {
	if !f.isArchTaskCommand() {
		return ArchTask{}, false
	}
	return ArchTask{TargetPath: f.TargetPath, ArchPath: f.ArchPath,
		Reversion: f.Reversion, RevDiffN: f.RevDiffN, RevDiffM: f.RevDiffM,
		Date: f.Date, DateDiffN: f.DateDiffN, DateDiffM: f.DateDiffM}, true
}

func (f *CmdFlags) init() error {
	f.EnvPath = f.getEnvPath()
	if f.isArchXmlCommand() {
		taskPath, exist := f.getFixPath(f.TaskPath)
		if !exist {
			return errors.New(fmt.Sprintf("Task Path[%s] is not exist! ", f.TaskPath))
		}
		f.TaskPath = taskPath
		return nil
	}

	targetPath, exist := f.getFixPath(f.TargetPath)
	if !exist {
		return errors.New(fmt.Sprintf("Target Path[%s] is not exist! ", f.TargetPath))
	}
	f.TargetPath = targetPath

	f.ArchPath, _ = f.getFixPath(f.ArchPath)
	return nil
}

func (f *CmdFlags) isLogCommand() bool {
	return f.TargetPath != "" && f.ArchPath == ""
}

func (f *CmdFlags) isArchXmlCommand() bool {
	return f.TaskPath != "" && f.TargetPath == ""
}

func (f *CmdFlags) isArchTaskCommand() bool {
	return f.TargetPath != "" && f.ArchPath != ""
}

func (f *CmdFlags) getEnvPath() string {
	runningRoot := osxu.GetRunningDir()
	if "" == f.EnvPath {
		return runningRoot
	}
	if filex.IsDir(f.EnvPath) {
		return filex.FormatPath(f.EnvPath)
	}
	return filex.Combine(runningRoot, f.EnvPath)
}

func (f *CmdFlags) getFixPath(path string) (newPath string, exist bool) {
	//fmt.Println("getFixPath:", path, filex.IsExist(path), filex.IsAbsFormat(path))
	if filex.IsExist(path) {
		return path, true
	}
	if filex.IsAbsFormat(path) {
		return filex.FormatPath(path), true
	}
	rs := filex.Combine(f.EnvPath, path)
	if filex.IsExist(rs) {
		return rs, true
	}
	return path, false
}

func ParseFlags() (flags *CmdFlags, err error) {
	// 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	envPath := flag.String("env", "", "Running Environment Path! ")
	target := flag.String("target", "", "Target Path! ")

	logSize := flag.Int("log", 0, "Max Log entry size! ")
	task := flag.String("task", "", "Task Config XML!")

	arch := flag.String("arch", "", "Arch File Path! ")
	rev := flag.Int("r", 0, "Reversion Number! ")
	revN := flag.Int("r0", 0, "Start Reversion Number! ")
	revM := flag.Int("r1", 0, "Target Reversion Number! ")
	date := flag.String("d", "", "Start Date! ")
	dateN := flag.String("d0", "", "Start Date! ")
	dateM := flag.String("d1", "", "Target Date! ")

	flag.Parse()
	rs := &CmdFlags{
		EnvPath: strings.TrimSpace(*envPath), LogSize: *logSize, TaskPath: strings.TrimSpace(*task),
		TargetPath: strings.TrimSpace(*target), ArchPath: strings.TrimSpace(*arch),
		Reversion: *rev, RevDiffN: *revN, RevDiffM: *revM,
		Date: strings.TrimSpace(*date), DateDiffN: strings.TrimSpace(*dateN), DateDiffM: strings.TrimSpace(*dateM)}
	err = rs.init()
	if nil != err {
		return
	}
	return rs, nil
}
