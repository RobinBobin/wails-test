package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct{}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here

	runtime.EventsOn(ctx, "resize", func(_ ...any) {
		isFullscreen := runtime.WindowIsFullscreen(ctx)

		if !isFullscreen {
			log.Println("fullscreen only.")

			runtime.WindowFullscreen(ctx)
		}
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
