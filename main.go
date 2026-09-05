package main

import (
	"log"

	"github.com/sonalgaud12/Microservices/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.fatal("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("sucessfully conntected to db")
}
