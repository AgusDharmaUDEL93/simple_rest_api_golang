package main

import (
	"belajar_api_sendiri/controller"
	"fmt"
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

	fmt.Println("Serving in api at port 8080")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
