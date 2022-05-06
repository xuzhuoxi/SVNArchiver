package main

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/svn"
)

func main() {
	path := `D:/workspaces/Project2204/Project2204_Common/Configs/source`

	log, err := svn.QueryLog(path)

	if nil != err {
		fmt.Println("错误：", err)
		return
	}
	fmt.Println("成功：", log.Name.Local, log.Name.Space, len(log.LogEntries))
}
