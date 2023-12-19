#include <chrono>
#include <iostream>
#include <map>
#include <ranges>
#include <string>

#include <Eigen/Eigen>
#include <polyscope/polyscope.h>
#include <polyscope/surface_mesh.h>

#include "../affinebody/affinebody.h"

using namespace std::literals;

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

    // Floor
    Eigen::Matrix<double,4,3> floor_V;
    floor_V << 0.0, 0.0, 0.0,
               1.0, 0.0, 0.0,
               0.0, 0.0, 1.0,
               1.0, 0.0, 1.0;
    Eigen::Matrix<int,2,3> floor_F;
    floor_F << 0, 1, 3,
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
         
    std::vector<std::vector<AffineBody>> frame_states {
        {
            AffineBody {simplex_V, simplex_F,
                Eigen::Vector<double,12>{
                    0.5, 0.5, 0.5, 2, 0, 0, 0, 2, 0, 0, 0, 2
                }},
            AffineBody {simplex_V, simplex_F,
                Eigen::Vector<double,12>{
                    0.0, 0.0, 0.0, 1, 0, 0, 0, 1, 0, 0, 0, 1
                }}
        }
    };

    auto frame_period = 30ms;
    int num_frames = 360;
    frame_states.resize(num_frames);
    for (int f = 1; f < num_frames; ++f) {
        frame_states[f] = affine_body_dynamics(frame_states[f-1], frame_period);
    }

    polyscope::init();

    auto init_state = frame_states[0];
    for (int i = 0; i < init_state.size(); ++i) {
        auto& ab = init_state[i];
        polyscope::registerSurfaceMesh(std::to_string(i), ab.V(), ab.F);
    }
    polyscope::registerSurfaceMesh("floor", floor_V, floor_F);

    int next_frame = 0;
    auto last_update = std::chrono::system_clock::now();
    polyscope::state::userCallback = [&last_update,&next_frame,&frame_states,&num_frames,&frame_period]() {
        auto now = std::chrono::system_clock::now();
        if (now - last_update < frame_period) {
            return;
        }

        auto& frame_state = frame_states[next_frame];
        for (int i = 0; i < frame_state.size(); ++i) {
            auto& ab = frame_state[i];
            polyscope::getSurfaceMesh(std::to_string(i))->updateVertexPositions(ab.V());
        }
        next_frame = (next_frame + 1) % num_frames;
    };

    polyscope::show();

    return 0;
}