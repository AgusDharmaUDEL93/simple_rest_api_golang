package main

import (
	"belajar_api_sendiri/controller"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controller.GetAllData(w)
		} else if r.Method == "POST" {
			controller.CreateData(w, r)
		} else {
			http.Error(w, "BadRequest", http.StatusBadRequest)
		}
	})
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			id := strings.TrimPrefix(r.URL.Path, "/person/")
			controller.GetDataById(w, id)
		} else {
			http.Error(w, "BadRequest", http.StatusBadRequest)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
