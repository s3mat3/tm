#[[
 This code is licensed under the MIT License, see the LICENSE file for details
 Copyright © 2026 s3mat3
 Author s3mat3

 Brief:
   Simple build target for GO language.
]]

#[[
Building an executable from GO source

add_golang_executable(
  target       -- Target name. (required)
  ENTRY_POINT  -- Entry point for executable. (required)
  GOMOD_DIR    -- GO app root. (option) default CMAKE_CURRENT_SOURCE_DIR
  [OUTPUT]     -- Target output destination. (option) default ${CMAKE_CURRENT_BINARY_DIR}/bin/${target}
  [FLAGS]      -- Build flag. (option) (list) default -trimpath
  [LD_FLAGS]   -- Linker flag. (option) (list) defult ""
  [SOURCES]    -- Source files to be built. (option) (list) default file(GLOB SOURCES "${CMAKE_CURRENT_SOURCE_DIR}/*.go")
  [DEPS]       -- Non-source build dependencies (option) (list) default empty
  [VERSION]    -- Project(repository) version. (option) defult no set and no effect
  [REVISION]   -- Project(repository) revision. (option) default no set and no effect
  [GOOS]       -- OS to build on. (option) default from CMAKE_SYSTEM_NAME
  [GOARCH]     -- CPU to build on. (option) defualt from CMAKE_SYSTEM_PROCESSOR
  [CGO]        -- c/c++ link flag. (option) defult 0 (Not link)
)

For example
set(TARGET "myapp")
set(TARGET_ENTRY_POINT "${CMAKE_CURRENT_SOURCE_DIR}/cmd/some-dir/main.go")
list(APPEND TARGET_BUILD_FLAGS -tags release)
list(APPEND TARGET_BUILD_LD_FLAGS -s -w)
add_golang_executable(
  ${TARGET}
  ENTRY_POINT ${TARGET_ENTRY_POINT}
  FLAGS ${TARGET_BUILD_FLAGS}
  LD_FLAGS ${TARGET_BUILD_LD_FLAGS}
  VERSION ${REPOS_VERSION}
  REVISION ${REPOS_REVISION}
)
]]

