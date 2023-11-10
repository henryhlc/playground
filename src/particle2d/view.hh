#pragma once

#include <stdexcept>
#include <memory>
#include <vector>

#include <SDL3/SDL.h>

#include "model.hh"

class ParticleView {
  public:
    ParticleView() {
        if (SDL_CreateWindowAndRenderer(windowWidth, windowHeight, 0, &window, &renderer) < 0) {
            SDL_Log("SDL_CreateWindowAndRenderer failed (%s)", SDL_GetError());
            throw std::runtime_error("SDL_CreateWindowAndRenderer failed!");
        }
        clearAndDrawBackground();
        SDL_RenderPresent(renderer);
    }

    ParticleView(ParticleView&& other) {
        *this = std::move(other);
    }

    void operator=(ParticleView&& other) {
        window = other.window;
        renderer = other.renderer;
        other.window = nullptr;
        other.renderer = nullptr;
    }

    ~ParticleView() {
        SDL_DestroyRenderer(renderer);
        SDL_DestroyWindow(window);
    }

    void clearAndDrawBackground() {
        SDL_SetRenderDrawColor(renderer, 0, 0, 0, SDL_ALPHA_OPAQUE);
        SDL_RenderClear(renderer);
    }

    void drawParticles(const std::vector<Particle2f> ps, float worldWidth, float worldHeight) {
        clearAndDrawBackground();
        float widthScaleFactor = windowWidth / worldWidth;
        float heightScaleFactor = windowHeight / worldHeight; 
        for (const auto& p : ps) {
            SDL_FRect rect {
                .x = (p.x.x() - p.r) * widthScaleFactor,
                .y = windowHeight - (p.x.y() - p.r) * heightScaleFactor,
                .w = 2 * p.r * widthScaleFactor,
                .h = 2 * p.r * heightScaleFactor,
            };
            SDL_SetRenderDrawColor(renderer, p.c.r, p.c.g, p.c.b, SDL_ALPHA_OPAQUE);
            SDL_RenderFillRect(renderer, &rect);
        }
        SDL_RenderPresent(renderer);
    }

  private:
    int windowWidth = 500;
    int windowHeight = 500;

    SDL_Window* window;
    SDL_Renderer* renderer;
};
