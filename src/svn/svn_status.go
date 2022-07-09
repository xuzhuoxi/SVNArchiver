// Create on 2022/7/8
// @author xuzhuoxi
package svn

import (
	"os/exec"
	"encoding/xml"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
)


// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.status.html
// 功能：针对每个文件或目录，查询最新状态
// path可以为本地副本路径， 也可以是URL
// path使用URL时支持支持多个路径
// 版本号是整个svn仓库唯一共享的，所以这里返回的会出现断层情况
func QueryStatus(path string) (l *model.StatusResult, err error) {
	cmd := exec.Command(MainCmd, SubCmdStatus, ArgVerbose, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &model.StatusResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	rs.HandleResult()
	return rs, nil
}
