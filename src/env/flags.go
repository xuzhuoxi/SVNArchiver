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
	XmlPath    string // 批量任务的配置文件路径。
	LogSize    int    // 查询条目显示的最大数量，要求>=0, 当值为0时，认定为无限制。
	TargetPath string // 归档处理的svn目录，可以是svn仓库的非根目录。
	ArchPath   string // 归档文件保存路径，支持通配符。
	Reversion  int    // 完整归档时使用, 用于指定具体版本号，并使用该版本号(或向前最近的版本号)进行归档。
	RevDiffN   int    // 差异归档时使用, 用于指定起始版本号。
	RevDiffM   int    // 差异归档时使用, 用于指定结束版本号。
	Date       string // 完整归档时使用, 用于指定一个时间点，并使用该时间点上的版本号(或向前最近的版本号)进行归档。
	DateDiffN  string // 差异归档时使用, 用于指定起始时间。
	DateDiffM  string // 差异归档时使用, 用于指定结束时间。
}

func (f *CmdFlags) GetArchXml() (ctx *ArchXml, err error) {
	if !f.isArchXmlCommand() {
		return nil, nil
	}
	os.ReadFile(f.XmlPath)
	bs, err := os.ReadFile(f.XmlPath)
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
		taskPath, exist := f.getFixPath(f.XmlPath)
		if !exist {
			return errors.New(fmt.Sprintf("Task Path[%s] is not exist! ", f.XmlPath))
		}
		f.XmlPath = taskPath
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
	return f.XmlPath != "" && f.TargetPath == ""
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
	envPath := flag.String("env", "", "Running Environment Path! ") // 【可选】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	xml := flag.String("xml", "", "Task Config XML!")               // 批量任务的配置文件路径。
	size := flag.Int("size", 0, "Max Log entry size! ")             // 查询条目显示的最大数量，要求>=0, 当值为0时，认定为无限制。

	target := flag.String("target", "", "Target Path! ")   // 归档处理的svn目录，可以是svn仓库的非根目录。
	arch := flag.String("arch", "", "Arch File Path! ")    // 归档文件保存路径，支持通配符。
	rev := flag.Int("r", 0, "Reversion Number! ")          // 完整归档时使用, 用于指定具体版本号，并使用该版本号(或向前最近的版本号)进行归档。
	revN := flag.Int("r0", 0, "Start Reversion Number! ")  // 差异归档时使用, 用于指定起始版本号。
	revM := flag.Int("r1", 0, "Target Reversion Number! ") // 差异归档时使用, 用于指定结束版本号。
	date := flag.String("d", "", "Start Date! ")           // 完整归档时使用, 用于指定一个时间点，并使用该时间点上的版本号(或向前最近的版本号)进行归档。
	dateN := flag.String("d0", "", "Start Date! ")         // 差异归档时使用, 用于指定起始时间。
	dateM := flag.String("d1", "", "Target Date! ")        // 差异归档时使用, 用于指定结束时间。

	flag.Parse()
	rs := &CmdFlags{
		EnvPath: strings.TrimSpace(*envPath), LogSize: *size, XmlPath: strings.TrimSpace(*xml),
		TargetPath: strings.TrimSpace(*target), ArchPath: strings.TrimSpace(*arch),
		Reversion: *rev, RevDiffN: *revN, RevDiffM: *revM,
		Date: strings.TrimSpace(*date), DateDiffN: strings.TrimSpace(*dateN), DateDiffM: strings.TrimSpace(*dateM)}
	err = rs.init()
	if nil != err {
		return
	}
	return rs, nil
}
