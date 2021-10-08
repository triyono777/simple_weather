package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetStatusWeather)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func GetStatusWeather(w http.ResponseWriter, r *http.Request) {
	data := Weather{
		Water: 5,
		Wind:  7,
	}
	dataWeather := map[string]Weather{
		"data": data,
	}

	jsonInBytes, err := json.Marshal(dataWeather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonInBytes)
	if err != nil {
		return
	}
}
