package main

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/core"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
	"github.com/xuzhuoxi/SVNArchiver/src/svnversion"
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

	//-----------------------
	ctx0, err := cmdFlags.GetRevArchContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	core.HandleRevArch(ctx0)

	//-----------------------
	ctx1, err := cmdFlags.GetRevDiffArchContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	core.HandleRevDiffArch(ctx1)

	//-----------------------
	ctx2, err := cmdFlags.GetDateArchContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	core.HandleDateArch(ctx2)

	//-----------------------
	ctx3, err := cmdFlags.GetDateDiffArchContext()
	if nil != err {
		fmt.Println(err)
		return
	}
	core.HandleDateDiffArch(ctx3)
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
