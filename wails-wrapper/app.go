package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) onClose(ctx context.Context) (prevent bool) {
	result, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Message: "Do you want to quit?",
		Type:    runtime.QuestionDialog,
	})

	if err != nil {
		log.Println(err)
	}

	prevent = result == "No"

	return
}

func (a *App) onStartup(ctx context.Context) {
	runtime.EventsOn(ctx, "resize", func(_ ...any) {
		isFullscreen := runtime.WindowIsFullscreen(ctx)

		if !isFullscreen {
			log.Println("fullscreen only.")

			runtime.WindowFullscreen(ctx)
		}
	})
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
