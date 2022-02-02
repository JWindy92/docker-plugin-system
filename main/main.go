package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pluginmanager "main.example.com/pkg/pluginmanager"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You hit the API")
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requesting the current weather report")
	resp, err := http.Get("http://weather:4444/current")
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf("Something went wrong")
	}
	jsonResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Something went wrong reading json")
	}
	var report interface{}

	err = json.Unmarshal([]byte(jsonResp), &report)
	if err != nil {
		fmt.Errorf("Somethign went wrong unmarshalling json")
	}
	fmt.Println(report)
	json.NewEncoder(w).Encode(report)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/weather", getWeather)

	fmt.Println("Main API starting on port 8888 with poopy")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func main() {
	pluginmanager.GetPlugins()
	handleRequests()
}
