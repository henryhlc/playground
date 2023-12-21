add_custom_command(OUTPUT massmatrix_integral.h
	COMMAND python3 "${CMAKE_CURRENT_SOURCE_DIR}/massmatrix_integral.py" > massmatrix_integral.h
	DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/massmatrix_integral.py"
)

add_custom_command(OUTPUT distance.h
	COMMAND python3 "${CMAKE_CURRENT_SOURCE_DIR}/distance.py" > distance.h
	DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/distance.py"
)

add_library(affinebody
	"${CMAKE_CURRENT_BINARY_DIR}/massmatrix_integral.h"
	"${CMAKE_CURRENT_BINARY_DIR}/distance.h"
	affinebody.h
	affinebody.cpp)
set_target_properties(affinebody PROPERTIES LINKER_LANGUAGE CXX)
target_link_libraries(affinebody PUBLIC Eigen3::Eigen ipc::toolkit)
target_include_directories(affinebody PUBLIC "${CMAKE_CURRENT_BINARY_DIR}")