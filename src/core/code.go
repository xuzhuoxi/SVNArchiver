// Create on 2022/7/15
// @author xuzhuoxi
package core

import (
	"crypto"
	"github.com/xuzhuoxi/infra-go/cryptox"
)

var codeTypeMap map[string]crypto.Hash

func init() {
	codeTypeMap = make(map[string]crypto.Hash)
	codeTypeMap["md5"] = crypto.MD5
	codeTypeMap["sha1"] = crypto.SHA1
}

func GetCode(path string, codeType string) string {
	if path == "" || codeType == "" {
		return ""
	}
	if hash, ok := codeTypeMap[codeType]; ok {
		return cryptox.HashFile2Hex(hash, path)
	}
	return ""
}
