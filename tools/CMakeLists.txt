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

# csv2sql
set(GOLANG_ENTRY_POINT "${TOOLS_ROOT}/cmd/csv2sql/main.go")
set(TARGET_CSV2SQL "csv2sql")
add_golang_executable(${TARGET_CSV2SQL})


add_custom_target("clean_${TARGET_CSV2SQL}"
  COMMAND rm ${TARGET_CSV2SQL}
  WORKING_DIRECTORY ${TOOLS_OUTPUT}
  COMMENT "Remove ${TOOLS_OUT}/${TARGET_CSV2SQL}"
)

add_custom_target(clean_tools
  DEPENDS "clean_${TARGET_CSV2SQL}"
  COMMENT "Remove ${TOOLS_OUT}/*"
)
