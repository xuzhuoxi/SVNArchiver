// Create on 2022/7/9
// @author xuzhuoxi
package core

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/model"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func HandleSvnLog(ctx *env.QueryLogContext) {
	if nil == ctx {
		return
	}
	if ctx.LogSize < 0 {
		Logger.Warnln(fmt.Sprintf(`HandleSvnLog log value[%d] should >= 0. `, ctx.LogSize))
		return
	}
	Logger.Infoln(`HandleSvnLog with command["svn log"]:`)
	rs, err := svn.QueryLog(ctx.TargetPath)
	if nil != err {
		Logger.Warnln(fmt.Sprintf(`HandleSvnLog ["svnQueryLog"] Error[%s]`, err))
		return
	}
	es := rs.LogEntries
	size := ctx.LogSize
	logLimit := size > 0
	Logger.Infoln(fmt.Sprintf(`HandleSvnLog Result[MaxSize=%d, PrintSize=%d]:`, len(es), size))
	for index := len(es) - 1; index >= 0; index -= 1 {
		printLogEntry(es[index])
		if logLimit {
			size -= 1
			if size == 0 {
				break
			}
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
	printStr := fmt.Sprintf("\t %d \t %s \t %s \t %s \t %s", reversion, string(actions), author, date, msg)
	Logger.Println(printStr)
}
