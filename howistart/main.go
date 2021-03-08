package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	// Handler function at root path of webserver
	// When a new request comes into the server matching root path, server
	// spawns a new goroutine executing hello()
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

// hello is a HandlerFunc returns a "hello" response to a request.
func hello(w http.ResponseWriter, r *http.Request) {
	// Cast string to byte-slice and write response to client
	w.Write([]byte("hello!"))
}

// query populates a weatherData struct
func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=API-KEY&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData // allocate weatherData struct

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}

type weatherData struct {
	Name string `json:"name"` // backticks surround tags
	// tags are used for customizing JSON key names
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
