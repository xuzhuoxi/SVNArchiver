package svn

import (
	"encoding/xml"
	"os/exec"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.list.html
func QueryList(path string, recursive bool) (l *model.ListResult, err error) {
	var cmd *exec.Cmd
	if recursive {
		cmd = exec.Command(MainCmd, SubCmdList, ArgXml, ArgRecursive, path)
	} else {
		cmd = exec.Command(MainCmd, SubCmdList, ArgXml, path)
	}
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &model.ListResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	return rs, nil
}
