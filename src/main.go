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

	if ctx := cmdFlags.GetLogContext(); nil != ctx {
		core.HandleLog(ctx)
	}

	ctx, err := cmdFlags.GetArchContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	if ctx != nil {
		core.HandleArch(ctx)
	}
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
