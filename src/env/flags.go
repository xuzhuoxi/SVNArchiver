package env

import (
	"errors"
	"flag"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strings"
)

type CmdFlags struct {
	EnvPath    string // 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	LogSize    int
	TargetPath string
	ArchPath   string

	Reversion int // 版本导出
	RevDiffN  int // 版本差异前
	RevDiffM  int // 版本差异后

	Date      string // 版本时间导出
	DateDiffN string // 版本时间差异前
	DateDiffM string // 版本时间差异后

	defaultEvn bool
}

func (f *CmdFlags) GetLogContext() (ctx *LogContext) {
	if !f.isLogCommand() {
		return nil
	}
	return &LogContext{TargetPath: f.TargetPath, LogSize: f.LogSize}
}

func (f *CmdFlags) GetRevDiffArchContext() (ctx *ArchRevDiffContext, err error) {
	if !f.isRevDiffArchCommand() {
		return nil, nil
	}
	return &ArchRevDiffContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath,
		RevStart: f.RevDiffN, RevTarget: f.RevDiffM}, nil
}

func (f *CmdFlags) GetRevArchContext() (ctx *ArchRevContext, err error) {
	if !f.isRevArchCommand() {
		return nil, nil
	}
	return &ArchRevContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath,
		Reversion: f.Reversion}, nil
}

func (f *CmdFlags) GetDateDiffArchContext() (ctx *ArchDateDiffContext, err error) {
	if !f.isDateDiffArchCommand() {
		return nil, nil
	}
	ctx = &ArchDateDiffContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath}
	if f.DateDiffN != "" {
		start, e := ParseInputDatetime(f.DateDiffN)
		if nil != e {
			return nil, e
		}
		ctx.DateStart, ctx.DateStartStr, ctx.ExistStart = start, f.DateDiffN, true
	}
	if f.DateDiffM != "" {
		target, e := ParseInputDatetime(f.DateDiffM)
		if nil != e {
			return nil, e
		}
		ctx.DateTarget, ctx.DateTargetStr, ctx.ExistTarget = target, f.DateDiffM, true
	}
	return
}

func (f *CmdFlags) GetDateArchContext() (ctx *ArchDateContext, err error) {
	if !f.isDateArchCommand() {
		return nil, nil
	}
	ctx = &ArchDateContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath, DateStr: f.Date}
	date, e := ParseInputDatetime(f.Date)
	if nil != e {
		return nil, e
	}
	ctx.Date = date
	return
}

func (f *CmdFlags) init() error {
	f.EnvPath, f.defaultEvn = f.getEnvPath()
	targetPath, exist := f.getTargetPath()
	if !exist {
		return errors.New("Target Path is not exist! ")
	}
	if filex.IsAbsFormat(targetPath) {
		f.TargetPath = targetPath
	} else {
		f.TargetPath = filex.Combine(f.TargetPath, targetPath)
	}
	if !filex.IsAbsFormat(f.ArchPath) {
		f.ArchPath = filex.Combine(f.EnvPath, f.ArchPath)
	}
	return nil
}

func (f *CmdFlags) isLogCommand() bool {
	return f.ArchPath == "" && !f.isArchCommand()
}

func (f *CmdFlags) isArchCommand() bool {
	return f.isRevArchCommand() || f.isRevDiffArchCommand() ||
		f.isDateArchCommand() || f.isDateDiffArchCommand()
}

func (f *CmdFlags) isRevArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.hasRevParams()
}

func (f *CmdFlags) isRevDiffArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.hasRevDiffParams()
}

func (f *CmdFlags) hasRevParams() bool {
	return f.Reversion > 0
}

func (f *CmdFlags) hasRevDiffParams() bool {
	return f.RevDiffN > 0 || f.RevDiffM > 0
}

func (f *CmdFlags) isDateArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.hasDateParams()
}

func (f *CmdFlags) isDateDiffArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.hasDateDiffParams()
}

func (f *CmdFlags) hasDateParams() bool {
	return f.Date != ""
}

func (f *CmdFlags) hasDateDiffParams() bool {
	return f.DateDiffN != "" || f.DateDiffM != ""
}

func (f *CmdFlags) getEnvPath() (evnPath string, isDefault bool) {
	runningRoot := osxu.GetRunningDir()
	if "" == f.EnvPath {
		return runningRoot, true
	}
	if filex.IsDir(f.EnvPath) {
		return f.EnvPath, false
	}
	return filex.Combine(runningRoot, f.EnvPath), false
}

func (f *CmdFlags) getTargetPath() (targetPath string, exist bool) {
	if filex.IsExist(f.TargetPath) {
		return f.TargetPath, true
	}
	rs := filex.Combine(f.EnvPath, f.TargetPath)
	if filex.IsExist(rs) {
		return rs, true
	}
	return f.TargetPath, false
}

func ParseFlags() (flags *CmdFlags, err error) {
	// 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	envPath := flag.String("env", "", "Running Environment Path! ")
	logSize := flag.Int("log", 0, "Max Log entry size! ")
	target := flag.String("target", "", "Target Path! ")

	arch := flag.String("arch", "", "Arch File Path! ")

	rev := flag.Int("r", 0, "Reversion Number! ")
	revN := flag.Int("r0", 0, "Start Reversion Number! ")
	revM := flag.Int("r1", 0, "Target Reversion Number! ")
	date := flag.String("d", "", "Start Date! ")
	dateN := flag.String("d0", "", "Start Date! ")
	dateM := flag.String("d1", "", "Target Date! ")

	flag.Parse()
	rs := &CmdFlags{
		EnvPath: strings.TrimSpace(*envPath), LogSize: *logSize,
		TargetPath: strings.TrimSpace(*target), ArchPath: strings.TrimSpace(*arch),
		Reversion: *rev, RevDiffN: *revN, RevDiffM: *revM,
		Date: strings.TrimSpace(*date), DateDiffN: strings.TrimSpace(*dateN), DateDiffM: strings.TrimSpace(*dateM)}
	err = rs.init()
	if nil != err {
		return
	}
	return rs, nil
}
