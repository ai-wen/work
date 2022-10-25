# [ cpu_features](https://github.com/syberia-project/platform_external_cpu_features)
经典的camke脚本，包含linux打包
C++ 编译器宏
cmake 编译器宏

[](https://github.com/pytorch/cpuinfo)
[](https://github.com/google/cpu_features)

# go 

https://cs.opensource.google/go/x/sys/+/master:cpu/cpu.go;bpv=0;bpt=0
https://pkg.go.dev/golang.org/x/sys/cpu#section-sourcefiles

```go
package main

import (
	"fmt"

	"golang.org/x/sys/cpu"
)

func main() {
	
	var supportSM4 = cpu.ARM64.HasSM4
	var supportsAES = cpu.X86.HasAES || cpu.ARM64.HasAES
	var supportsGFMUL = cpu.X86.HasPCLMULQDQ || cpu.ARM64.HasPMULL
	var useAVX2 = cpu.X86.HasAVX2 && cpu.X86.HasBMI2

	if supportSM4 {
		//return newCipherNI(key)
		fmt.Println("supportSM4")
	}

	if !supportsAES {
		//return newCipherGeneric(key)
	}

	//blocks := 4
	if useAVX2 {
		//blocks = 8
	}
	//c := &sm4CipherAsm{sm4Cipher{make([]uint32, rounds), make([]uint32, rounds)}, blocks, blocks * BlockSize}
	//expandKeyAsm(&key[0], &ck[0], &c.enc[0], &c.dec[0], INST_AES)
	if supportsGFMUL {
		//return &sm4CipherGCM{c}, nil
	}
	//return c, nil
}

```