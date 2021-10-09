package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var timeout = 5

func main() {

	go timer()

	http.HandleFunc("/", GetStatusWeather)

	fmt.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err.Error())
		return
	}
}

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomWeather() {

	min := 1
	max := 100

	data = Weather{
		Water: rand.Intn(max - min),
		Wind:  rand.Intn(max - min),
	}
	fmt.Println(data)

}

var data Weather

func GetStatusWeather(w http.ResponseWriter, r *http.Request) {

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

func timer() {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		RandomWeather()
		timer()
	})
}
