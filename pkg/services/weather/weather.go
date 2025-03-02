package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var (
	apiKey    = os.Getenv("OPENWEATHERMAP_API_KEY")
	latitude  = getEnvWithDefault("WEATHER_LATITUDE", "41.4006716")
	longitude = getEnvWithDefault("WEATHER_LONGITUDE", "2.1832604")

	urlCurrentWeather = "https://api.openweathermap.org/data/2.5/weather"
	urlHourlyWeather  = "https://api.openweathermap.org/data/2.5/forecast"
	urlCurrentAQI     = "http://api.openweathermap.org/data/2.5/air_pollution"
)

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func iconURL(iconCode string) string {
	return fmt.Sprintf("http://openweathermap.org/img/wn/%s@2x.png", iconCode)
}

func iconImage(iconCode string) ([]byte, error) {
	fileName := filepath.Join("cache", iconCode+".png")
	log.WithField("fileName", fileName).Debugln("Getting weather icon")

	// If the file exists, return it
	if data, err := os.ReadFile(fileName); err == nil {
		log.WithField("fileName", fileName).Debugln("Using cached weather icon")
		return data, nil
	}

	// Get the icon URL
	url := iconURL(iconCode)

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

	return iconData, nil
}

func currentWeather() (map[string]interface{}, error) {
	log.Debugln("Fetching weather info")

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlCurrentWeather, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lat", latitude)
	q.Add("lon", longitude)
	q.Add("appid", apiKey)
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func currentAQI() (map[string]interface{}, error) {
	log.Debugln("Fetching air quality info")

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlCurrentAQI, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lat", latitude)
	q.Add("lon", longitude)
	q.Add("appid", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func hourlyWeather() (map[string]interface{}, error) {
	log.Debugln("Fetching hourly weather info")

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlHourlyWeather, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lat", latitude)
	q.Add("lon", longitude)
	q.Add("appid", apiKey)
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
