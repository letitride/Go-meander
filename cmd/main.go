package main

import (
	"encoding/json"
	"meander"
	"net/http"
	"runtime"
)

func main() {
	//利用するCPU数の最大値の指定=すべて
	runtime.GOMAXPROCS(runtime.NumCPU())
	//meander.APIKey = "TODO"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
