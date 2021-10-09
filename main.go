package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path"
	"time"
)

var timeout = 5

func main() {

	go timer()

	http.HandleFunc("/", HomeWeather)
	http.HandleFunc("/getWeather", GetStatusWeather)

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
	checkDataweather()

}

var data Weather

func HomeWeather(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dataContoh = map[string]interface{}{
		"water": data.Water,
		"wind":  data.Wind,
	}

	err = tmpl.Execute(w, dataContoh)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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

func checkDataweather() {
	dataWater := data.Water
	dataWind := data.Wind
	if dataWater < 6 {
		fmt.Printf("\n Water %s m,  status aman", dataWater)
	}
	if dataWater >= 6 && dataWater <= 8 {
		fmt.Printf("\nWater %s m,  status siaga", dataWater)

	}
	if dataWater > 8 {
		fmt.Printf("\nWater %s m,  status bahaya", dataWater)
	}
	if dataWind < 7 {
		fmt.Printf("\nWind %s m,  status aman", dataWind)

	}
	if dataWind >= 7 && dataWind <= 15 {
		fmt.Printf("\nWind %s m,  status siaga", dataWind)
	}
	if dataWind > 15 {
		fmt.Printf("\nWind %s m,  status bahaya", dataWind)
	}

}
