package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/tools/playground"

	"github.com/kedare/dashink/pkg/gui/widgets/weather"
)

const (
	WIDTH  = 800
	HEIGHT = 480
)

func BuildCanvas(app fyne.App) fyne.Canvas {
	c := playground.NewSoftwareCanvas()
	c.SetPadded(false)

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

	top := canvas.NewText("top bar", color.White)
	weatherWidget := weather.NewWeatherWidget(41.4006716, 2.1832604).CreateContent()
	content := container.NewBorder(top, nil, nil, nil, container.NewCenter(weatherWidget))
	c.SetContent(content)
	c.Resize(fyne.NewSize(WIDTH, HEIGHT))

	return c
}
