// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This program dumps current host data to the standard output.
// Output can be text or json if the `--json` flag is passed.

#include <assert.h>
#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "cpu_features_macros.h"

#include "cpuinfo_x86.h"
#include "cpuinfo_arm.h"
#include "cpuinfo_aarch64.h"
#include "cpuinfo_mips.h"
#include "cpuinfo_ppc.h"


int main()
{
    X86Info* info = GetX86Info();
   // const CacheInfo* cache_info = GetX86CacheInfo();

    /*
    bool supportSM4 = cpu.ARM64.HasSM4
    bool supportsAES = cpu.X86.HasAES || cpu.ARM64.HasAES;
    bool supportsGFMUL = cpu.X86.HasPCLMULQDQ || cpu.ARM64.HasPMULL;
    bool useAVX2 = cpu.X86.HasAVX2 && cpu.X86.HasBMI2;
    */

    ArmInfo* info1 = GetArmInfo();
   
    Aarch64Info* info2 = GetAarch64Info();
    
    MipsInfo* info3 = GetMipsInfo();
    
    PPCInfo* info4 = GetPPCInfo();
 
    return 0;
}

