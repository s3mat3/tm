#[[
usage cmake -DCMAKE_BUILD_TYPE=(Debug | Release | '') \
 -DTM_BUILD_BACKEND=(on | off default) \
 -DTM_BUILD_FRONTEND=(on | off default) \
 -DTM_BACKEND_TEST=(on | off default) \
 -DTM_FRONTEND_TEST=(on | off default)
]]
cmake_minimum_required(VERSION 3.24)

file(STRINGS VERSION TM_VERSION_LINE)
string(REGEX MATCH "([0-9]+\\.?)+" REPOS_VERSION ${TM_VERSION_LINE})
execute_process (COMMAND git rev-parse --short HEAD
  OUTPUT_VARIABLE REPOS_REVISION OUTPUT_STRIP_TRAILING_WHITESPACE)

message("Repository veresion: ${REPOS_VERSION} ${REPOS_REVISION}")

set(TARGET tman)
project("${TARGET}"
  VERSION ${REPOS_VERSION}
  LANGUAGES NONE)

set(TM_ROOT_NAME ${TARGET})
###############
# directories #
###############
set(TM_ROOT_DIRECTORY ${PROJECT_SOURCE_DIR})
set(TM_BUILD_DIRECTORY ${PROJECT_BINARY_DIR})

set(TM_SCRIPTS ${TM_ROOT_DIRECTORY}/scripts)

set(TM_FRONTEND_ROOT ${TM_ROOT_DIRECTORY}/frontend)
set(TM_BACKEND_ROOT ${TM_ROOT_DIRECTORY}/backend)
set(TM_DOCKER_ROOT ${TM_ROOT_DIRECTORY}/docker)

## output directories
set(TM_OUTPUT_ROOT ${TM_BUILD_DIRECTORY})

###########
# options #
###########
option(BUILD_BACKEND "Server side compile." OFF)
option(BUILD_BACKEND_TEST "Test for server side" OFF)
option(BUILD_FRONTEND "Client side compile." OFF)
option(BUILD_FRONTEND_TEST "Test for client side" OFF)

if(BUILD_BACKEND)
  set(TM_BUILD_BACKEND ON)
  add_subdirectory(${TM_BACKEND_ROOT})
endif()

if(BUILD_FRONTEND)
  set(TM_BUILD_FRONTEND ON)
  add_subdirectory(${TM_FRONTEND_ROOT})
endif()

################
# make targets #
################
add_custom_target(up
  COMMAND ${TM_SCRIPTS}/up
  WORKING_DIRECTORY ${TM_DOCKER_ROOT}
  COMMENT "Running developing environment..."
  )

add_custom_target(down
  COMMAND ${TM_SCRIPTS}/down
  WORKING_DIRECTORY ${TM_DOCKER_ROOT}
  COMMENT "Stoping developing environment..."
  )
