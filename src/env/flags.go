package env

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"errors"
	"strings"
	"strconv"
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/lib"
	"time"
)

type CmdFlags struct {
	EnvPath    string // 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	LogSize    int
	TargetPath string
	ArchPath   string
	VerRange   string
	DateRange  string

	defaultEvn bool
}

func (f *CmdFlags) GetLogContext() (ctx *LogContext) {
	if !f.isLogCommand() {
		return nil
	}
	return &LogContext{TargetPath: f.TargetPath, LogSize: f.LogSize}
}

func (f *CmdFlags) GetArchContext() (ctx *ArchContext, err error) {
	if !f.isVerArchCommand() {
		return nil, nil
	}
	s, t, e := f.parseVer()
	if nil != e {
		return nil, e
	}
	return &ArchContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath, StartVer: s, TargetVer: t}, nil
}

func (f *CmdFlags) GetDateArchContext() (ctx *ArchContext, err error) {
	if !f.isDateArchCommand() {
		return nil, nil
	}
	s, t, e := f.parseDate()
	if nil != e {
		return nil, e
	}
	dCtx := &ArchDateContext{TargetPath: f.TargetPath, ArchPath: f.ArchPath, StartDate: s, TargetDate: t}
	return dCtx.GetArchContext(), nil
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
	return nil
}

func (f *CmdFlags) isLogCommand() bool {
	return f.ArchPath == ""
}

func (f *CmdFlags) isArchCommand() bool {
	return f.isVerArchCommand() || f.isDateArchCommand()
}

func (f *CmdFlags) isVerArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.VerRange != ""
}

func (f *CmdFlags) isDateArchCommand() bool {
	return f.LogSize == 0 && f.ArchPath != "" && f.DateRange != ""
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

func (f *CmdFlags) parseVer() (start int, target int, err error) {
	if "" == f.VerRange {
		err = errors.New(fmt.Sprintf("VerRange format error: [%s]", f.VerRange))
		return
	}
	if strings.Contains(f.VerRange, SepVer) {
		arr := strings.Split(f.VerRange, SepVer)
		if len(arr) == 2 {
			start, err = strconv.Atoi(arr[0])
			if nil != err {
				return
			}
			target, err = strconv.Atoi(arr[1])
			if nil != err {
				return
			}
			return
		}
		err = errors.New(fmt.Sprintf("VerRange format error: [%s]", f.VerRange))
	}
	target, err = strconv.Atoi(f.VerRange)
	if nil != err {
		return
	}
	return
}

func (f *CmdFlags) parseDate() (start time.Time, target time.Time, err error) {
	if "" == f.DateRange {
		err = errors.New(fmt.Sprintf("DateRange format error: [%s]", f.VerRange))
		return
	}
	if strings.Contains(f.VerRange, SepVer) {
		arr := strings.Split(f.VerRange, SepVer)
		if len(arr) == 2 {
			start, err = lib.ParseDatetime(arr[0])
			if nil != err {
				return
			}
			target, err = lib.ParseDatetime(arr[1])
			if nil != err {
				return
			}
			return
		}
		err = errors.New(fmt.Sprintf("VerRange format error: [%s]", f.VerRange))
	}
	target, err = lib.ParseDatetime(f.VerRange)
	if nil != err {
		return
	}
	start = lib.DatetimeZero
	return
}

func ParseFlags() (flags *CmdFlags, err error) {
	// 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	envPath := flag.String("env", "", "Running Environment Path! ")
	logSize := flag.Int("log", 0, "Version Info Count! ")
	target := flag.String("target", "", "Target Path! ")
	arch := flag.String("arch", "", "Arch File Path! ")
	v := flag.String("v", "", "Version Setting! ")
	d := flag.String("d", "", "Date Setting! ")

	flag.Parse()
	rs := &CmdFlags{
		EnvPath:    strings.TrimSpace(*envPath), LogSize: *logSize,
		TargetPath: strings.TrimSpace(*target), ArchPath: strings.TrimSpace(*arch),
		VerRange:   strings.TrimSpace(*v), DateRange: strings.TrimSpace(*d)}
	err = rs.init()
	if nil != err {
		return
	}
	return rs, nil
}
