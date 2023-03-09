package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	_, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connecct database")
	}

	fmt.Printf("dada")

}
