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

#ifndef CPU_FEATURES_INCLUDE_CPUINFO_MIPS_H_
#define CPU_FEATURES_INCLUDE_CPUINFO_MIPS_H_

#include "cpu_features_cache_info.h"
#include "cpu_features_macros.h"
#include <stdbool.h>

CPU_FEATURES_START_CPP_NAMESPACE

typedef struct {
  bool msa : 1;  // MIPS SIMD Architecture
                // https://www.mips.com/products/architectures/ase/simd/
  bool eva : 1;  // Enhanced Virtual Addressing
                // https://www.mips.com/products/architectures/mips64/
  bool r6 : 1;   // True if is release 6 of the processor.

  // Make sure to update MipsFeaturesEnum below if you add a field here.
} MipsFeatures;

typedef struct {
  MipsFeatures features;
} MipsInfo;

MipsInfo* GetMipsInfo(void);

////////////////////////////////////////////////////////////////////////////////
// Introspection functions

typedef enum {
  MIPS_MSA,
  MIPS_EVA,
  MIPS_R6,
  MIPS_LAST_,
} MipsFeaturesEnum;



CPU_FEATURES_END_CPP_NAMESPACE



#endif  // CPU_FEATURES_INCLUDE_CPUINFO_MIPS_H_
