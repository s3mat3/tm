cmake_minimum_required(VERSION 3.24)

set(TARGET tools)

project(${TARGET}
  VERSION ${REPOS_VERSION}
  LANGUAGES NONE)

if(NOT TM_BUILD_TOOLS)
  message("Build request from direct")
  set(TOOLS_BUILD_DIRECTORY ${PROJECT_BINARY_DIR})
  set(TOOLS_ROOT ${PROJECT_SOURCE_DIR})
  set(TOOLS_SCRIPTS ${TOOLS_ROOT}/../scripts)
  set(TM_ROOT_DIRECTORY ${TOOLS_ROOT}/..)
else ()
  message("Build requested by tm root CMakeLists")
  set(TOOLS_BUILD_DIRECTORY ${TM_BUILD_DIRECTORY})
  set(TOOLS_ROOT ${TM_TOOLS_ROOT})
  set(TOOLS_SCRIPTS ${TM_SCRIPTS})
endif ()

include(${TM_ROOT_DIRECTORY}/cmake/golang.cmake)

#######################
# Setup build options #
#######################
set(TARGET_DIR "debug")

if (${CMAKE_BUILD_TYPE} STREQUAL "Release")
  set(TARGET_DIR "release")
  set(GOLANG_BUILD_FLAGS -tags release)
  set(GOLANG_BUILD_LDFLAGS -s -w)
else ()
  set(TARGET_DIR "debug")
  set(GOLANG_BUILD_FLAGS -tags debug)
  #set(GOLANG_BUILD_LDFLAGS -race)
endif ()

set(TOOLS_OUTPUT "${TOOLS_BUILD_DIRECTORY}/${TARGET}/${TARGET_DIR}")

set(GOLANG_GOMOD_DIRECTORY ${TM_ROOT})
set(GOLANG_BUILD_OUT_DIRECTORY ${TOOLS_OUTPUT})

set(BUILD_TOOLS_CSV2SQL On)
add_subdirectory(${TOOLS_ROOT}/csv2sql)
