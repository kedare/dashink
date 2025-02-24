package output

import (
	"image/png"
	"log"
	"os"

	"fyne.io/fyne/v2"
)

// CaptureCanvasToFile captures the current content of a Fyne window's canvas
// and saves it to a file with the specified fileName.
// It creates a new RGBA image from the captured content and encodes it as a PNG.
func CaptureCanvasToFile(c fyne.Canvas, fileName string) {
	// Capture the current window's canvas content
	img := c.Capture()

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
