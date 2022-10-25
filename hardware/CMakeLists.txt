cmake_minimum_required(VERSION 3.0)

# option() honors normal variables.
# see: https://cmake.org/cmake/help/git-stage/policy/CMP0077.html
if(POLICY CMP0077)
  cmake_policy(SET CMP0077 NEW)
endif()

project(CpuFeatures VERSION 0.6.0 LANGUAGES C)

set(CMAKE_C_STANDARD 99)

# Default Build Type to be Release
if(NOT CMAKE_BUILD_TYPE)
  set(CMAKE_BUILD_TYPE "Release" CACHE STRING
      "Choose the type of build, options are: Debug Release RelWithDebInfo MinSizeRel."
      FORCE)
endif(NOT CMAKE_BUILD_TYPE)

option(BUILD_SHARED_LIBS "Build library as shared." ON)

if(BUILD_SHARED_LIBS AND UNIX)
  set(BUILD_PIC ON)
  #set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC ")
endif()


macro(setup_include_and_definitions TARGET_NAME)
  target_include_directories(${TARGET_NAME}
    PUBLIC  $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include>
    PRIVATE $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include/internal>
  )
  target_compile_definitions(${TARGET_NAME}
    PUBLIC STACK_LINE_READER_BUFFER_SIZE=1024
  )
endmacro()

#
# library : utils
#

add_library(utils OBJECT
  ${PROJECT_SOURCE_DIR}/include/internal/bit_utils.h
  ${PROJECT_SOURCE_DIR}/include/internal/filesystem.h
  ${PROJECT_SOURCE_DIR}/include/internal/stack_line_reader.h
  ${PROJECT_SOURCE_DIR}/include/internal/string_view.h
  ${PROJECT_SOURCE_DIR}/src/internal/filesystem.c
  ${PROJECT_SOURCE_DIR}/src/internal/stack_line_reader.c
  ${PROJECT_SOURCE_DIR}/src/internal/string_view.c
)
set_property(TARGET utils PROPERTY POSITION_INDEPENDENT_CODE ${BUILD_PIC})
#setup_include_and_definitions(utils)
target_include_directories(utils
    PUBLIC  $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include>
    PRIVATE $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include/internal>
  )
target_compile_definitions(utils
    PUBLIC STACK_LINE_READER_BUFFER_SIZE=1024
  )


#
# library : unix_based_hardware_detection
#

if(UNIX)
  add_library(unix_based_hardware_detection OBJECT
    ${PROJECT_SOURCE_DIR}/include/internal/hwcaps.h
    ${PROJECT_SOURCE_DIR}/src/internal/hwcaps.c
  )
  #setup_include_and_definitions(unix_based_hardware_detection)
  target_include_directories(unix_based_hardware_detection
    PUBLIC  $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include>
    PRIVATE $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include/internal>
    )
  target_compile_definitions(unix_based_hardware_detection
    PUBLIC STACK_LINE_READER_BUFFER_SIZE=1024
    )

  check_include_file(dlfcn.h HAVE_DLFCN_H)
  if(HAVE_DLFCN_H)
    target_compile_definitions(unix_based_hardware_detection PRIVATE HAVE_DLFCN_H)
  endif()
  check_symbol_exists(getauxval "sys/auxv.h" HAVE_STRONG_GETAUXVAL)
  if(HAVE_STRONG_GETAUXVAL)
    target_compile_definitions(unix_based_hardware_detection PRIVATE HAVE_STRONG_GETAUXVAL)
  endif()
  set_property(TARGET unix_based_hardware_detection PROPERTY POSITION_INDEPENDENT_CODE ${BUILD_PIC})
endif()



set(PROCESSOR_IS_MIPS FALSE)
set(PROCESSOR_IS_ARM FALSE)
set(PROCESSOR_IS_AARCH64 FALSE)
set(PROCESSOR_IS_X86 FALSE)
set(PROCESSOR_IS_POWER FALSE)

if(CMAKE_SYSTEM_PROCESSOR MATCHES "^mips")
  set(PROCESSOR_IS_MIPS TRUE)
