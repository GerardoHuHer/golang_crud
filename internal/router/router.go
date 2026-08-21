package router

import (
	"github.com/GerardoHuHer/go_crud/internal/handler"
	"github.com/gorilla/mux"
)

func New(roverHandler *handler.RoverHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/create_vehiculo", roverHandler.Create).Methods("POST")
	api.HandleFunc("/get_vehiculo/{id}", roverHandler.GetByIdHandler).Methods("GET")
	api.HandleFunc("/update_vehiculo/{id}", roverHandler.UpdateByIdHandler).Methods("PATCH")
	api.HandleFunc("/delete_vehiculo/{id}", roverHandler.DeleteByIdHandler).Methods("DELETE")
	api.HandleFunc("/get_all_vehiculos", roverHandler.GetAllHandler).Methods("GET")

	return r
}
