package main

import (
	"image/png"
	_ "image/png"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  = 800
	HEIGHT = 480
)

// captureWindowToFile captures the current content of a Fyne window's canvas
// and saves it to a file with the specified fileName.
// It creates a new RGBA image from the captured content and encodes it as a PNG.
func captureWindowToFile(w fyne.Window, fileName string) {
	// Capture the current window's canvas content
	img := w.Canvas().Capture()

	// Create a file to save the image
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Encode the image as a PNG and write it to the file
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Screenshot saved to", fileName)
}

func main() {
	a := app.New()
	w := a.NewWindow("Dashink")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			captureWindowToFile(w, "screenshot.png")
			w.Close()
		}),
	))

	w.ShowAndRun()
}
