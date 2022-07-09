// Create on 2022/7/7
// @author xuzhuoxi
package svnversion

import (
	"os/exec"
	"strings"
	"strconv"
	"fmt"
)

const (
	CommandName = "svnversion"
	sep         = ":"
)

const (
	ArgNoNewline = "-n" // 不要打印换行符.
	ArgCommitted = "-c" // 使用最近一次产生修改的版本号, 而不是当前版本号 (当前版本号是 本地可获得的, 值最大的版本号).
)

const (
	MarkModify = "M" // 有本地悠
	MarkSwitch = "S" // 切换过
	MarkSparse = "P" // 稀疏
)

type VersionResult struct {
	Min, Max               int
	Modify, Switch, Sparse bool
}

func (r VersionResult) String() string {
	return fmt.Sprintf("{Min:%d, Max:%d}", r.Min, r.Max)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryVersion(path string) (r *VersionResult, err error) {
	cmd := exec.Command(CommandName, ArgNoNewline, ArgCommitted, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	markModify := strings.Contains(outStr, MarkModify)
	markSwitch := strings.Contains(outStr, MarkSwitch)
	markSparse := strings.Contains(outStr, MarkSparse)
	outStr = clearMarks(outStr, MarkModify, MarkSparse, MarkSwitch)
	if "" == outStr {
		return nil, nil
	}

	if !strings.Contains(outStr, sep) {
		version, err := strconv.ParseInt(outStr, 10, 32)
		if nil != err {
			return nil, err
		}
		return &VersionResult{Min: int(version), Max: int(version),
			Modify: markModify, Switch: markSwitch, Sparse: markSparse}, nil
	}

	arr := strings.Split(outStr, sep)
	min, err := strconv.ParseInt(arr[0], 10, 32)
	if nil != err {
		return nil, err
	}
	max, err := strconv.ParseInt(arr[1], 10, 32)
	if nil != err {
		return nil, err
	}
	return &VersionResult{Min: int(min), Max: int(max),
		Modify: markModify, Switch: markSwitch, Sparse: markSparse}, nil
}

func clearMarks(rs string, marks ...string) string {
	if len(marks) == 0 {
		return rs
	}
	for _, m := range marks {
		rs = strings.ReplaceAll(rs, m, "")
	}
	return rs
}
