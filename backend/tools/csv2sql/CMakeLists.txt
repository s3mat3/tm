cmake_minimum_required(VERSION 3.24)

set(TARGET csv2sql)

project(${TARGET}
  VERSION ${REPOS_VERSION}
  LANGUAGES NONE)

if (BUILD_TOOLS_CSV2SQL)
  set(GOLANG_ENTRY_POINT "${TM_BACKEND_ROOT}/tools/csv2sql/cmd/main.go")
  add_golang_executable(${TARGET})
endif()
