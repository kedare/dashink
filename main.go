package main

import (
	_ "image/png"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/kedare/go-dashink/pkg/gui"
	"github.com/kedare/go-dashink/pkg/output"
)

func main() {
	log.Println("Starting")
	app := app.New()
	window := gui.BuildWindow(app)
	log.Println("Showing GUI")
	go func() {
		log.Println("Starting screenshot goroutine")
		time.Sleep(1 * time.Second)
		output.CaptureWindowToFile(window, "screenshot.png")
		os.Exit(0)
	}()
	window.ShowAndRun()
}
