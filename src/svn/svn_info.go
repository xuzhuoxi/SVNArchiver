package svn

import (
	"encoding/xml"
	"os/exec"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.info.html
func QueryInfo(path string) (l *model.InfoResult, err error) {
	cmd := exec.Command(MainCmd, SubCmdInfo, ArgXml, path)
	out, err := cmd.CombinedOutput()
	if nil != err {
		return nil, err
	}
	rs := &model.InfoResult{}
	err = xml.Unmarshal(out, rs)
	if nil != err {
		return nil, err
	}
	return rs, nil
}
