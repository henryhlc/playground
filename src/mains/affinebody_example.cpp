#include <iostream>
#include <map>

#include <Eigen/Eigen>
#include <polyscope/polyscope.h>
#include <polyscope/surface_mesh.h>

#include "../affinebody/affinebody.h"

int main(int argc, char** argv) {

    // Simplex
    Eigen::Matrix<double,4,3> simplex_V;
    simplex_V << 0.0, 0.0, 0.0,
                 1.0, 0.0, 0.0,
                 0.0, 1.0, 0.0,
                 0.0, 0.0, 1.0;
    Eigen::Matrix<int,4,3> simplex_F;
    simplex_F << 0, 2, 1,
                 0, 1, 3,
                 1, 2, 3,
                 0, 3, 2;

    // Box
    Eigen::Matrix<double,8,3> box_V;
    box_V << 0.0, 0.0, 0.0,
             0.5, 0.0, 0.0,
             0.0, 0.1, 0.0,
             0.0, 0.0, 1.0,
             0.0, 0.1, 1.0,
             0.5, 0.0, 1.0,
             0.5, 0.1, 0.0,
             0.5, 0.1, 1.0;
    Eigen::Matrix<int,12,3> box_F;
    box_F << 0, 2, 1,
             1, 2, 6,
             0, 1, 5,
             0, 5, 3,
             0, 3, 4,
             0, 4, 2,
             3, 5, 4,
             5, 7, 4,
             1, 6, 7,
             7, 5, 1,
             2, 4, 7,
             2, 7, 6;
         
    Eigen::Vector<double,12> q1;
    q1 << 0.5, 0.5, 0.5, 2, 0, 0, 0, 2, 0, 0, 0, 2;
    AffineBody simplex1_ab {simplex_V, simplex_F, q1};

    Eigen::Vector<double,12> q2;
    q2 << 0.0, 0.0, 1.0, 0.5, 0, 0, 0, 0.5, 0, 0, 0, 0.5;
    AffineBody simplex2_ab {simplex_V, simplex_F, q2};

    polyscope::init();
    polyscope::registerSurfaceMesh("simplex1", simplex1_ab.V(), simplex1_ab.F);
    polyscope::registerSurfaceMesh("simplex2", simplex2_ab.V(), simplex2_ab.F);
    polyscope::show();

    return 0;
}