elseif(CMAKE_SYSTEM_PROCESSOR MATCHES "^arm")
  set(PROCESSOR_IS_ARM TRUE)
elseif(CMAKE_SYSTEM_PROCESSOR MATCHES "^aarch64")
  set(PROCESSOR_IS_AARCH64 TRUE)
elseif(CMAKE_SYSTEM_PROCESSOR MATCHES "(x86_64)|(AMD64|amd64)|(^i.86$)")
  set(PROCESSOR_IS_X86 TRUE)
elseif(CMAKE_SYSTEM_PROCESSOR MATCHES "^(powerpc|ppc)")
  set(PROCESSOR_IS_POWER TRUE)
endif()


macro(add_cpu_features_headers_and_sources HDRS_LIST_NAME SRCS_LIST_NAME)
  list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpu_features_macros.h)
  list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpu_features_cache_info.h)
 #[[
  if(PROCESSOR_IS_MIPS)
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_mips.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_mips.c)
  elseif(PROCESSOR_IS_ARM)
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_arm.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_arm.c)
  elseif(PROCESSOR_IS_AARCH64)
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_aarch64.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_aarch64.c)
  elseif(PROCESSOR_IS_X86)
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_x86.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/internal/cpuid_x86.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_x86.c)
  elseif(PROCESSOR_IS_POWER)
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_ppc.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_ppc.c)
  else()
    message(FATAL_ERROR "Unsupported architectures ${CMAKE_SYSTEM_PROCESSOR}")
  endif()
]]
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_mips.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_mips.c)

      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_arm.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_arm.c)

      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_aarch64.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_aarch64.c)

      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_x86.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/internal/cpuid_x86.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_x86.c)
  
      list(APPEND ${HDRS_LIST_NAME} ${PROJECT_SOURCE_DIR}/include/cpuinfo_ppc.h)
      list(APPEND ${SRCS_LIST_NAME} ${PROJECT_SOURCE_DIR}/src/cpuinfo_ppc.c)

endmacro()



#
# library : cpu_features
#
set (CPU_FEATURES_HDRS)
set (CPU_FEATURES_SRCS)
add_cpu_features_headers_and_sources(CPU_FEATURES_HDRS CPU_FEATURES_SRCS)
list(APPEND CPU_FEATURES_SRCS $<TARGET_OBJECTS:utils>)

if(NOT PROCESSOR_IS_X86 AND UNIX)
  list(APPEND CPU_FEATURES_SRCS $<TARGET_OBJECTS:unix_based_hardware_detection>)
endif()

add_library(cpu_features ${CPU_FEATURES_HDRS} ${CPU_FEATURES_SRCS} ${PROJECT_SOURCE_DIR}/cpu_features.def)

set_target_properties(cpu_features PROPERTIES PUBLIC_HEADER "${CPU_FEATURES_HDRS}")

setup_include_and_definitions(cpu_features)
target_link_libraries(cpu_features PUBLIC ${CMAKE_DL_LIBS})
set_property(TARGET cpu_features PROPERTY POSITION_INDEPENDENT_CODE ${BUILD_PIC})
target_include_directories(cpu_features
  PUBLIC $<INSTALL_INTERFACE:${CMAKE_INSTALL_INCLUDEDIR}/cpu_features>
)
if(PROCESSOR_IS_X86)
  if(APPLE)
    target_compile_definitions(cpu_features PRIVATE HAVE_SYSCTLBYNAME)
  endif()
endif()
add_library(CpuFeature::cpu_features ALIAS cpu_features)

#
# ndk_compat
#

if(ANDROID)
add_subdirectory(ndk_compat)
endif()


#
# program : list_cpu_features
#

add_executable(list_cpu_features ${PROJECT_SOURCE_DIR}/src/test/list_cpu_features.c)
target_link_libraries(list_cpu_features PRIVATE cpu_features)
add_executable(CpuFeature::list_cpu_features ALIAS list_cpu_features)
