add_custom_command(OUTPUT massmatrix_integral.cpp
	COMMAND python3 "${CMAKE_CURRENT_SOURCE_DIR}/massmatrix_integral.py" > massmatrix_integral.cpp
	DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/massmatrix_integral.py"
)

add_custom_command(OUTPUT distance.cpp
	COMMAND python3 "${CMAKE_CURRENT_SOURCE_DIR}/distance.py" > distance.cpp
	DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/distance.py"
)

add_library(affinebody_distance
	distance.h
	"${CMAKE_CURRENT_BINARY_DIR}/distance.cpp"
)
target_link_libraries(affinebody_distance PUBLIC Eigen3::Eigen ipc::toolkit)

add_library(affinebody
	massmatrix_integral.h "${CMAKE_CURRENT_BINARY_DIR}/massmatrix_integral.cpp"
	affinebody.h affinebody.cpp
	contact_ip.h contact_ip.cpp
)
set_target_properties(affinebody PROPERTIES LINKER_LANGUAGE CXX)
target_link_libraries(affinebody PUBLIC affinebody_distance Eigen3::Eigen ipc::toolkit)