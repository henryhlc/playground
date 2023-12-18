add_executable(polyscope_example polyscope_example.cpp)
target_link_libraries(polyscope_example PRIVATE Eigen3::Eigen igl::core polyscope)

add_executable(sdl_example sdl_example.cpp)
target_link_libraries(sdl_example PRIVATE SDL3::SDL3 particle2d)

add_executable(affinebody_example affinebody_example.cpp)
target_link_libraries(affinebody_example PRIVATE Eigen3::Eigen igl::core polyscope affinebody)
target_include_directories(affinebody_example PRIVATE "${CMAKE_CURRENT_BINARY_DIR}/..")