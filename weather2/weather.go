package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WeatherReport struct {
	Report string `json:"report"`
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	var resp = WeatherReport{
		Report: "Returning the weather from weather 2",
	}
	fmt.Println("Getting the weather report")
	fmt.Println(resp)
	json.NewEncoder(w).Encode(resp)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/current", getCurrentWeather)

	fmt.Println("Weather plugin 2 starting on port 4444")
	log.Fatal(http.ListenAndServe(":4444", router))
}

func main() {
	handleRequests()
}
