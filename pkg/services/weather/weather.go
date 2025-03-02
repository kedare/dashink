package weather

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/briandowns/openweathermap"
	log "github.com/sirupsen/logrus"
)

var (
	apiKey = os.Getenv("OPENWEATHERMAP_API_KEY")
)

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func IconURL(iconCode string) string {
	return fmt.Sprintf("http://openweathermap.org/img/wn/%s@2x.png", iconCode)
}

func IconImage(iconCode string) (image.Image, error) {
	cacheDir := "cache"
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, 0755)
	}
	fileName := filepath.Join(cacheDir, iconCode+".png")
	log.WithField("fileName", fileName).Debugln("Getting weather icon")

	// If the file exists, return it
	if data, err := os.ReadFile(fileName); err == nil {
		log.WithField("fileName", fileName).Debugln("Using cached weather icon")
		image, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		return image, nil
	}

	// Get the icon URL
	url := IconURL(iconCode)

	// Download the icon
	log.WithField("url", url).Debugln("Downloading weather icon")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorln("Unexpected status code")
		return nil, errors.New("unexpected status code")
	}

	// Read the icon data
	iconData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Save the icon
	log.WithField("fileName", fileName).Debugln("Saving weather icon")
	if err := os.WriteFile(fileName, iconData, 0644); err != nil {
		return nil, err
	}

	image, err := png.Decode(bytes.NewReader(iconData))
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetCurrentWeather retrieves the current weather data including max and min temperatures
// from OpenWeatherMap API for the specified location using the briandowns/openweathermap library
func GetCurrentWeather(lat, lon float64) (*openweathermap.CurrentWeatherData, error) {
	w, err := openweathermap.NewCurrent("C", "en", apiKey)
	if err != nil {
		log.WithError(err).Errorln("Failed to create OpenWeatherMap client")
		return nil, err
	}

	log.WithFields(log.Fields{
		"lat": lat,
		"lon": lon,
	}).Debugln("Fetching current weather data")

	err = w.CurrentByCoordinates(&openweathermap.Coordinates{
		Latitude:  lat,
		Longitude: lon,
	})

	if err != nil {
		log.WithError(err).Errorln("Failed to fetch weather data")
		return nil, err
	}

	log.WithFields(log.Fields{
		"temp":        w.Main.Temp,
		"temp_min":    w.Main.TempMin,
		"temp_max":    w.Main.TempMax,
		"description": w.Weather[0].Description,
	}).Debugln("Weather data retrieved successfully")

	return w, nil
}

// GetCurrentAQI retrieves the current Air Quality Index (AQI) data
// from OpenWeatherMap API for the specified location using the briandowns/openweathermap library
func GetCurrentAQI(lat, lon float64) (*openweathermap.PollutionData, error) {
	a, err := openweathermap.NewPollution(apiKey)
	if err != nil {
		log.WithError(err).Errorln("Failed to create OpenWeatherMap AQI client")
		return nil, err
	}

	log.WithFields(log.Fields{
		"lat": lat,
		"lon": lon,
	}).Debugln("Fetching current AQI data")

	params := openweathermap.PollutionParameters{
		Location: openweathermap.Coordinates{
			Latitude:  lat,
			Longitude: lon,
		},
	}

	err = a.PollutionByParams(&params)

	if err != nil {
		log.WithError(err).Errorln("Failed to fetch AQI data")
		return nil, err
	}

	aqi := a.List[0].Main.Aqi

	log.WithFields(log.Fields{
		"aqi":        aqi,
		"components": a.List[0].Components,
		"time":       a.List[0].Dt,
	}).Debugln("AQI data retrieved successfully")

	return &a.List[0], nil
}
