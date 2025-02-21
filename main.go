package main

import (
	_ "image/png"
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/kedare/go-dashink/pkg/gui"
	"github.com/kedare/go-dashink/pkg/output"
)

func main() {
	log.Println("Starting")
	app := app.New()
	window := gui.BuildWindow(app)
	log.Println("Showing GUI")
	window.ShowAndRun()
	output.CaptureWindowToFile(window, "screenshot.png")
}
