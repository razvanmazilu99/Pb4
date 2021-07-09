package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb4/db"
	"pb4/entity"
)

func GetClass(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var class entity.Class

	result := db.GetDB().Where("id=?", id).Find(&class)

	if result.RecordNotFound() {
		http.Error(rw, "No record found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	classBytes, err := json.Marshal(class)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(classBytes)
}

func PostClass(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var class entity.Class
	err = json.Unmarshal(bodyBytes, &class)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	db.GetDB().Create(&class)

	fmt.Println(class)
	rw.Write(bodyBytes)
}

func UpdateClass(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var inputClass entity.Class
	err = json.Unmarshal(bodyBytes, &inputClass)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var class entity.Class
	db.GetDB().Model(&class).Updates(inputClass)

	fmt.Println(class)
	rw.Write(bodyBytes)
}

func DeleteClass(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	result := db.GetDB().Delete(&entity.Class{}, "id=?", id)

	if result.Error != nil {
		http.Error(rw, "Internal error. Please try again after a while", http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Record successfully deleted"))
}

func ListOfClasses(rw http.ResponseWriter, r *http.Request) {

	var class entity.Class

	result := db.GetDB().Find(&class)

	if result.RecordNotFound() {
		http.Error(rw, "No record found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	classBytes, err := json.Marshal(class)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(classBytes)
}
