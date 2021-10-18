package main

import (
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
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
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

type DataWeather struct {
	Data        Weather `json:"data"`
	StatusWater string  `json:"status_water"`
	StatusWind  string  `json:"status_wind"`
}

func RandomWeather() {

	min := 1
	max := 100

	data = Weather{
		Water: rand.Intn(max - min),
		Wind:  rand.Intn(max - min),
	}
	//fmt.Println(data)

}

var data Weather

func GetStatusWeather(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")

	var statusWater string
	var statusWind string
	dataWater := data.Water
	dataWind := data.Wind
	if dataWater < 6 {
		statusWater = "status aman"
	}
	if dataWater >= 6 && dataWater <= 8 {
		statusWater = "status siaga"

	}
	if dataWater > 8 {
		statusWater = "status bahaya"
	}
	if dataWind < 7 {
		statusWind = "status aman"

	}
	if dataWind >= 7 && dataWind <= 15 {
		statusWind = "status siaga"
	}
	if dataWind > 15 {
		statusWind = "status bahaya"
	}

	dataWeather := DataWeather{
		Data:        data,
		StatusWater: statusWater,
		StatusWind:  statusWind,
	}

	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, dataWeather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// json data

	//jsonInBytes, err := json.Marshal(dataWeather)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//_, err = w.Write(jsonInBytes)
	//if err != nil {
	//	return
	//}
}

func timer() {
	RandomWeather()
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		timer()
	})
}
