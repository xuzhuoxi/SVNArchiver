package svn

import (
	"encoding/xml"
	"fmt"
	"os/exec"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryLog(path string) (l *model.LogResult, err error) {
	rsVersion, err := svnversion.QueryVersion(path)
	if nil != err {
		return nil, err
	}
	verInfo := fmt.Sprintf("%d:%d", rsVersion.Min, rsVersion.Max)
	cmd := exec.Command(MainCmd, SubCmdLog, ArgVerbose, ArgRevision, verInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &model.LogResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	return rs, nil
}
