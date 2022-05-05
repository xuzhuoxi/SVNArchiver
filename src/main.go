package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
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

	decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(out)
	fmt.Println("返回：", string(decodeBytes))
}
