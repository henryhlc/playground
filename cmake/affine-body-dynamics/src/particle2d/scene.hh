#pragma once

#include "model.hh"
#include "view.hh"

class ParticleScene {
  public:
    static ParticleScene createExampleScene();

    ParticleScene(ParticleView&& ui, ParticleModel&& model): ui(std::move(ui)), model(std::move(model)) {}

    void loop() {
        ui.drawParticles(model.particles, model.worldWidth, model.worldHeight);
        bool quit = false;
        SDL_Event event;
        while (!quit) {
            while (SDL_PollEvent(&event) != 0) {
                switch(event.type) {
                    case SDL_EVENT_WINDOW_CLOSE_REQUESTED:
                        quit = true;
                        break;
                    case SDL_EVENT_KEY_UP:
                        model.tickMs(100);
                        ui.drawParticles(model.particles, model.worldWidth, model.worldHeight);
                        break;
                }
            }
        }
    }

  private:
    ParticleView ui;
    ParticleModel model;
};