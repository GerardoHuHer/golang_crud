package main

import (
	"net/http"
	"github.com/GerardoHuHer/go_crud/routes"
	"github.com/gorilla/mux"
	"github.com/GerardoHuHer/go_crud/db"
	"github.com/GerardoHuHer/go_crud/models"
)


func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Rover{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":4000", r)
}
