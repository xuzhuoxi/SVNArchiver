package svn

const (
	MainCmd = "svn"
)

// 子命令。目标为仓库 URL、带历史版本 -r（非 BASE）、或 status -u 时都会连接远程仓库。
const (
	SubCmdInfo   = "info"   // WC 路径且不带 -r（或 -r BASE）时读本地；URL 或 -r(非 BASE) 时连远程
	SubCmdDiff   = "diff"   // 不带 -r 时本地（工作文件 vs BASE）；-r N:M 比较历史版本时连远程
	SubCmdLog    = "log"    // 必须连远程（提交历史不在工作副本中）
	SubCmdList   = "list"   // 必须连远程（即使 path 是 WC 也会转成 URL）
	SubCmdExport = "export" // 不带 -r 的 WC 导出读本地；URL 或 -r N 时连远程
	SubCmdStatus = "status" // 不加 -u 时读本地；-u 时连远程对比是否过期
)

const (
	ArgXml       = "--xml"
	ArgRecursive = "--recursive"
	ArgSummarize = "--summarize"
	ArgVerbose   = "--verbose"
	ArgRevision  = "--revision"
	ArgQuiet     = "--quiet"
)
