package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
	"log"
)


var DSN string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", 
	os.Getenv("DB_HOST"), 
	os.Getenv("POSTGRES_USER"), 
	os.Getenv("POSTGRES_PASSWORD"), 
	os.Getenv("POSTGRES_DB"))

var DB *gorm.DB


func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
