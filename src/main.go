package main

import (
	"fmt"
	"github.com/xuzhuoxi/SVNArchiver/src/command"
	"os/exec"
)

func main() {
	//cmd := exec.Command("ipconfig")
	cmd := exec.Command("svn", "help")
	out, err := cmd.CombinedOutput()
	if nil != err {
		fmt.Println("错误：", err)
		return
	}

	fmt.Println("返回：", command.Bytes2String(out))

}
