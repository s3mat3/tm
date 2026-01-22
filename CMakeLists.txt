#[[
usage cmake -DCMAKE_BUILD_TYPE=(Debug | Release | '') \
 -DCMAKE_VERBOSE_MAKEFILE=(on | off) \
 -DBUILD_TOOLS=(on | off default) \
 -DBUILD_APP=(on | off default)
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

set(TM_APP_ROOT ${TM_ROOT_DIRECTORY}/app)
set(TM_TOOLS_ROOT ${TM_ROOT_DIRECTORY}/tools)
set(TM_DOCKER_ROOT ${TM_ROOT_DIRECTORY}/docker)

## output directories
set(TM_OUTPUT_ROOT ${TM_BUILD_DIRECTORY})

###########
# options #
###########
option(BUILD_TOOLS "Build tools" Off)
option(BUILD_APP "Build app" Off)

if(BUILD_TOOLS)
  set(TM_BUILD_TOOLS ON)
  add_subdirectory(${TM_TOOLS_ROOT})
endif()

if(BUILD_APP)
  set(TM_BUILD_APP ON)
  add_subdirectory(${TM_APP_ROOT})
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
