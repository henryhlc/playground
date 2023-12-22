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

    // CCW faces for correct normal direction.
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
         
    std::vector<AffineBody> pp_states {
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.87, 1.98, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                -0.82, 2.01, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.02, 3.35, 0.03, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.02, 2.57, -0.82, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
    };
    std::vector<AffineBody> p_states {
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.8, 2.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                -0.8, 2.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.0, 3.32, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
        AffineBody {simplex_V, simplex_F,
            Eigen::Vector<double,12>{
                0.02, 2.6, -0.77, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0
            }},
    };

    auto simulation_start = std::chrono::steady_clock::now();
    auto frame_period = 10ms;
    size_t num_frames = 120;
    int steps_per_frame = 1;
    auto step_period = frame_period / steps_per_frame;
    std::vector<std::vector<AffineBody>> frame_states {num_frames};
    frame_states[0] = p_states;
    for (int f = 1; f < num_frames; ++f) {
        auto frame_start = std::chrono::steady_clock::now();
        std::cout << "Frame " << f << std::endl;
        for (int s = 0; s < steps_per_frame; ++s) {

            auto step_start = std::chrono::steady_clock::now();
            auto curr_states = affine_body_dynamics(p_states, pp_states, step_period);
            pp_states = p_states;
            p_states = curr_states;
            auto step_end = std::chrono::steady_clock::now();
            auto step_duration = std::chrono::round<std::chrono::milliseconds>(step_end - step_start);
            std::cout << "Step " << s << ": " << step_duration << std::endl << std::endl;

        }
        frame_states[f] = p_states;
        auto frame_end = std::chrono::steady_clock::now();
        auto frame_duration = std::chrono::round<std::chrono::milliseconds>(frame_end - frame_start);
        std::cout << "Frame " << f << ": " << frame_duration << std::endl << std::endl;
    }
    auto simulation_end = std::chrono::steady_clock::now();
    auto simulation_duration = std::chrono::round<std::chrono::milliseconds>(simulation_end - simulation_start);
    std::cout << "Total simulation time: " << simulation_duration << std::endl;

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