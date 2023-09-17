#include <iostream>
#include <stdexcept>
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

class ParticleWorld {
  public:
    float worldWidth = 100;
    float worldHeight = 100;
    std::vector<Particle2f> particles;

    ParticleWorld(std::vector<Particle2f>&& particles): particles(particles) {}

    void tickMs(int dtMs) {
        float dtS = dtMs / 1000.0f;
        for (auto& p : particles) {
            p.v = p.v + dtS * g / p.m;
            p.x = p.x + dtS * p.v;
        }
    }

  private:
    const Eigen::Vector2f g {0, -9.8};
};

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

class ParticleApp {
  public:
    static ParticleApp createExampleApp() {
        return ParticleApp(ParticleView {}, ParticleWorld { std::vector<Particle2f> {
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
        });
    }

    ParticleApp(ParticleView&& ui, ParticleWorld&& world): ui(std::move(ui)), world(std::move(world)) {}

    void loop() {
        bool quit = false;
        SDL_Event event;
        while (!quit) {
            while (SDL_PollEvent(&event) != 0) {
                switch(event.type) {
                    case SDL_EVENT_WINDOW_CLOSE_REQUESTED:
                        quit = true;
                        break;
                    case SDL_EVENT_KEY_UP:
                        world.tickMs(100);
                        ui.drawParticles(world.particles, world.worldWidth, world.worldHeight);
                        break;
                }
            }
        }
    }

  private:
    ParticleView ui;
    ParticleWorld world;
};

int main(int argc, char* argv[]) {
    if (SDL_Init(SDL_INIT_VIDEO | SDL_INIT_EVENTS) < 0) {
        SDL_Log("SDL_Init failed (%s)", SDL_GetError());
        return 1;
    }

    SDL_Window* window = SDL_CreateWindow("Index", 300, 300, 0);
    if (window == nullptr) {
        SDL_Log("SDL_CreateWindow failed (%s)", SDL_GetError());
        return 1;
    }

    SDL_Event e;
    bool quit = false;
    while (!quit) {
        while (SDL_PollEvent(&e) != 0) {
            if (e.type == SDL_EVENT_QUIT) {
                quit = true;
            }
            if (e.type == SDL_EVENT_KEY_UP) {
                switch (e.key.keysym.sym) {
                    case SDLK_1:
                        ParticleApp::createExampleApp().loop();
                }
            }
        }
    }

    SDL_DestroyWindow(window);

    return 0;
}