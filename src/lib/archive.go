// Create on 2022/7/10
// @author xuzhuoxi
package lib

import (
	"archive/tar"
	"bufio"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"io"
	"io/fs"
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
	tarWriter := tar.NewWriter(f)
	defer tarWriter.Close()
	//gzipWriter := gzip.NewWriter(tarWriter)
	//defer gzipWriter.Close()

	if filex.IsDir(filePath) {
		filePathLen := len(filePath)
		filex.WaldAllFiles(filePath, func(path string, info fs.FileInfo, err error) error {
			if nil != err {
				return err
			}
			if info.IsDir() {
				return nil
			}
			tarHeardName := path[filePathLen+1:]
			logErr := writeFileToTar(path, tarHeardName, tarWriter)
			if nil != logErr {
				fmt.Println(logErr)
			}
			return nil
		})
	} else {
		_, filename := filex.Split(filePath)
		logErr := writeFileToTar(filePath, filename, tarWriter)
		if nil != logErr {
			fmt.Println(logErr)
		}
	}
	return nil

}

func writeFileToTar(filePath string, tarHeaderName string, tarWriter *tar.Writer) error {
	fileInfo, err := os.Stat(filePath)
	if nil != err {
		return err
	}
	srcFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		return err
	}
	//hdr.Name = UTF8ToHZGB2312(tarHeaderName)
	hdr.Name = tarHeaderName
	hdr.Format = tar.FormatGNU
	tarWriter.WriteHeader(hdr)
	_, err = io.Copy(tarWriter, reader)
	if err != nil {
		return err
	}
	return nil
}
