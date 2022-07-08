# SVNArchiver

SVN归档工具

## <span id="a1">兼容性</span>

go 1.16.15

## <span id="a2">开始</span>

### <span id="a2.1">1. 下载</span>

+ 下载发行版本 [这里](https://github.com/xuzhuoxi/SVNArchiver/releases).

+ 下载仓库:

  ```sh
  go get -u github.com/xuzhuoxi/SVNArchiver
  ```

### <span id="a2.2">2. 构建<span>

+ 如果你已经下载整个仓库及相关依赖仓库，你可以执行构造脚本获得执行程序。

  + Windows下执行[goxc_build.bat](/build/goxc_build.bat)

  + Linux下执行[goxc_build.sh](/build/goxc_build.sh)

+ 如有必要，你可以修改 ([goxc_build.sh](/goxc_build/build.sh))来进行自定义的构造，构造工具的说明在[这里](https://github.com/laher/goxc).

### <span id="a2.3">3. 运行<span>

+ 仅支持命令行运行

+ 参考命令：

  `SVNArchiver -version -path=处理目录路径`

  `SVNArchiver -list=50 -path=处理目录路径`

  `SVNArchiver -arch=归档路径 -v=目标版本号 -path=处理目录路径`

  `SVNArchiver -arch=归档路径 -v=旧版本号:目标版本号 -path=处理目录路径`

## <span id="a3">用户手册<span>

### <span id="a3.1">3.1 参数说明<span>

+ -env

  指定环境路径

+ -list

+ -target

+ -arch

+ -v

+ -d

### <span id="a3.1">3.2 功能说明<span>

+ 查询目录的版本区间

  + **示例**： 
  
    `SVNArchiver -version -target=处理目录路径`
  
  + 结果会在控制能输出版本区域信息，如： `Min=0, Max=9`

+ 查询目录的版本详细列表信息

  + **示例**： 
  
    `SVNArchiver -list=50 -target=处理目录路径`
  
  + **-list**： 支持一个大于0的整数，表示最大版本条目数量，顺序为按时间近到远。

  + 结果会在控制能输出版本详细列表信息，如： 
      
    `****************************`

+ 默认归档

  + 使用**最近版本**与**上一版本**进行差异归档。

  + **示例**： 
  
    `SVNArchiver -evn=环境路径 -arch=归档路径 -target=处理目录路径`

+ 使用版本号信息进行差异归档

  + **示例**： 
  
    `SVNArchiver -evn=环境路径 -arch=归档路径 -v=开始版本号:目标版本号 -target=处理目录路径`
    `SVNArchiver -evn=环境路径 -arch=归档路径 -v=目标版本号 -target=处理目录路径`

  + **-v**： 版本号参数信息
    
	+ **只包含**目标版本时，表示使用**目标版本**与**上一版本**进行差异归档。

	+ **同时**包含**开始版本**与**目标版本**时，表示使用**开始版本**与**目标版本**进行差异归档。

+ 使用日期信息进行差异归档

  + **示例**： 
  
    `SVNArchiver -evn=环境路径 -arch=归档路径 -d=开始时间:目标时间 -target=处理目录路径`
    `SVNArchiver -evn=环境路径 -arch=归档路径 -d=目标时间 -target=处理目录路径`

  + **-d**： 时间参数信息
    
	+ **只包含**目标时间时，表示使用目标时间**向前最近**的版本与**上一版本**进行差异归档。

	+ **同时**包含**开始时间**与**目标时间**时，表示使用目标时间**向前最近**的版本与开始时间**向前最近**的版本进行差异归档。

### <span id="a3.2">3.2 注意<span>

1. 要求安装svn客户工具，并设置为环境路径。


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