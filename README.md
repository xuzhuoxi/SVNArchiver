# SVNArchiver

SVN归档工具

## 兼容性

go 1.16.15

## 获取


可通过下载执行文件或下载源码编译获得执行文件

- 执行文件[下载地址](https://github.com/xuzhuoxi/SVNArchiver/releases)。

- 通过github下载源码编译

	1. 执行代码。

	```
	go get -u github.com/xuzhuoxi/SVNArchiver
	```
	
	2. 编译工程。 

	Windows下执行[goxc_build.bat](/build/goxc_build.bat)
	
	Linux下执行[goxc_build.sh](/build/goxc_build.sh)

## 运行

程序只允许通过命令行运行

支持的命令行参数包括：-vlist, -arch, -v, -v0, -v1, -path

```
SVNArchiver -vlist=50 -path=处理目录路径
```

```
SVNArchiver -arch=归档路径 -path=处理目录路径
```

```
SVNArchiver -arch=归档路径 -v0=旧版本号 -v1=目标版本号 -path=处理目录路径
```

归档路径格式： path-version.zip path-version0-version1.zip


## 功能说明

- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.export.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.info.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.list.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.status.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.update.html


## 依赖性

- infra-go(库依赖) [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc(编译依赖) [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## 联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> 或 <mailxuzhuoxi@163.com>

## 开源许可证

~~ExcelExporter 源代码基于[MIT许可证](/LICENSE)进行开源。~~