package helper

import (
	"belajar_api_sendiri/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JsonOutput(w http.ResponseWriter, o interface{}, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var status int
	if code == 200 {
		status = 1
	} else {
		status = 0
		outputErr := responseErr{
			Status:  status,
			Message: message,
		}
		resErr, err := json.Marshal(outputErr)
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(resErr)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	output := response{
		Status:  status,
		Message: message,
		Data:    o,
	}

	res, err := json.Marshal(output)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func JsonInput(r *http.Request) *models.Person {
	decoder := json.NewDecoder(r.Body)
	var result models.Person
	err := decoder.Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &result
}
