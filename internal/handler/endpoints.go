package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"weather-tracker/internal/models"
)

func(h *Handler) weather(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")
	data, err := query(city)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func query(city string) (models.WeatherData, error) {
	apiConfig, err := LoadApiConfig(".apiConfig")
	if err != nil {
		return models.WeatherData{}, err
	}

	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,uk&appid=%s&units=metric", city, apiConfig.OpenWeatherMapApiKey))

	defer resp.Body.Close()

	var d models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return models.WeatherData{}, err
	}

	return d, nil
}

func LoadApiConfig(filename string) (models.ApiConfigData, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return models.ApiConfigData{}, err
	}

	var c models.ApiConfigData

	err = json.Unmarshal(content, &c)
	if err != nil {
		return models.ApiConfigData{}, err
	}

	return c, nil
}