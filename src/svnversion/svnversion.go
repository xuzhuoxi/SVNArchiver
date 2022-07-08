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
	ArgNoNewline = "-n"
	ArgCommitted = "-c"
)

type VersionResult struct {
	RevisionMin, RevisionMax int
}

func (r VersionResult) String() string {
	return fmt.Sprintf("{Min:%d, Max:%d}", r.RevisionMin, r.RevisionMax)
}

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryVersion(path string) (l *VersionResult, err error) {
	cmd := exec.Command(CommandName, ArgNoNewline, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	if "" == outStr {
		return nil, nil
	}

	if !strings.Contains(outStr, sep) {
		version, err := strconv.ParseInt(outStr, 10, 32)
		if nil != err {
			return nil, err
		}
		return &VersionResult{RevisionMin: int(version), RevisionMax: int(version)}, nil
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
	return &VersionResult{RevisionMin: int(min), RevisionMax: int(max)}, nil
}
