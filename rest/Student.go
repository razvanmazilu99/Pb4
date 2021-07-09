package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb4/db"
	"pb4/entity"

	"github.com/sirupsen/logrus"
)

func GetStudent(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var student entity.Student

	result := db.GetDB().Where("id=?", id).Find(&student)

	if result.RecordNotFound() {
		http.Error(rw, "No record found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	studentBytes, err := json.Marshal(student)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(studentBytes)
}

func PostStudent(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	db.GetDB().Create(&student)

	fmt.Println(student)
	rw.Write(bodyBytes)
}

func UpdateStudent(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var inputStudent entity.Student
	err = json.Unmarshal(bodyBytes, &inputStudent)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var student entity.Student
	db.GetDB().Model(&student).Updates(inputStudent)

	fmt.Println(student)
	rw.Write(bodyBytes)
}

func DeleteStudent(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	result := db.GetDB().Delete(&entity.Student{}, "id=?", id)

	if result.Error != nil {
		http.Error(rw, "Internal error. Please try again after a while", http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Record successfully deleted"))
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)

	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}

	return false
}

func ListOfStudents(rw http.ResponseWriter, r *http.Request) {

	var students []entity.Student

	result := db.GetDB().Find(&students)

	if result.RecordNotFound() {
		http.Error(rw, "No record found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	studentBytes, err := json.Marshal(students)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(studentBytes)
}

func EnrollStudent(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	//db.GetDB().Preload("classes").Find(&student)
	res := db.GetDB().Model(student).Association("classes").Append(student.Classes) 

	if res.Error != nil {
		http.Error(rw, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(student)
	rw.Write(bodyBytes)
}
