#[[
 This code is licensed under the MIT License, see the LICENSE file for details
 Copyright © 2026 s3mat3
 Author s3mat3

 Brief:
   Utilty functions for cmake
]]

#[[
Utility functions to display variables during debugging.

cmake -S source-dir -B dist-dir -DCMAKE_MESSAGE_LOG_LEVEL=DEBUG <- This is displayed when this command is executed or when --target rebuild_cache is executed.
cmake -S source-dir -B dist-dir --log-level=DEBUG <- This is displayed when this command is executed.

set(STRING_ARRAY qwe;asd;zxc;rty;fgh;vbn)
s3_debug_print("STRING_ARRAY" ${STRING_ARRAY})
=== console ===
-- *** STRING_ARRAY size = 6 ***
-- item> qwe
-- item> asd
-- item> zxc
-- item> rty
-- item> fgh
-- item> vbn
-- *** end of list items ***
=== endconsole ===

set(SCALAR "Hello")
s3_debug_print("SCALAR" ${SCALAR})
=== console ===
-- *** SCALAR size = 1 ***
-- >>> Not a list: Hello
-- *** end of list items ***
=== endconsole ===

s3_debug_print(${SCALAR})
=== console ===
-- *** Hello size = 0 ***
-- >>> Nothin to print
-- *** end of list items ***
=== endconsole ===

s3_debug_print("SCALAR")
=== console ===
-- *** SCALAR size = 0 ***
-- >>> Nothin to print
-- *** end of list items ***
=== endconsole ===

]]
function(s3_debug_print s3_debug_print_arg_msg)
  cmake_parse_arguments(DEBUG_LIST "" "" "" ${ARGN})
  set(print_target_list ${ARGN})
  list(LENGTH print_target_list list_size)
  message(DEBUG "*** ${s3_debug_print_arg_msg} size = ${list_size} ***")
  if (${list_size} LESS 1)
    message(DEBUG ">>> Nothing to print")
  elseif (${list_size} LESS 2)
    message(DEBUG ">>> Not a list: ${print_target_list}")
  else ()
    foreach(item ${print_target_list})
      message(DEBUG "item> ${item}")
    endforeach()
  endif ()
  message(DEBUG "*** end of list items ***\n")
endfunction()

