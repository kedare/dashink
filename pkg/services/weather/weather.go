package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	logger = log.Default()

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
	logger.Printf("Getting weather icon: %s", fileName)

	// If the file exists, return it
	if data, err := os.ReadFile(fileName); err == nil {
		logger.Printf("Using cached weather icon: %s", fileName)
		return data, nil
	}

	// Get the icon URL
	url := iconURL(iconCode)

	// Download the icon
	logger.Printf("Downloading weather icon: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the icon data
	iconData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Save the icon
	logger.Printf("Saving weather icon: %s", fileName)
	if err := os.WriteFile(fileName, iconData, 0644); err != nil {
		return nil, err
	}

	return iconData, nil
}

func currentWeather() (map[string]interface{}, error) {
	logger.Println("Fetching weather info")

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
	logger.Println("Fetching air quality info")

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
	logger.Println("Fetching hourly weather info")

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
