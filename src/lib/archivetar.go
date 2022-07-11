// Create on 2022/7/11
// @author xuzhuoxi
package lib

import (
	"archive/tar"
	"bufio"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"io"
	"io/fs"
	"os"
)

func ArchiveTar(targetWriter io.Writer, srcPath string) {
	tarWriter := tar.NewWriter(targetWriter)
	defer tarWriter.Close()
	//gzipWriter := gzip.NewWriter(tarWriter)
	//defer gzipWriter.Close()

	if filex.IsDir(srcPath) {
		writeDirToTar(srcPath, tarWriter)
	} else {
		_, filename := filex.Split(srcPath)
		logErr := writeFileToTar(srcPath, filename, tarWriter)
		if nil != logErr {
			fmt.Println(logErr)
		}
	}
}

func writeDirToTar(filePath string, tarWriter *tar.Writer) error {
	filePathLen := len(filePath)
	filex.WalkAllFiles(filePath, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			return err
		}
		tarHeardName := path[filePathLen+1:]
		logErr := writeFileToTar(path, tarHeardName, tarWriter)
		if nil != logErr {
			fmt.Println(logErr)
		}
		return nil
	})
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
