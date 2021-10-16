package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homework-backend/storage"
	"log"
)

func main() {
	s, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	s.Model(&storage.Result{}).Where("1 = 1").Delete(storage.Result{})
	s.Model(&storage.Answer{}).Where("1 = 1").Delete(storage.Answer{})

	log.Printf("\U0001F9F9 Database is succesfully reset.")
}
