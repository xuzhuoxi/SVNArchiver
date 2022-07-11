// Create on 2022/7/10
// @author xuzhuoxi
package lib

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
)

func Archive(filePath string, archPath string, override bool) error {
	if !filex.IsExist(filePath) {
		return errors.New(fmt.Sprintf("FilePath[%s] is not exist!", filePath))
	}
	if filex.IsFile(archPath) {
		if !override {
			return errors.New(fmt.Sprintf("ArchPath[%s] is exist!", archPath))
		}
		filex.Remove(archPath)
	}
	dir, err := filex.GetUpDir(archPath)
	if nil != err {
		return errors.New(fmt.Sprintf("Parse archPath[%s] fail! ", archPath))
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if nil != err {
		return errors.New(fmt.Sprintf("Mkdir archPath[%s] Parent folder fail! ", archPath))
	}
	f, err := os.Create(archPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Create arch[%s] fail! ", archPath))
	}
	defer f.Close()

	ArchiveZip(f, filePath)
	return nil
}
