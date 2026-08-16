package router

import (
	"github.com/GerardoHuHer/go_crud/internal/handler"
	"github.com/gorilla/mux"
)

func New(roverHandler *handler.RoverHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/create_rover", roverHandler.Create).Methods("POST")
	api.HandleFunc("/get_rover/{id}", roverHandler.GetByIdHandler).Methods("GET")
	api.HandleFunc("/update_rover/{id}", roverHandler.UpdateByIdHandler).Methods("PATCH")
	api.HandleFunc("/delete_rover/{id}", roverHandler.DeleteByIdHandler).Methods("DELETE")

	return r
}
