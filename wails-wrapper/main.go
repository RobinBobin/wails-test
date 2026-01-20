package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []any{
			app,
		},
		Frameless:        true,
		OnBeforeClose:    app.onClose,
		OnStartup:        app.onStartup,
		Title:            "wails-wrapper",
		WindowStartState: options.Fullscreen,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
