cmake_minimum_required(VERSION 3.7..3.27)

if (${CMAKE_VERSION} VERSION_LESS 3.12)
    cmake_policy(VERSION ${CMAKE_MAJOR_VERSION}.${CMAKE_MINOR_VERSION})
endif()

project(playground VERSION 0.1
    DESCRIPTION "Playground to explore 3D system implementations."
    LANGUAGES C CXX)
set (CMAKE_CXX_STANDARD 23)

include(FetchContent)
FetchContent_Declare(
  SDL
  GIT_REPOSITORY https://github.com/libsdl-org/SDL.git
  GIT_TAG main
  GIT_SHALLOW TRUE
  GIT_PROGRESS TRUE
)
FetchContent_Declare(
  Eigen
  GIT_REPOSITORY https://gitlab.com/libeigen/eigen.git
  GIT_TAG master
  GIT_SHALLOW TRUE
  GIT_PROGRESS TRUE
)

FetchContent_MakeAvailable(SDL Eigen)

add_executable(main src/main.cc)
target_link_libraries(main PRIVATE SDL3::SDL3-static Eigen3::Eigen)