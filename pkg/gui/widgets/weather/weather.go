package weather

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	log "github.com/sirupsen/logrus"

	"github.com/kedare/dashink/pkg/services/weather"
)

// WeatherWidget represents a widget displaying weather information
type WeatherWidget struct {
	fyne.Widget
	lat    float64
	lon    float64
	apiKey string
}

func NewWeatherWidget(lat, lon float64) *WeatherWidget {
	return &WeatherWidget{
		lat: lat,
		lon: lon,
	}
}

func (w *WeatherWidget) CreateContent() fyne.CanvasObject {
	// Get current weather data
	currentWeather, err := weather.GetCurrentWeather(w.lat, w.lon)
	if err != nil {
		// Show error state
		return container.NewCenter(widget.NewLabel("Weather data unavailable"))
	}

	// Get weather icon
	iconImage, err := weather.IconImage(currentWeather.Weather[0].Icon)
	var weatherIcon *canvas.Image
	if err != nil {
		weatherIcon = canvas.NewImageFromResource(theme.WarningIcon())
		log.WithError(err).Errorln("Error getting weather icon")
	} else {
		weatherIcon = canvas.NewImageFromImage(iconImage)
	}
	weatherIcon.SetMinSize(fyne.NewSize(50, 50))
	weatherIcon.FillMode = canvas.ImageFillOriginal
	weatherIcon.Resize(fyne.NewSize(50, 50))

	// Create temperature labels
	currentTemp := canvas.NewText(fmt.Sprintf("%.1f°C", currentWeather.Main.Temp), color.White)
	currentTemp.TextSize = 24

	minMaxTemp := canvas.NewText(
		fmt.Sprintf("%.1f°C / %.1f°C",
			currentWeather.Main.TempMin,
			currentWeather.Main.TempMax,
		),
		color.White,
	)
	minMaxTemp.TextSize = 16

	// Create description label
	description := canvas.NewText(currentWeather.Weather[0].Description, color.White)
	description.TextSize = 14

	// Layout the widgets
	tempContainer := container.NewVBox(
		currentTemp,
		minMaxTemp,
		description,
	)

	return container.NewHBox(
		weatherIcon,
		tempContainer,
	)
}
