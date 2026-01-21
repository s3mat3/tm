cmake_minimum_required(VERSION 3.24)

set(BACKEND_TARGET tm-backend)
project(${BACKEND_TARGET}
  VERSION ${REPOS_VERSION}
  LANGUAGES NONE)

if(NOT TM_BUILD_BACKEND)
  set(BACKEND_BUILD_DIRECTORY ${PROJECT_BINARY_DIR})
  set(BACKEND_ROOT ${PROJECT_SOURCE_DIR})
  set(BACKEND_SCRIPTS ${BACKEND_ROOT}/../scripts)
else ()
  set(BACKEND_BUILD_DIRECTORY ${TM_BUILD_DIRECTORY})
  set(BACKEND_ROOT ${TM_BACKEND_ROOT})
  set(BACKEND_SCRIPTS ${TM_SCRIPTS})
endif ()

set(BACKEND_TOOLS "${BACKEND_ROOT}/tools")

include(${BACKEND_ROOT}/cmake/golang.cmake)

#######################
# Setup build options #
#######################
set(SUFFIX "debug")

if (${CMAKE_BUILD_TYPE} STREQUAL "Release")
  set(SUFFIX "release")
  set(GOLANG_BUILD_FLAGS -tags release)
  set(GOLANG_BUILD_LDFLAGS -s -w)
else ()
  set(SUFFIX "debug")
  set(GOLANG_BUILD_FLAGS -tags debug)
  #set(GOLANG_BUILD_LDFLAGS -race)
endif ()

set(BACKEND_OUTPUT_ROOT ${BACKEND_BUILD_DIRECTORY})
set(BACKEND_OUTPUT "${BACKEND_OUTPUT_ROOT}/backend/${SUFFIX}")

set(GOLANG_GOMOD_DIRECTORY ${BACKEND_ROOT})
set(GOLANG_BUILD_OUT_DIRECTORY ${BACKEND_OUTPUT})

set(BUILD_TOOLS_CSV2SQL On)
add_subdirectory(${BACKEND_TOOLS}/csv2sql)
