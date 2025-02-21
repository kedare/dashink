package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kedare/go-dashink/pkg/output"
)

const (
	WIDTH  = 800
	HEIGHT = 480
)

func BuildWindow(app fyne.App) fyne.Window {
	w := app.NewWindow("Dashink")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			output.CaptureWindowToFile(w, "screenshot.png")
			w.Close()
		}),
	))

	return w
}
