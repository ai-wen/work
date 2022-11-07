# 下载
https://golang.google.cn/dl/

# golang cgo windows mingw64 环境搭建
MingW 分 32位和64位版本：下载地址分别如下：
http://sourceforge.net/projects/mingw/
http://sourceforge.net/projects/mingw-w64/

gcc 主要有三种不同的线程库的定义，分别是 Win32，OS/2，以及 POSIX
前两种定义只适合于他们各自的平台，而 POSIX wiki 定义的线程库是适用于所有的计算平台的，故肯定使用 threads-posix

C++ Exceptions有三种处理方式
DWARF 不传递错误，需要使用DWARF-2（或DWARF-3）调试信息，生成的包会很大
SJLJ 长跳转，即使抛出异常也正常的执行，可以在没有使用GCC编译的代码或者没有调用堆栈展开信息的代码中工作
SEH 结构化异常处理，Windows使用自己的异常处理机制

[x86_64-8.1.0-release-posix-seh-rt_v6-rev0 选择](https://nchc.dl.sourceforge.net/project/mingw-w64/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z)

# 网络访问 配置国内代理
go env -w GOPROXY=https://goproxy.cn

# VSCOD 使用go
命令行设置
go env -w GO111MODULE=off
go env -w GOPROXY=https://goproxy.cn,direct

安装插件 Go 和 Go Nightly
Ctrl+Shift+P
输入 Go:Install/Update Tools
全部勾选，再点击Ok安装

# 使用Go Modules管理依赖
这是官方推荐的为了替代GOPATH而诞生的一个Go语言依赖库管理器。之前所有的包都丢在GOPATH中，它的最大的好处就是，我们依赖的包可以指定版本。其次所有程序依赖的包，只会存在同一份。不会像npm那样，同一个包还能有n多个存在。这样我们的电脑就很省空间了
使用起来也非常简单，常用命令就一个go mod tidy，通俗来说就是将当前的库源码文件所依赖的包，全部安装并记录下来，多的包就删掉，少了的就自动补上
go mod init 

# 下载安装
go get -u github.com/motemen/gore
u 强制使用网络去更新包和它的依赖包
github.com 网站域名：表示代码托管的网站，类似于电子邮件 @ 后面的服务器地址。
motemen 作者或机构：表明这个项目的归属，一般为网站的用户名，如果需要找到这个作者下的所有项目，可以直接在网站上通过搜索“域名/作者”进行查看。这部分类似于电子邮件 @ 前面的部分。
gore 项目名：每个网站下的作者或机构可能会同时拥有很多的项目，图中标示的部分表示项目名称。

默认情况下，go get 可以直接使用。 go get github.com/motemen/gore
获取前，请确保 GOPATH 已经设置。Go 1.8 版本之后，GOPATH 默认在用户目录的 go 文件夹下。GOPATH=C:\Users\lw\go


删除：
直接删除源文件目录及编译后的package目录即可。
在源码目录$GOPATH/src下找到你要删除的package名，直接删除；
然后在$GOPATH/pkg/<architecture>下删除编译后的package目标文件目录。

go clean命令自动删除编译后的package目录，再手动删除源文件目录
	
# 编译 go 项目

go build main.go
go build src/main.go 


## 测试类编译
go help test

import (
	"testing"
)

src/xxx_test.go

cd src 
go test  -c  编译二进制
go test      运行
go test  -v  运行并输出中间结果


要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以_test结尾。默认的情况下，go test命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。
单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Test为前缀，例如：
func TestXXX( t *testing.T )


# 可执行文件 
package main  #必须

func main() {  #必须
	
	
}

# GO111MODULE环境变量

GO111MODULE=off，无模块支持，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找
GO111MODULE=on，模块支持，go命令行会使用modules，而一点也不会去GOPATH目录下查找

GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
（1）当前目录在GOPATH/src之外且该目录包含go.mod文件，开启模块支持。
（2）当前文件在包含go.mod文件的目录下面

go env -w GO111MODULE=auto
或者
go env -w GO111MODULE=off


# 下载包
cannot find package "golang.org/x/sys/cpu" in any of:
        D:\SW\GO\src\golang.org\x\sys\cpu (from $GOROOT)
        C:\Users\lm\go\src\golang.org\x\sys\cpu (from $GOPATH)

go get -u golang.org/x/sys


go env
手动从github上将该包clone下来
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/sys.git




## 编译动态库 静态库
```go
//动态库基本元素
package main		--必须	
import "C"       	--必须

//export TestFun    --必须和函数之间无空行 只有导出函数才会生成 .h头文件
func TestFun() (float64) {
	return 1
}

func main() {  //--必须
    // Need a main function to make CGO compile package as C shared library
}
```

go build -buildmode=c-shared -o test.dll testdll.go
go build -buildmode=c-archive -o test.a testdll.go


## 静态库包
go install xxpkg

编译静态库 生成.a文件
如果包在 %GOPATH%\src 目录下，编译静态库xxx.a 直接使用命令 go install randomness
静态库生成在 %GOPATH%\pkg\windows_amd64

使用静态库 生成.o文件
go tool compile -I %GOPATH%\pkg\windows_amd64 testa.go
-I选项指定了xxx包的安装路径，供testa.go导入使用

链接生成可执行文件
go tool link -o testa.exe -L %GOPATH%\pkg\windows_amd64 testa.o
-L选项指定了静态库xxx.a的路径


## go help buildmode命令可以查看C静态库和C动态库的构建说明

##  从非main包导出C函数，或者是多个包导出C函数
https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-06-static-shared-lib.html

## 32位
go env
set GOARCH=386 
set GOARCH=amd64

GOOS=windows GOARCH=386 go build main.go
GOOS=windows GOARCH=amd64 go build main.go
