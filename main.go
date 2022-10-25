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