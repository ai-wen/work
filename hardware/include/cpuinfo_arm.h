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

#ifndef CPU_FEATURES_INCLUDE_CPUINFO_ARM_H_
#define CPU_FEATURES_INCLUDE_CPUINFO_ARM_H_

#include <stdint.h>  // uint32_t
#include <stdbool.h>
#include "cpu_features_cache_info.h"
#include "cpu_features_macros.h"

CPU_FEATURES_START_CPP_NAMESPACE


// ARM64 contains the supported CPU features of the
// current ARMv8(aarch64) platform. If the current platform
// is not arm64 then all feature flags are false.
typedef struct {	
	bool		fp; // Floating-point instruction set (always available)
	bool		asimd; // Advanced SIMD (always available)
	bool		evtstrm; // Event stream support
	bool		aes; // AES hardware implementation
	bool		pmull; // Polynomial multiplication instruction set
	bool		sha1; // SHA1 hardware implementation
	bool		sha2; // SHA2 hardware implementation
	bool		crc32; // CRC32 hardware implementation
	bool		atomics; // Atomic memory operation instruction set
	bool		fphp; // Half precision floating-point instruction set
	bool		asimdhp; // Advanced SIMD half precision instruction set
	bool		cpuid; // CPUID identification scheme registers
	bool		asimdrdm; // Rounding double multiply add/subtract instruction set
	bool		jscvt; // Javascript conversion from floating-point to integer
	bool		fcam; // Floating-point multiplication and addition of complex numbers
	bool		lrcpc; // Release Consistent processor consistent support
	bool		dcpop; // Persistent memory support
	bool		sha3; // SHA3 hardware implementation
	bool		sm3; // SM3 hardware implementation
	bool		sm4; // SM4 hardware implementation
	bool		asimddp; // Advanced SIMD double precision instruction set
	bool		sha512; // SHA512 hardware implementation
	bool		sve; // Scalable Vector Extensions
	bool		asimdfhm; // Advanced SIMD multiplication FP16 to FP32	
}Arm64Features;


// ARM contains the supported CPU features of the current ARM (32-bit) platform.
typedef struct {
  bool swp : 1;       // SWP instruction (atomic read-modify-write)
  bool half : 1;      // Half-word loads and stores
  bool thumb : 1;     // Thumb (16-bit instruction set)
  bool _26bit : 1;    // "26 Bit" Model (Processor status register folded into
                   // program counter)
  bool fastmult : 1;  // 32x32->64-bit multiplication
  bool fpa : 1;       // Floating point accelerator
  bool vfp : 1;       // Vector Floating Point.
  bool edsp : 1;     // DSP extensions (the 'e' variant of the ARM9 CPUs, and all
                  // others above)
  bool java : 1;     // Jazelle (Java bytecode accelerator)
  bool iwmmxt : 1;   // Intel Wireless MMX Technology.
  bool crunch : 1;   // MaverickCrunch coprocessor
  bool thumbee : 1;  // ThumbEE
  bool neon : 1;     // Advanced SIMD.
  bool vfpv3 : 1;    // VFP version 3
  bool vfpv3d16 : 1;  // VFP version 3 with 16 D-registers
  bool tls : 1;       // TLS register
  bool vfpv4 : 1;     // VFP version 4 with fast context switching
  bool idiva : 1;     // SDIV and UDIV hardware division in ARM mode.
  bool idivt : 1;     // SDIV and UDIV hardware division in Thumb mode.
  bool vfpd32 : 1;    // VFP with 32 D-registers
  bool lpae : 1;     // Large Physical Address Extension (>4GB physical memory on
                  // 32-bit architecture)
  bool evtstrm : 1;  // kernel event stream using generic architected timer
  bool aes : 1;      // Hardware-accelerated Advanced Encryption Standard.
  bool pmull : 1;    // Polynomial multiply long.
  bool sha1 : 1;     // Hardware-accelerated SHA1.
  bool sha2 : 1;     // Hardware-accelerated SHA2-256.
  bool crc32 : 1;    // Hardware-accelerated CRC-32.

  // Make sure to update ArmFeaturesEnum below if you add a field here.
} ArmFeatures;

typedef struct {
  ArmFeatures features;
  int implementer;
  int architecture;
  int variant;
  int part;
  int revision;
} ArmInfo;

// TODO(user): Add macros to know which features are present at compile
// time.

ArmInfo* GetArmInfo(void);

// Compute CpuId from ArmInfo.
uint32_t GetArmCpuId(const ArmInfo* const info);

////////////////////////////////////////////////////////////////////////////////
// Introspection functions

typedef enum {
  ARM_SWP,
  ARM_HALF,
  ARM_THUMB,
  ARM_26BIT,
  ARM_FASTMULT,
  ARM_FPA,
  ARM_VFP,
  ARM_EDSP,
  ARM_JAVA,
  ARM_IWMMXT,
  ARM_CRUNCH,
  ARM_THUMBEE,
  ARM_NEON,
  ARM_VFPV3,
  ARM_VFPV3D16,
  ARM_TLS,
  ARM_VFPV4,
  ARM_IDIVA,
  ARM_IDIVT,
  ARM_VFPD32,
  ARM_LPAE,
  ARM_EVTSTRM,
  ARM_AES,
  ARM_PMULL,
  ARM_SHA1,
  ARM_SHA2,
  ARM_CRC32,
  ARM_LAST_,
} ArmFeaturesEnum;



CPU_FEATURES_END_CPP_NAMESPACE


#endif  // CPU_FEATURES_INCLUDE_CPUINFO_ARM_H_
