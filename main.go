package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type weatherData struct{
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:main`
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/weather/",city_weather)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!!"))
}

func city_weather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path,"/",3)[2]

	data, err := getCityWeatherInfo(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	defer r.Body.Close()

	w.Header(). Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func getCityWeatherInfo(city string) (weatherData, error){
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d , nil
}
