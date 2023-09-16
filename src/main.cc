#include <SDL3/SDL.h>

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
        }
    }

    SDL_DestroyWindow(window);

    return 0;
}