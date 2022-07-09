// Create on 2022/7/9
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func HandleLog(ctx *env.LogContext) {
	if nil == ctx {
		return
	}
	fmt.Println(`Handle "svn log" Command:`)
	rs, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		fmt.Println("QueryList Error:", err)
		return
	}
	es := rs.LogEntries
	size := ctx.LogSize
	for index := len(es) - 1; index >= 0; index -= 1 {
		printLogEntry(es[index])
		size -= 1
		if size == 0 {
			break
		}
	}
}

func printLogEntry(e *model.LogResultEntry) {
	reversion := e.Revision
	actions := []byte("   ")
	as := []byte(e.GetActions())
	copy(actions, as)
	author := e.Author
	date := e.Date[:22]
	msg := e.Msg
	fmt.Println(fmt.Sprintf("%d \t %s \t %s \t %s \t %s", reversion, string(actions), author, date, msg))
}
