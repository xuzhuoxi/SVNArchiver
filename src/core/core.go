// Create on 2022/7/10
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/osxu"
	"os"
	"strconv"
)

var (
	runningTemp = filex.Combine(osxu.GetRunningDir(), ".temp_export")
	tempIndex   = 1
)

func ClearTempDir() {
	filex.RemoveAll(runningTemp)
}

func genNextTempDir() string {
	for {
		temp := getNextTempDir()
		if !filex.IsExist(temp) {
			os.MkdirAll(temp, os.ModePerm)
			return temp
		}
	}
	return ""
}

func getNextTempDir() string {
	temp := filex.Combine(runningTemp, strconv.Itoa(tempIndex))
	tempIndex += 1
	return temp
}
