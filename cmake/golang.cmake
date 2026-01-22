#[[
add_golang_executable

 Variables set externally and used inside the function

 REPOS_VERSION              {string} - Project repository version from VCS
 REPOS_REVISION             {string} - Project repository revision from VCS
 GOLANG_GOMOD_DIRECTORY     {string} - Go module root directory (compile on this directory)
 GOLANG_ENTRY_POINT         {string} - Entry point of target application
 GOLANG_BUILD_FLAGS         {list}   - Go build option(s)
 GOLANG_BUILD_LDFLAGS       {list}   - Go linker options
 GOLANG_BUILD_MODE          {string} - Go build mode (Not yet)
 GOLANG_BUILD_TARGET        {string} - Target name for application
 GOLANG_BUILD_OUT_DIRECTORY {string} - Go builder output binary dhirectory

USAGE:
set(REPOS_VERSION "0.0.0")
set(REPOS_REVISION "deadbeef")
list(APPEND GOLANG_BUILD_FLAGS -tags release)
list(APPEND GOLANG_BUILD_LDFLAGS -s -w)
set(GOLANG_GOMOD_DIRECTORY ${BACKEND_ROOT})
set(GOLANG_ENTRY_POINT "${BACKEND_ROOT}/tools/csv2sql/cmd/main.go")
set(GOLANG_BUILD_OUT_DIRECTORY ${BACKEND_OUTPUT})

add_golang_executable(${TARGET})
]]
function(add_golang_executable GOLANG_BUILD_TARGET)
  list(REMOVE_DUPLICATES GOLANG_BUILD_FLAGS)
  list(APPEND l_flags ${GOLANG_BUILD_FLAGS} -a -trimpath)

  list(REMOVE_DUPLICATES GOLANG_BUILD_LD_FLAGS)
  list(APPEND l_ldflags ${GOLANG_BUILD_LDFLAGS} -X main.Version=${REPOS_VERSION} -X main.Revision=${REPOS_REVISION})
  message("${REPOS_VERSION}/${REPOS_REVISION}")
  # build command
  add_custom_command(OUTPUT ${GOLANG_BUILD_OUT_DIRECTRY}/${GOLANG_BUILD_TARGET}
    COMMAND go build
    -o "${GOLANG_BUILD_OUT_DIRECTORY}/${GOLANG_BUILD_TARGET}"
    ${l_flags}
    -ldflags="${l_ldflags}"
    ${GOLANG_ENTRY_POINT}
    #VERBATIM
    WORKING_DIRECTORY ${GOLANG_GOMOD_DIRECTORY}
    COMMENT "Bulding ${GOLANG_BUILD_TARGET} from ${GOLANG_ENTRY_POINT}"
  )
  # depend checker
  add_custom_target(${GOLANG_BUILD_TARGET} ALL
    SOURCES ${GOLANG_BUILD_OUT_DIRECTRY}/${GOLANG_BUILD_TARGET}
    DEPENDS ${GOLANG_BUILD_OUT_DIRECTRY}/${GOLANG_BUILD_TARGET}
    COMMENT "Build command done...."
  )
endfunction(add_golang_executable)
