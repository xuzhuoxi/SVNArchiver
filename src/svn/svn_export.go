package svn

import (
	"fmt"
	"os/exec"
)

// https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.export.html
func Export(path string, revision int, dist string) error {
	vStr := fmt.Sprintf("-r%d", revision)
	cmd := exec.Command(MainCmd, SubCmdExport, vStr, path, dist)
	_, err := cmd.CombinedOutput()
	return err
}
