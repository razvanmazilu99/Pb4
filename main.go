package main

import (
	"log"
	"net/http"
	"pb4/config"
	"pb4/db"
	"pb4/rest"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Not able to load env file")
	}

	err = config.InitConfig()

	if err != nil {
		log.Fatal("Not able to create config")
	}

	db.Initdatabase()

	var endpoint1 = "/student"
	var endpoint2 = "/class"

	router := chi.NewRouter()
	router.Route("/" + config.GetConfig().APPVersion, func(r chi.Router) {
		r.Get(endpoint1, rest.GetStudent)
		r.Post(endpoint1, rest.PostStudent)
		r.Delete(endpoint1, rest.DeleteStudent)
		r.Put(endpoint1, rest.UpdateStudent)

		r.Get(endpoint2, rest.GetClass)
		r.Post(endpoint2, rest.PostClass)
		r.Delete(endpoint2, rest.DeleteClass)
		r.Put(endpoint2, rest.UpdateClass)

		r.Get(endpoint1, rest.ListOfStudents)
		r.Get(endpoint2, rest.ListOfClasses)

	//	r.Post("/enroll", rest.Enroll)
	})
	http.ListenAndServe(":" + config.GetConfig().Port, router)
	
}