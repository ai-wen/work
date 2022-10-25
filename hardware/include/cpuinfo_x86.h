// Copyright 2017 Google LLC
// Copyright 2020 Intel Corporation
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

#ifndef CPU_FEATURES_INCLUDE_CPUINFO_X86_H_
#define CPU_FEATURES_INCLUDE_CPUINFO_X86_H_

#include "cpu_features_cache_info.h"
#include "cpu_features_macros.h"
#include <stdbool.h>

CPU_FEATURES_START_CPP_NAMESPACE

// See https://en.wikipedia.org/wiki/CPUID for a list of x86 cpu features.
// The field names are based on the short name provided in the wikipedia tables.
typedef struct {
  bool fpu : 1;
  bool tsc : 1;
  bool cx8 : 1;
  bool clfsh : 1;
  bool mmx : 1;
  bool aes : 1;
  bool erms : 1;
  bool f16c : 1;
  bool fma4 : 1;
  bool fma3 : 1;
  bool vaes : 1;
  bool vpclmulqdq : 1;
  bool bmi1 : 1;
  bool hle : 1;
  bool bmi2 : 1;
  bool rtm : 1;
  bool rdseed : 1;
  bool clflushopt : 1;
  bool clwb : 1;

  bool sse : 1;
  bool sse2 : 1;
  bool sse3 : 1;
  bool ssse3 : 1;
  bool sse4_1 : 1;
  bool sse4_2 : 1;
  bool sse4a : 1;

  bool avx : 1;
  bool avx2 : 1;

  bool avx512f : 1;
  bool avx512cd : 1;
  bool avx512er : 1;
  bool avx512pf : 1;
  bool avx512bw : 1;
  bool avx512dq : 1;
  bool avx512vl : 1;
  bool avx512ifma : 1;
  bool avx512vbmi : 1;
  bool avx512vbmi2 : 1;
  bool avx512vnni : 1;
  bool avx512bitalg : 1;
  bool avx512vpopcntdq : 1;
  bool avx512_4vnniw : 1;
  bool avx512_4vbmi2 : 1;
  bool avx512_second_fma : 1;
  bool avx512_4fmaps : 1;
  bool avx512_bf16 : 1;
  bool avx512_vp2intersect : 1;
  bool amx_bf16 : 1;
  bool amx_tile : 1;
  bool amx_int8 : 1;

  bool pclmulqdq : 1;
  bool smx : 1;
  bool sgx : 1;
  bool cx16 : 1;  // aka. CMPXCHG16B
  bool sha : 1;
  bool popcnt : 1;
  bool movbe : 1;
  bool rdrnd : 1;

  bool dca : 1;
  bool ss : 1;
  // Make sure to update X86FeaturesEnum below if you add a field here.
} X86Features;

typedef struct {
  X86Features features;
  int family;
  int model;
  int stepping;
  char vendor[13];  // 0 terminated string
} X86Info;

// Calls cpuid and returns an initialized X86info.
// This function is guaranteed to be malloc, memset and memcpy free.
X86Info* GetX86Info(void);

// Returns cache hierarchy informations.
// Can call cpuid multiple times.
// Only works on Intel CPU at the moment.
// This function is guaranteed to be malloc, memset and memcpy free.
CacheInfo* GetX86CacheInfo(void);

typedef enum {
  X86_UNKNOWN,
  INTEL_CORE,      // CORE
  INTEL_PNR,       // PENRYN
  INTEL_NHM,       // NEHALEM
  INTEL_ATOM_BNL,  // BONNELL
  INTEL_WSM,       // WESTMERE
  INTEL_SNB,       // SANDYBRIDGE
  INTEL_IVB,       // IVYBRIDGE
  INTEL_ATOM_SMT,  // SILVERMONT
  INTEL_HSW,       // HASWELL
  INTEL_BDW,       // BROADWELL
  INTEL_SKL,       // SKYLAKE
  INTEL_ATOM_GMT,  // GOLDMONT
  INTEL_KBL,       // KABY LAKE
  INTEL_CFL,       // COFFEE LAKE
  INTEL_WHL,       // WHISKEY LAKE
  INTEL_CNL,       // CANNON LAKE
  INTEL_ICL,       // ICE LAKE
  INTEL_TGL,       // TIGER LAKE
  INTEL_SPR,       // SAPPHIRE RAPIDS
  AMD_HAMMER,      // K8
  AMD_K10,         // K10
  AMD_BOBCAT,      // K14
  AMD_BULLDOZER,   // K15
  AMD_JAGUAR,      // K16
  AMD_ZEN,         // K17
} X86Microarchitecture;

// Returns the underlying microarchitecture by looking at X86Info's vendor,
// family and model.
X86Microarchitecture GetX86Microarchitecture(const X86Info* info);

// Calls cpuid and fills the brand_string.
// - brand_string *must* be of size 49 (beware of array decaying).
// - brand_string will be zero terminated.
// - This function calls memcpy.
void FillX86BrandString(char brand_string[49]);

const char* GetX86MicroarchitectureName(X86Microarchitecture);



CPU_FEATURES_END_CPP_NAMESPACE

#endif  // CPU_FEATURES_INCLUDE_CPUINFO_X86_H_
