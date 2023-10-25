package main

import (
	"embed"

	"github.com/arun07as/wails-postman/modules/app"
	"github.com/arun07as/wails-postman/modules/collections"
	"github.com/arun07as/wails-postman/modules/menu"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	myApp := app.NewApp()

	collectionManager := &collections.CollectionManager{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-postman",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        myApp.Startup,

		Menu: menu.GenerateMenu(myApp), // reference the menu above

		Bind: []interface{}{
			myApp,
			collectionManager,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
