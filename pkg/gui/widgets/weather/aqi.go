package weather

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/kedare/dashink/pkg/services/weather"
)

type AQIWidget struct {
	lat, lon float64
}

func NewAQIWidget(lat, lon float64) *AQIWidget {
	return &AQIWidget{
		lat: lat,
		lon: lon,
	}
}

func (w *AQIWidget) CreateContent() fyne.CanvasObject {
	// Get current AQI data
	aqi, err := weather.GetCurrentAQI(w.lat, w.lon)
	if err != nil {
		// Show error state
		return container.NewCenter(widget.NewLabel("AQI data unavailable"))
	}

	// Create AQI value label
	aqiValue := canvas.NewText(fmt.Sprintf("AQI: %v", aqi.Main.Aqi), getAQIColor(aqi.Main.Aqi))
	aqiValue.TextSize = 24

	// Create AQI description label
	aqiDesc := canvas.NewText(getAQIDescription(aqi.Main.Aqi), color.White)
	aqiDesc.TextSize = 14

	// Layout the widgets
	container := container.NewVBox(
		aqiValue,
		aqiDesc,
	)

	return container
}

func getAQIColor(value float64) color.Color {
	switch {
	case value <= 50:
		return color.RGBA{0, 228, 0, 255} // Green
	case value <= 100:
		return color.RGBA{255, 255, 0, 255} // Yellow
	case value <= 150:
		return color.RGBA{255, 126, 0, 255} // Orange
	case value <= 200:
		return color.RGBA{255, 0, 0, 255} // Red
	case value <= 300:
		return color.RGBA{143, 63, 151, 255} // Purple
	default:
		return color.RGBA{126, 0, 35, 255} // Maroon
	}
}

func getAQIDescription(value float64) string {
	switch {
	case value <= 50:
		return "Good"
	case value <= 100:
		return "Moderate"
	case value <= 150:
		return "Unhealthy for Sensitive Groups"
	case value <= 200:
		return "Unhealthy"
	case value <= 300:
		return "Very Unhealthy"
	default:
		return "Hazardous"
	}
}
