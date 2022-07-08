package main

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/core"
)

func main() {
	cmdFlags, err := env.ParseFlags()
	if nil != err {
		fmt.Println(err)
		return
	}
	vCtx, lCtx, archCtx, err := cmdFlags.GetContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	core.HandleVersion(vCtx)
	core.HandleStatus(lCtx)
	core.HandleArch(archCtx)
}

func demo() {
	path := `H:/SvnTest`

	//log, err := svn.QueryLog(path)
	log, err := svnversion.QueryVersion(path)

	if nil != err {
		fmt.Println("错误：", err)
		return
	}
	fmt.Println("成功：", log)
}
