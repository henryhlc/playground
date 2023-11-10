add_subdirectory(particle2d)

add_executable(main main.cc)
target_link_libraries(main PRIVATE SDL3::SDL3-static Eigen3::Eigen igl::core particle2d)