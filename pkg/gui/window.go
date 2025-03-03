package gui

import (
  "time"
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

  dateString := time.Now().Format("2006-01-02T15:04:05-07:00")
	top := canvas.NewText(dateString, color.White)
	weatherWidget := weather.NewWeatherWidget(41.4006716, 2.1832604).CreateContent()
	aqiWidget := weather.NewAQIWidget(41.4006716, 2.1832604).CreateContent()

	// Create a grid layout for the dashboard
	grid := container.NewGridWithColumns(2,
		container.NewPadded(weatherWidget),
		container.NewPadded(aqiWidget),
	)

	content := container.NewBorder(top, nil, nil, nil, grid)
	c.SetContent(content)
	c.Resize(fyne.NewSize(WIDTH, HEIGHT))

	return c
}
