package main

import (
	_ "image/png"
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/kedare/dashink/pkg/gui"
	"github.com/kedare/dashink/pkg/output"
)

func main() {
	log.Println("Starting")
	app := app.New()
	canvas := gui.BuildCanvas(app)
	output.CaptureCanvasToFile(canvas, "screenshot.png")
	os.Exit(0)
}
