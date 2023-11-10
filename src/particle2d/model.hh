#pragma once

#include <vector>
#include <Eigen/Eigen>
#include <SDL3/SDL.h>

template <typename S>
struct Particle2 {
    S m;
    Eigen::Vector2<S> x;
    Eigen::Vector2<S> v;
    S r;
    SDL_Color c;
};
using Particle2f = Particle2<float>;

class ParticleModel {
  public:
    float worldWidth = 100;
    float worldHeight = 100;
    std::vector<Particle2f> particles;

    ParticleModel(std::vector<Particle2f>&& particles): particles(particles) {}

    void tickMs(int dtMs) {
        float dtS = dtMs / 1000.0f;
        for (auto& p : particles) {
            p.v = p.v + dtS * g / p.m;
            p.x = p.x + dtS * p.v;
        }
    }

  private:
    static inline const Eigen::Vector2f g {0, -9.8};
};