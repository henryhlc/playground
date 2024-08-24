#include <iostream>
#include <stdexcept>
#include <chrono>

#include <Eigen/Eigen>
#include <igl/readOBJ.h>

#include <polyscope/polyscope.h>
#include <polyscope/surface_mesh.h>

using namespace std::literals;

int main(int argc, char* argv[]) {
    polyscope::init();

    Eigen::MatrixXd V;
    Eigen::MatrixXi F;

    igl::readOBJ("../data/stanford-bunny.obj", V, F);

    polyscope::registerSurfaceMesh("input mesh", V, F);

    auto lastUpdate = std::chrono::system_clock::now();

    polyscope::state::userCallback = [&lastUpdate, &V]() {
        auto now = std::chrono::system_clock::now();
        if (now - lastUpdate > 1s) {
            lastUpdate = now;
            for (auto r : V.rowwise()) {
                r += Eigen::Vector3d {{0.03, 0.03, 0.03}};
            }
            polyscope::getSurfaceMesh("input mesh")->updateVertexPositions(V);
        }
    };

    polyscope::show();

    return 0;
}