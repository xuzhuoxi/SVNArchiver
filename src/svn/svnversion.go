// Package svn
// Create on 2022/7/7
// @author xuzhuoxi
package svn

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CommandName = "svnversion"
	sep         = ":"
)

const (
	ArgNoNewline = "-n" // 不要打印换行符.
	ArgCommitted = "-c" // 使用各节点最后一次提交修改的版本号 (changed_revision), 而不是检出版本号 (revision).
)

const (
	MarkModify = "M" // 有本地修改
	MarkSwitch = "S" // 有 switch
	MarkSparse = "P" // 稀疏检出
)

type VersionResult struct {
	Min, Max               int  // 单一版本时两者相同; 混合版本或 -c 时为区间两端
	Modify, Switch, Sparse bool // 对应输出后缀 M / S / P
}

func (r VersionResult) String() string {
	return fmt.Sprintf("{Min:%d, Max:%d}", r.Min, r.Max)
}

// QueryLocalVersion 读取本地工作副本元数据，不连接仓库。
// 实际命令: svnversion -n -c <path>
// path 只能是工作副本路径，不能是仓库 URL。
// 输出为单一版本号 N，或混合区间 min:max，末尾可带 M/S/P。
// 使用 -c 时 min:max 是当前树上各节点 last-changed 的最小/最大，
// 不等于 svn log 的完整历史（中间版本可能不连续，已被后续修改覆盖的提交也不会出现）。
// 非工作副本时 svnversion 会输出英文说明（如 Unversioned directory），此时解析失败。
// https://svnbook.red-bean.com/zh/1.8/svn.ref.svnversion.re.html
func QueryLocalVersion(path string) (r *VersionResult, err error) {
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
