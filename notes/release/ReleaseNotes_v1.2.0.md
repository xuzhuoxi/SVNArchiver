## Release Notes

+ 自 v1.1.0 起：调整 SVN 日志查询策略——`QueryLog` 改为远程查询全部提交记录，并新增按版本区间查询的 `QueryLogRange`；归档错误会正确向上返回。同时完善发版工作流与文档。

### Notable Changes

+ `svn.QueryLog` 不再用 `svnversion` 的本地版本区间截断日志，改为远程查询全部提交记录。

+ 新增 `svn.QueryLogRange`，按指定版本区间远程查询提交记录。

+ 版本号差异归档改为使用 `QueryLogRange`，只拉取任务所需区间的日志。

### Improvements

+ 新增 GitHub Actions ReleaseNote：可在网页上手动运行，用 `notes/release/ReleaseNotes_<tag>.md` 更新已有 Release 的说明正文。

+ Release 工作流改为把手写说明与 GitHub 自动生成 notes 合并写入正文；若 Release 已存在则覆盖附件并更新说明。

+ README 补充 TortoiseSVN 命令行工具安装说明，以及 svn 本地/远程命令对照。

+ 补充各包注释，并格式化源代码。

### Breaking Changes

+ `src/svnversion` 包已移除，相关能力迁入 `src/svn`；`QueryVersion` 重命名为 `QueryLocalVersion`。

+ `HandleDateArch`、`HandleRevArch`、`HandleDateDiffArch`、`HandleRevDiffArch` 的返回值由单独的 `archPath` 改为 `(archPath, error)`。

### API Changes

+ 新增 `svn.QueryLogRange(path, rMin, rMax)`。

+ `model.LogResult` 新增 `HandleLogs`、`MinRev`、`MaxRev`；解析日志后按版本号升序整理。

+ `LogResultEntry.String` 改为指针接收者。

### Notable Fixes

+ 归档过程中的 SVN 导出、打包、日志查询错误会返回给调用方，而不再只打日志后继续当作成功。

### Fixes

+ `getArchPathD` / `getArchDiffPathD` 在找不到对应 log entry 时返回错误，不再忽略。

+ 版本号差异归档在版本区间无效时提前失败。

+ 归档任务失败时不再写入空的归档信息。

### Changelog

+ `67f0e94` svn.QueryLog 不再优先按本地版本范围查询；新增 QueryLogRange；补充注释与 README。

+ `61c0a9d` 格式化源代码，补充代码说明，更新 README。

+ `387b602` 更新 README。

+ `9120005` 更新 v1.1.0 的 Release 说明。

+ `1607dc6` 更新 ReleaseNote 工作流说明，并更新 v1.1.0 的 Release 说明。

+ `ea417e0` 新增 ReleaseNote 工作流，用于更新已有 Release 的显示内容。

**Full Changelog**: https://github.com/xuzhuoxi/SVNArchiver/compare/v1.1.0...v1.2.0
