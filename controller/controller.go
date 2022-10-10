package controller

import (
	"belajar_api_sendiri/connection"
	"belajar_api_sendiri/helper"
	"belajar_api_sendiri/models"
	"database/sql"
	"fmt"
	"net/http"
)

func GetAllData(w http.ResponseWriter) {
	db, err := connection.ConnectAPI()
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("Select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(rows)
	var result []models.Person
	for rows.Next() {
		each := models.Person{}
		err := rows.Scan(&each.Id, &each.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = append(result, each)
	}
	helper.JsonOutput(w, result, http.StatusOK, "Succes get all data")
}

func GetDataById(w http.ResponseWriter, id string) {
	db, err := connection.ConnectAPI()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := db.Query("SELECT * from person where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data models.Person
	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	helper.JsonOutput(w, data, http.StatusOK, "Succes get data by id")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(rows)
}

func CreateData(w http.ResponseWriter, r *http.Request) {

	db, err := connection.ConnectAPI()
	if err != nil {
		fmt.Println(err)
	}

	var temp = helper.JsonInput(r)
	if err != nil {
		fmt.Println(err)
	}
	insert, err := db.Query("INSERT into person (name) values (?)", temp.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result models.Person
	response, err := db.Query("SELECT * from person where name = ?", temp.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(response *sql.Rows) {
		err := response.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(response)
	for response.Next() {
		err = response.Scan(&result.Id, &result.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	helper.JsonOutput(w, result, http.StatusOK, "Succes create a data")
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(insert)

}
