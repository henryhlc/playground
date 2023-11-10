add_library(particle2d
    particle2d.hh particle2d.cc
    model.hh model.cc
    view.hh view.cc
    scene.hh scene.cc)
target_link_libraries(particle2d PRIVATE SDL3::SDL3-static Eigen3::Eigen)