// Create on 2022/7/8
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func HandleSvnStatus(ctx *env.QueryLogContext) {
	if nil == ctx {
		return
	}
	Logger.Println(`Handle "svn status" Command:`)
	rs, err := svn.QueryStatus(ctx.TargetPath)
	if nil != err {
		Logger.Warnln("QueryStatus Error:", err)
		return
	}
	for _, v := range rs.StatusTarget.EntryList {
		committed := v.WcStatus.Commit
		author := committed.Author
		if author == "" {
			author = "unknown"
		}
		date := committed.DateString()
		Logger.Println(fmt.Sprintf("%d \t %s \t %s \t %s", committed.Revision, date, author, v.Path))
	}
}
