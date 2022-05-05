package command

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"runtime"
)

var (
	isWin          = isWinOS()
	gb18030Decoder = simplifiedchinese.GB18030.NewDecoder()
)

func Bytes2String(cmdResult []byte) string {
	if isWin {
		bs, err := gb18030Decoder.Bytes(cmdResult)
		if nil != err {
			return ""
		}
		return string(bs)
	}
	return string(cmdResult)
}

func isWinOS() bool {
	return "windows" == runtime.GOOS
}
