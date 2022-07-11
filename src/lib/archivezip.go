// Create on 2022/7/11
// @author xuzhuoxi
package lib

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"io"
	"io/fs"
	"os"
)

func ArchiveZip(targetWriter io.Writer, srcPath string) {
	tarWriter := zip.NewWriter(targetWriter)
	defer tarWriter.Close()
	//gzipWriter := gzip.NewWriter(tarWriter)
	//defer gzipWriter.Close()

	if filex.IsDir(srcPath) {
		writeDirToZip(srcPath, tarWriter)
	} else {
		_, filename := filex.Split(srcPath)
		logErr := writeFileToZip(srcPath, filename, tarWriter)
		if nil != logErr {
			fmt.Println(logErr)
		}
	}
}

func writeDirToZip(filePath string, tarWriter *zip.Writer) error {
	filePathLen := len(filePath)
	filex.WalkAllFiles(filePath, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}
		tarHeardName := path[filePathLen+1:]
		logErr := writeFileToZip(path, tarHeardName, tarWriter)
		if nil != logErr {
			fmt.Println(logErr)
		}
		return nil
	})
	return nil
}

func writeFileToZip(filePath string, tarHeaderName string, zipWriter *zip.Writer) error {
	fileInfo, err := os.Stat(filePath)
	if nil != err {
		return err
	}

	hdr, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	hdr.Name = tarHeaderName
	subWriter, err := zipWriter.CreateHeader(hdr)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	_, err = io.Copy(subWriter, reader)
	if err != nil {
		return err
	}
	return nil
}
