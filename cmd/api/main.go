package main

import (
	"log"
	"net/http"

	"github.com/GerardoHuHer/go_crud/internal/config"
	"github.com/GerardoHuHer/go_crud/internal/domain"
	"github.com/GerardoHuHer/go_crud/internal/handler"
	"github.com/GerardoHuHer/go_crud/internal/repository/postgres"
	"github.com/GerardoHuHer/go_crud/internal/router"
	"github.com/GerardoHuHer/go_crud/internal/service"
)

func main() {
	db, err := config.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Para hacer las migraciones
	if err := db.AutoMigrate(&domain.Rover{}); err != nil {
		log.Fatal(err)
	}

	roverRepo := postgres.NewRoverRepository(db)
	roverSvc := service.NewRoverService(roverRepo)
	roverHandler := handler.NewRoverHandler(roverSvc)

	r := router.New(roverHandler)

	log.Println("listening on port :4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
