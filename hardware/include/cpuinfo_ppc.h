// Copyright 2018 IBM
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

#ifndef CPU_FEATURES_INCLUDE_CPUINFO_PPC_H_
#define CPU_FEATURES_INCLUDE_CPUINFO_PPC_H_

#include "cpu_features_cache_info.h"
#include "cpu_features_macros.h"
#include "internal/hwcaps.h"
#include <stdbool.h>

CPU_FEATURES_START_CPP_NAMESPACE

typedef struct {
  bool ppc32 : 1;
  bool ppc64 : 1;
  bool ppc601 : 1;
  bool altivec : 1;
  bool fpu : 1;
  bool mmu : 1;
  bool mac_4xx : 1;
  bool unifiedcache : 1;
  bool spe : 1;
  bool efpsingle : 1;
  bool efpdouble : 1;
  bool no_tb : 1;
  bool power4 : 1;
  bool power5 : 1;
  bool power5plus : 1;
  bool cell : 1;
  bool booke : 1;
  bool smt : 1;
  bool icachesnoop : 1;
  bool arch205 : 1;
  bool pa6t : 1;
  bool dfp : 1;
  bool power6ext : 1;
  bool arch206 : 1;
  bool vsx : 1;
  bool pseries_perfmon_compat : 1;
  bool truele : 1;
  bool ppcle : 1;
  bool arch207 : 1;
  bool htm : 1;
  bool dscr : 1;
  bool ebb : 1;
  bool isel : 1;
  bool tar : 1;
  bool vcrypto : 1;
  bool htm_nosc : 1;
  bool arch300 : 1;
  bool ieee128 : 1;
  bool darn : 1;
  bool scv : 1;
  bool htm_no_suspend : 1;

  // Make sure to update PPCFeaturesEnum below if you add a field here.
} PPCFeatures;

typedef struct {
  PPCFeatures features;
} PPCInfo;

// This function is guaranteed to be malloc, memset and memcpy free.
PPCInfo* GetPPCInfo(void);

typedef struct {
  char platform[64];  // 0 terminated string
  char model[64];     // 0 terminated string
  char machine[64];   // 0 terminated string
  char cpu[64];       // 0 terminated string
  PlatformType type;
} PPCPlatformStrings;

PPCPlatformStrings GetPPCPlatformStrings(void);

////////////////////////////////////////////////////////////////////////////////
// Introspection functions

typedef enum {
  PPC_32,          /* 32 bit mode execution */
  PPC_64,          /* 64 bit mode execution */
  PPC_601_INSTR,   /* Old POWER ISA */
  PPC_HAS_ALTIVEC, /* SIMD Unit*/
  PPC_HAS_FPU,     /* Floating Point Unit */
  PPC_HAS_MMU,     /* Memory management unit */
  PPC_HAS_4xxMAC,
  PPC_UNIFIED_CACHE,  /* Unified instruction and data cache */
  PPC_HAS_SPE,        /* Signal processing extention unit */
  PPC_HAS_EFP_SINGLE, /* SPE single precision fpu */
  PPC_HAS_EFP_DOUBLE, /* SPE double precision fpu */
  PPC_NO_TB,          /* No timebase */
  PPC_POWER4,
  PPC_POWER5,
  PPC_POWER5_PLUS,
  PPC_CELL,  /* Cell broadband engine */
  PPC_BOOKE, /* Embedded ISA */
  PPC_SMT,   /* Simultaneous multi-threading */
  PPC_ICACHE_SNOOP,
  PPC_ARCH_2_05, /* ISA 2.05 - POWER6 */
  PPC_PA6T,      /* PA Semi 6T core ISA */
  PPC_HAS_DFP,   /* Decimal floating point unit */
  PPC_POWER6_EXT,
  PPC_ARCH_2_06,              /* ISA 2.06 - POWER7 */
  PPC_HAS_VSX,                /* Vector-scalar extension */
  PPC_PSERIES_PERFMON_COMPAT, /* Set of backwards compatibile performance
                                 monitoring events */
  PPC_TRUE_LE,
  PPC_PPC_LE,
  PPC_ARCH_2_07,      /* ISA 2.07 - POWER8 */
  PPC_HTM,            /* Hardware Transactional Memory */
  PPC_DSCR,           /* Data stream control register */
  PPC_EBB,            /* Event base branching */
  PPC_ISEL,           /* Integer select instructions */
  PPC_TAR,            /* Target address register */
  PPC_VEC_CRYPTO,     /* Vector cryptography instructions */
  PPC_HTM_NOSC,       /* Transactions aborted when syscall made*/
  PPC_ARCH_3_00,      /* ISA 3.00 - POWER9 */
  PPC_HAS_IEEE128,    /* VSX IEEE Binary Float 128-bit */
  PPC_DARN,           /* Deliver a random number instruction */
  PPC_SCV,            /* scv syscall */
  PPC_HTM_NO_SUSPEND, /* TM w/out suspended state */
  PPC_LAST_,
} PPCFeaturesEnum;



CPU_FEATURES_END_CPP_NAMESPACE


#endif  // CPU_FEATURES_INCLUDE_CPUINFO_PPC_H_
