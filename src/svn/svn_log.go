package svn

import (
	"encoding/xml"
	"fmt"
	"os/exec"

	"github.com/xuzhuoxi/SVNArchiver/src/model"
)

// QueryLog 远程查询全部的提交记录。
// 因此日志范围不等于 svn log 默认的全程（Tortoise Show log）。
// path 须为工作副本路径（QueryLocalVersion 不接受 URL）。
// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
func QueryLog(path string) (l *model.LogResult, err error) {
	cmd := exec.Command(MainCmd, SubCmdLog, ArgVerbose, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	return parseLogResult(out)
}

// QueryLogRange 远程查询指定版本区间的提交记录。
// 实际命令: svn log -v --revision rMin:rMax --xml <path>
// path 可以为本地工作副本路径，也可以是仓库 URL。
// 版本号全仓库共享，某条路径上的提交不一定连续。
// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
func QueryLogRange(path string, rMin int, rMax int) (l *model.LogResult, err error) {
	verInfo := fmt.Sprintf("%d:%d", rMin, rMax)
	cmd := exec.Command(MainCmd, SubCmdLog, ArgVerbose, ArgRevision, verInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	return parseLogResult(out)
}

func parseLogResult(result []byte) (*model.LogResult, error) {
	rs := &model.LogResult{}
	err := xml.Unmarshal(result, rs)
	if nil != err {
		return nil, err
	}
	rs.HandleLogs()
	return rs, nil
}
