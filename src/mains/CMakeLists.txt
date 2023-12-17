add_executable(polyscope_example polyscope_example.cpp)
target_link_libraries(polyscope_example PRIVATE Eigen3::Eigen igl::core polyscope)

add_executable(sdl_example sdl_example.cpp)
target_link_libraries(sdl_example SDL3::SDL3 particle2d)