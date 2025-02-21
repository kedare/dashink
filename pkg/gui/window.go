package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	WIDTH  = 800
	HEIGHT = 480
)

func BuildWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Dashink")
	window.Resize(fyne.NewSize(WIDTH, HEIGHT))
	window.SetFixedSize(true)

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(top, nil, left, nil, middle)
	window.SetContent(content)

	return window
}
