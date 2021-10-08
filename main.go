package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

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
//	var d = 15 * time.Second
//	var t = time.Now().Add(d)
//
//	for {
//		if time.Now().Before(t) {
//			continue
//		}
//// do somthing
//	}
	min := 1
	max := 100

	data = Weather{
		Water: rand.Intn(max-min),
		Wind:  rand.Intn(max-min),
	}

}

var data Weather

func GetStatusWeather(w http.ResponseWriter, r *http.Request) {
	RandomWeather()
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
