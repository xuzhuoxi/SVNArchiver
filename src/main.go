package main

import (
	"github.com/xuzhuoxi/SVNArchiver/src/core"
	"github.com/xuzhuoxi/SVNArchiver/src/env"
)

func main() {
	core.InitLogger()
	cmdFlags, err := env.ParseFlags()
	if nil != err {
		core.Logger.Errorln("ParseFlogs Error[%s]!", err)
		return
	}

	core.ClearTempDir()

	core.TryHandleQueryLog(cmdFlags)
	core.TryHandleArchContext(cmdFlags)
	core.TryHandleArchXml(cmdFlags)

	core.ClearTempDir()
}
