# SVNArchiver

SVN归档工具

中文 | [English](README_EN.md)

## <span id="a1">兼容性</span>

go 1.16.15

## <span id="a2">开始</span>

### <span id="a2.1">2.1 下载</span>

+ 下载发行版本 [这里](https://github.com/xuzhuoxi/SVNArchiver/releases).

+ 下载仓库:

  ```sh
  go get -u github.com/xuzhuoxi/SVNArchiver
  ```

### <span id="a2.2">2.2 构建<span>

+ 如果你已经下载整个仓库及相关依赖仓库，你可以执行构造脚本获得执行程序。

  + Windows下执行[goxc_build.bat](/build/goxc_build.bat)

  + Linux下执行[goxc_build.sh](/build/goxc_build.sh)

+ 如有必要，你可以修改 ([goxc_build.sh](/goxc_build/build.sh))来进行自定义的构造，构造工具的说明在[这里](https://github.com/laher/goxc).

### <span id="a2.3">2.3 运行<span>

+ 仅支持命令行运行

+ 参考命令：

  + 查询提交信息功能： 

    + [查询提交信息](#a3.2.1)

      `SVNArchiver -log=50 -target=Svn目录`

  + 完整归档功能：  

    + [版本号完整归档](#a3.2.2)

      `SVNArchiver -env=环境路径 -r=12 -target=与环境路径相关的相对Svn目录 -arch=归档文件路径_{d}_{r}.zip`

    + [时间点完整归档](#a3.2.3)

      `SVNArchiver -d=20060102T150405 -target=Svn目录 -arch=归档文件路径_{d}_{r}.zip`

  + 差异归档功能：  

    + [版本号差异归档](#a3.2.4)

      `SVNArchiver -env=环境路径 -r0=1 -r1=15 -target=与环境路径相关的相对Svn目录 -arch=归档文件路径_{d0}_{d1}_{r0}_{r1}.zip`

    + [时间点差异归档](#a3.2.5)

      `SVNArchiver -d0=20220703 -d1=20220709T150405 -target=Svn目录 -arch=归档文件路径_{d0}_{d1}_{r0}_{r1}.zip`

## <span id="a3">用户手册<span>

### <span id="a3.1">3.1 参数说明<span>

+ <span id="a3.1.1">-env<span>

  环境路径, 用于指定一个基础路径，覆盖执行程序"SVNArchiver"的目录，然后-target和-arch则可使用相对路径。
  
  若不指定，则使用执行文件所在目录作为环境路径。

+ <span id="a3.1.2">-log<span>

  查询条目显示的最大数量，要求>=0, 当值为0时，认定为无限制。

+ <span id="a3.1.3">-r<span>

  完整归档时使用, 用于指定具体版本号，并使用该版本号(或向前最近的版本号)进行归档。

+ <span id="a3.1.4">-d<span>

  完整归档时使用, 用于指定一个时间点，并使用该时间点上的版本号(或向前最近的版本号)进行归档。

+ <span id="a3.1.5">-r0<span>
 
  差异归档时使用, 用于指定起始版本号。

+ <span id="a3.1.6">-r1<span>

  差异归档时使用, 用于指定结束版本号。

+ <span id="a3.1.7">-d0<span>

  差异归档时使用, 用于指定起始时间。

+ <span id="a3.1.8">-d1<span>

  差异归档时使用, 用于指定结束时间。

+ **注意**： 

  1. 版本号指定**要求确定**，如果svn目录上为**非连贯版本号**且没有指定的版本号，则使用**向前最近**的版本号代替。

  2. 时间指定**不要求确定**，最终使用的时当前时间上的或向前最近的版本号。

  3. 使用版本号配置进行差异归档时，-r0与-r1至少有一个。 当只有-r0时，-r1会设置为**最大的版本号**。 当只有-r1时，-r0全设置为**相对于-r1**的**上一个版本号**。
  
  4. 使用时间配置进行差异归档时，-d0与-d1至少有一个。 当只有-d0时，-d1会设置为**最新的时间**。 当只有-d1时，-d0会设置为相对-d1关联到的版本号的**上一个版本号时间**。

  5. 带"r"的参数与带"d"的参数不应该同时使用。
  
  6. -arch可使用固定路径，也可以使用通配符路径, 通配符会被**实际使用的**版本号或时间填充。
  
  7. 完整归档时-arch支持带**{d}**、**{r}**的通配符路径。

  8. 差异归档时-arch支持带**{d0}**、**{d1}**、**{r0}**、**{r1}**的通配符路径。

### <span id="a3.2">3.2 功能说明<span>

+ <span id="a3.2.1">查询提交信息功能<span>
	 
  示例： `SVNArchiver -log=50 -target=Svn目录`

  + -log: **非必要**
  
    打印条目的数量限制，没有则代表没限制。

+ <span id="a3.2.2">基于版本号的完整归档<span>

  示例： `SVNArchiver -env=环境路径 -r=12 -target=归档Svn目录 -arch=归档文件路径_{d}_{r}.zip`

  + -env:**非必要**
   
	指定一个新的环境路径，用于覆盖执行程序"SVNArchiver"的目录。

  + -r:**必要**

	归档版本号，要求是svn仓库中真实存在的。

  + -target:**必要**

	进行归档的svn目录(可以不是svn根目录)。 支持绝对路径，也支持相对于-env的相对路径。

  + -arch:**必要**

	归档文件保存路径。 支持相对路径，也支持相对于-env的相对路径。 路径支持通配符"{d}"、"{r}"。

+ <span id="a3.2.3">基于时间点的完整归档<span>

  示例： `SVNArchiver -env=环境路径 -d=20060102T150405 -target=归档Svn目录 -arch=归档文件路径_{d}_{r}.zip`

  + -env:**非必要**
   
	指定一个新的环境路径，用于覆盖执行程序"SVNArchiver"的目录。

  + -d:**必要**

	归档时间点，实际使用的是**<=时间点**下**最近**的版本号进行归档。

  + -target:**必要**

	进行归档的svn目录(可以不是svn根目录)。 支持绝对路径，也支持相对于-env的相对路径。

  + -arch:**必要**

	归档文件保存路径。 支持相对路径，也支持相对于-env的相对路径。 路径支持通配符"{d}"、"{r}"。

+ <span id="a3.2.4">基于版本号的差异归档<span>

  示例： `SVNArchiver -env=环境路径 -r0=1 -r1=12 -target=归档Svn目录 -arch=归档文件路径_{d0}_{d1}_{r0}_{r1}.zip`

  + -env:**非必要**
   
	指定一个新的环境路径，用于覆盖执行程序"SVNArchiver"的目录。

  + -r0 和 -r1:**至少有一个**

	归档版本号。 -r0为起始版本号, -r1为结束版本号。

    当只有-r0时，-r1会设置为**最大的版本号**。 当只有-r1时，-r0全设置为**相对于-r1**的**上一个版本号**。

  + -target:**必要**

	进行归档的svn目录(可以不是svn根目录)。 支持绝对路径，也支持相对于-env的相对路径。

  + -arch:**必要**

	归档文件保存路径。 支持相对路径，也支持相对于-env的相对路径。 路径支持通配符"{d0}"、"{d1}"、"{r0}"、"{r1}"。

+ <span id="a3.2.5">基于时间点的差异归档<span>

  示例： `SVNArchiver -env=环境路径 -d0=20220703 -d1=20220709T150405 -target=归档Svn目录 -arch=归档文件路径_{d0}_{d1}_{r0}_{r1}.zip`

  + -env:**非必要**
   
	指定一个新的环境路径，用于覆盖执行程序"SVNArchiver"的目录。

  + -d0 和 -d1:**至少有一个**

	归档时间点。 -d0为起始时间点, -d1为结束时间点。

    当只有-d0时，-d1会设置为**最新的时间**。 当只有-d1时，-d0会设置为相对-d1关联到的版本号的**上一个版本号时间**。

  + -target:**必要**

	进行归档的svn目录(可以不是svn根目录)。 支持绝对路径，也支持相对于-env的相对路径。

  + -arch:**必要**

	归档文件保存路径。 支持相对路径，也支持相对于-env的相对路径。 路径支持通配符"{d0}"、"{d1}"、"{r0}"、"{r1}"。

### <span id="a3.3">3.3 注意<span>

+ 要求安装svn客户端工具，并设置为环境变量。

## <span id="a4">参考文献<span>

- https://svnbook.red-bean.com/zh/1.8/svn.ref.svnversion.re.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.log.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.diff.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.export.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.list.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.status.html
- https://svnbook.red-bean.com/zh/1.8/svn.ref.svn.c.update.html

## 依赖性

- infra-go(库依赖) [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc(编译依赖) [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## 联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> 或 <mailxuzhuoxi@163.com>

## 开源许可证

~~SVNArchiver 源代码基于[MIT许可证](/LICENSE)进行开源。~~