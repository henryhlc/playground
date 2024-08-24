#include "particle2d.hh"

#include "scene.hh"

void startParticle2dScene() {
    ParticleScene(ParticleView {}, ParticleModel { std::vector<Particle2f> {
        Particle2f {
            .m = 1.0f,
            .x = Eigen::Vector2f {0.0, 0.0},
            .v = 0.5 * Eigen::Vector2f {50.0, 80.0},
            .r = 1.0f,
            .c = SDL_Color {
                .r = 255,
                .g = 0,
                .b = 0,
            },
        }
        , Particle2f {
            .m = 1.0f,
            .x = Eigen::Vector2f {0.0, 0.0},
            .v = 0.4 * Eigen::Vector2f {50.0, 80.0},
            .r = 1.0f,
            .c = SDL_Color {
                .r = 0,
                .g = 255,
                .b = 0,
            },
        }
        , Particle2f {
            .m = 1.0f,
            .x = Eigen::Vector2f {0.0, 0.0},
            .v = 0.3 * Eigen::Vector2f {50.0, 80.0},
            .r = 1.0f,
            .c = SDL_Color {
                .r = 0,
                .g = 0,
                .b = 255,
            },
        }
        , Particle2f {
            .m = 1.0f,
            .x = Eigen::Vector2f {50.0, 80.0},
            .v = Eigen::Vector2f {0.0, 0.0},
            .r = 1.0f,
            .c = SDL_Color {
                .r = 255,
                .g = 255,
                .b = 255,
            },
        }}
    }).loop();
}