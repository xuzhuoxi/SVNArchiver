package lib

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"runtime"
)

var (
	isWin           = isWinOS()
	gb18030Decoder  = simplifiedchinese.GB18030.NewDecoder()
	gbkDecoder      = simplifiedchinese.GBK.NewDecoder()
	hzgb2313Decoder = simplifiedchinese.HZGB2312.NewDecoder()
)

func UTF8ToGB18030(utf8 string) string {
	return utf8To(utf8, gb18030Decoder)
}

func UTF8ToGBK(utf8 string) string {
	return utf8To(utf8, gbkDecoder)
}

func UTF8ToHZGB2312(utf8 string) string {
	return utf8To(utf8, hzgb2313Decoder)
}

func utf8To(utf8 string, decoder *encoding.Decoder) string {
	bs, err := decoder.Bytes([]byte(utf8))
	if nil != err {
		return ""
	}
	return string(bs)
}

func isWinOS() bool {
	return "windows" == runtime.GOOS
}
