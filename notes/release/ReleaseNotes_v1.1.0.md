# SVNArchiver

## Release Notes

+ 自 v1.0.2 起：升级 Go 与依赖，适配 Go module 构建，并增加 GitHub Actions 的 CI / 发版流程。归档功能本身无行为变更。

### Improvements

+ 升级 Go 版本至 1.24，并适配 Go module 模式编译。

+ 新增本地构造脚本 `build/build.bat`、`build/build.sh`，支持单元测试、编译可执行文件、打包源代码。

+ 新增 GitHub Actions CI：在 `main` 的推送与 Pull Request 上自动构建与测试。

+ 新增 GitHub Actions Release：在 `main` 上推送 `v*.*.*` tag 时交叉编译并发布 GitHub Release。

+ 新增 GitHub Actions ReleaseNote：可在网页上手动运行，用 `notes/release/ReleaseNotes_<tag>.md` 更新已有 Release 的说明正文。

### Changes

+ 依赖 `github.com/xuzhuoxi/infra-go` 由 v1.0.2 更新至 v1.4.1。

+ 依赖 `golang.org/x/text` 由 v0.3.7 更新至 v0.34.0。

+ 更新 README 中英文文档样式与说明。

### Library Changes

+ `infra-go`：v1.0.2 → v1.4.1

+ `golang.org/x/text`：v0.3.7 → v0.34.0