function(add_golang_executable target)
  # find golang
  find_program(GOLANG_EXECUTABLE go)
  if (NOT GOLANG_EXECUTABLE)
    message(FATAL_ERROR "add_golang_executable: Is golang installed? Check your executable path.")
  endif (NOT GOLANG_EXECUTABLE)
  message("Founded GO lang in : ${GOLANG_EXECUTABLE}")
  #
  cmake_parse_arguments(
    GO_BUILD_ARG
    ""
    "ENTRY_POINT"
    "GOMOD_DIR;OUTPUT;FLAGS;LD_FLAGS;SOURCES;DEPS;VERSION;REVISION;GOOS;GOARCH;CGO"
    ${ARGN})
  # Check entry point. if no given to fatal error.
  if (NOT DEFINED GO_BUILD_ARG_ENTRY_POINT)
    message(FATAL_ERROR "add_golang_executable: No entry point...")
  endif(NOT DEFINED GO_BUILD_ARG_ENTRY_POINT)
  #
  message(STATUS "add_golang_executable: Build ${target} from ${GO_BUILD_ARG_ENTRY_POINT}")
  #
  if (NOT DEFINED GO_BUILD_ARG_GOMOD_DIR)
    set(GO_BUILD_ARG_GOMOD_DIR ${CMAKE_CURRENT_SOURCE_DIR})
    message(STATUS "add_golang_executable: Set default GOMOD_DIR => ${GO_BUILD_ARG_GOMOD_DIR}")
  endif()
  #
  if (NOT DEFINED GO_BUILD_ARG_OUTPUT)
    set(GO_BUILD_ARG_OUTPUT "${CMAKE_CURRENT_BINARY_DIR}/bin/${target}")
    message(STATUS "add_golang_executable: Set default OUTPUT => ${GO_BUILD_ARG_OUTPUT}")
  endif ()
  # Setup default sources (Recommend run rebuild_cache)
  if (NOT DEFINED GO_BUILD_ARG_SOURCES)
    file(GLOB_RECURSE
      GO_BUILD_ARG_SOURCES
      CONFIGURE_DEPENDS
      "${CMAKE_CURRENT_SOURCE_DIR}/*.go"
    )
    message(STATUS "add_golang_executable: Set default SOURCES => ${GO_BUILD_ARG_SOURCES}")
  endif()

  list(APPEND GO_BUILD_ARG_FLAGS ${GO_BUILD_ARG_FLAGS} -trimpath)
  list(REMOVE_DUPLICATES GO_BUILD_ARG_FLAGS)
  # What to do if a version is set
  if (DEFINED GO_BUILD_ARG_VERSION)
    list(APPEND GO_BUILD_ARG_LD_FLAGS -X main.Version=${GO_BUILD_ARG_VERSION})
  endif()
  # What to do if a revision is set
  if (DEFINED GO_BUILD_ARG_REVISION)
    list(APPEND GO_BUILD_ARG_LD_FLAGS -X main.Revision=${GO_BUILD_ARG_REVISION})
  endif()
  # Setup default GOOS
  # Cross-compiling is probably not possible
  if (NOT DEFINED GO_BUILD_ARG_GOOS)
    string(TOLOWER ${CMAKE_SYSTEM_NAME} GO_BUILD_ARG_GOOS)
    message(STATUS "add_golang_executable: Set default GOOS => ${GO_BUILD_ARG_GOOS}")
  endif ()
  # Setup default GOARCH
  if (NOT GO_BUILD_ARG_GOARCH)
    if (${CMAKE_SYSTEM_PROCESSOR} STREQUAL "i686")
      set(GO_BUILD_ARG_GOARCH "386")
    elseif(${CMAKE_SYSTEM_PROCESSOR} STREQUAL "aarch64" OR ${CMAKE_SYSTEM_PROCESSOR} STREQUAL "arm64")
      set(GO_BUILD_ARG_GOARCH "arm64")
    elseif(${CMAKE_SYSTEM_PROCESSOR} STREQUAL "arm")
      set(GO_BUILD_ARG_GOARCH "arm")
    else () # AMD64(windows) OR x86_64
      set(GO_BUILD_ARG_GOARCH "amd64")
    endif ()
    message(STATUS "add_golang_executable: Set default GOARCH => ${GO_BUILD_ARG_GOARCH}")
  endif ()
  # Setup default CGO
  if (NOT GO_BUILD_ARG_CGO)
    set(GO_BUILD_ARG_CGO "0")
    message(STATUS "add_golang_executable: Set default CGO = ${GO_BUILD_ARG_CGO}")
  endif ()
  #
  message(VERBOSE "TARGET      = ${target}")
  message(VERBOSE "ENTRY_POINT = ${GO_BUILD_ARG_ENTRY_POINT}")
  message(VERBOSE "GOMOD_DIR   = ${GO_BUILD_ARG_GOMOD_DIR}")
  message(VERBOSE "OUTPUT      = ${GO_BUILD_ARG_OUTPUT}")
  message(VERBOSE "FLAGS       = ${GO_BUILD_ARG_FLAGS}")
  message(VERBOSE "LD_FLAGS    = ${GO_BUILD_ARG_LD_FLAGS}")
  message(VERBOSE "SOURCES     = ${GO_BUILD_ARG_SOURCES}")
  message(VERBOSE "VERSION     = ${GO_BUILD_ARG_VERSION}")
  message(VERBOSE "REVISION    = ${GO_BUILD_ARG_REVISION}")
  message(VERBOSE "GOOS        = ${GO_BUILD_ARG_GOOS}")
  message(VERBOSE "GOARCH      = ${GO_BUILD_ARG_GOARCH}")
  message(VERBOSE "CGO         = ${GO_BUILD_ARG_CGO}")
  # Build custom command
  add_custom_command(
    OUTPUT ${GO_BUILD_ARG_OUTPUT}
    COMMAND CGO=${GO_BUILD_ARG_CGO} GOOS=${GO_BUILD_ARG_GOOS} GOARCH=${GO_BUILD_ARG_GOARCH}
    go build
    -o ${GO_BUILD_ARG_OUTPUT}
    ${GO_BUILD_ARG_FLAGS}
    -ldflags="${GO_BUILD_ARG_LD_FLAGS}"
    ${GO_BUILD_ARG_ENTRY_POINT}
    WORKING_DIRECTORY ${GO_BUILD_ARG_GOMOD_DIR}
    DEPENDS ${GO_BUILD_ARG_SOURCES} ${GO_BUILD_ARG_DEPS}
    COMMENT "Building ..."
  )
  # Build custom target
  add_custom_target(${target} ALL
    DEPENDS ${GO_BUILD_ARG_OUTPUT}
    COMMENT "Done build proccess ..."
  )
endfunction(add_golang_executable)
