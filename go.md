# 下载
https://golang.google.cn/dl/

# 网络访问 配置国内代理
go env -w GOPROXY=https://goproxy.cn


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
//go build -buildmode=c-shared -o test.dll testdll.go
//go build -buildmode=c-archive -o test.a testdll.go


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



```go
package main

import "C"

import (
	"fmt"
	"randomness"	
)

/*
[ 1] 单比特频数检测 MonoBitFrequencyTest
[ 2] 块内频数检测 FrequencyWithinBlockTest
[ 3] 扑克检测 PokerTest
[ 4] 重叠子序列检测 OverlappingTemplateMatchingTest
[ 5] 游程总数检测 RunsTest
[ 6] 游程分布检测 RunsDistributionTest
[ 7] 块内最大游程检测 LongestRunOfOnesInABlockTest
[ 8] 二元推导检测 BinaryDerivativeTest
[ 9] 自相关检测 AutocorrelationTest
[10] 矩阵秩检测 MatrixRankTest
[11] 累加和检测 CumulativeTest
[12] 近似熵检测 ApproximateEntropyTest
[13] 线型复杂度检测 LinearComplexityTest
[14] Maurer通用统计检测 MaurerUniversalTest
[15] 离散傅里叶检测 DiscreteFourierTransformTest
*/


//export MBFTest
func MBFTest(data []byte) (float64) {	
	p, q := randomness.MonoBitFrequencyTest(randomness.B2bitArr(data))

	if fmt.Sprintf("%.6f", p) != "0.215925" || fmt.Sprintf("%.6f", q) != "0.892038" {
		return 0
	}
	return 1
}

//export FWBTest
func FWBTest(data []byte) (float64){	
	p, q := randomness.FrequencyWithinBlockTest(randomness.B2bitArr(data))
	if fmt.Sprintf("%.6f", p) != "0.706438" || fmt.Sprintf("%.6f", q) != "0.706438" {
		return 0
	}
	return 1
}

//export PKTest
func PKTest(data []byte, m int) (float64){
	p, q := randomness.PokerProto(randomness.B2bitArr(data), m)

	if fmt.Sprintf("%.6f", p) != "0.213734" || fmt.Sprintf("%.6f", q) != "0.213734" {
		return 0
	}
	return 1
}


func main() {
    // Need a main function to make CGO compile package as C shared library
}


//编译动态库 静态库
//go build -buildmode=c-shared -o test.dll testdll.go
//go build -buildmode=c-archive -o test.a testdll.go

//https://github.com/Trisia/randomness  国密最新随机数检测规范



```

```go
package main

import (
	"fmt"
	"math/rand"
	"randomness"
	"randomness/detect"
		
)

func createDistributions(s, m int) [][]float64 {
	res := make([][]float64, m)
	for i := 0; i < m; i++ {
		res[i] = make([]float64, s)
	}
	return res
}

// PowerOnDetect 上电自检，15种检测，每组 10^6比特，分20组
func main() {

	s := 20
	t := detect.Threshold(s)

	buf := make([]byte, 1000_000/8)
	counters := make([]int, 15)

	distributions := createDistributions(s, 15)

	for i := 0; i < s; i++ {

		_, _ = rand.Read(buf)

		resArr := detect.Round15(buf)
		for idx, result := range resArr {
			distributions[idx][i] = result.Q
			if result.Pass {
				counters[idx]++
			}
		}
	}
	for i, n := range counters {
		if n < t {
			 fmt.Printf("%s %d/%d", randomness.TestMethodArr[i].Name, n, s)
			 return
		}
	}
	for i := range distributions {
		Pt := detect.ThresholdQ(distributions[i])
		if Pt < randomness.AlphaT {
			fmt.Printf("%s %f", randomness.TestMethodArr[i].Name, Pt)
			return 
		}
	}

	fmt.Printf("15种算法 上电自检 20组 10^6 \n")
	
	return
}
```