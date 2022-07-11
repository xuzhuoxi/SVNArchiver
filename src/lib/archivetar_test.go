// Create on 2022/7/11
// @author xuzhuoxi
package lib

//
//var (
//	srcPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\SVNArchiver\res\diff2.xml`
//	tarPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\SVNArchiver\res\diff2.tar`
//)
//
//func TestPaxNonAscii(t *testing.T) {
//	// Create an archive with non ascii. These should trigger a pax header
//	// because pax headers have a defined utf-8 encoding.
//	fileinfo, err := os.Stat(`D:\workspaces\GoPath\src\github.com\xuzhuoxi\SVNArchiver\res\diff2.xml`)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	hdr, err := tar.FileInfoHeader(fileinfo, "")
//	if err != nil {
//		t.Fatalf("os.Stat:1 %v", err)
//	}
//
//	// some sample data
//	chineseFilename := "文件名.xml"
//	chineseGroupname := "組"
//	chineseUsername := "用戶名"
//
//	hdr.Name = chineseFilename
//	hdr.Gname = chineseGroupname
//	hdr.Uname = chineseUsername
//	hdr.Format = tar.FormatPAX
//
//	srcFile, _ := os.Open(srcPath)
//	defer srcFile.Close()
//	tarFile, _ := os.Create(tarPath)
//	defer tarFile.Close()
//	writer := tar.NewWriter(tarFile)
//	defer writer.Close()
//
//	writer.WriteHeader(hdr)
//	io.Copy(writer, srcFile)
//}
