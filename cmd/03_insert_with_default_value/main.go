package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   int64
	Name string `gorm:"default:galeone"` // 这就是默认值
	Age  int64  `gorm:"default:18"` // 默认值
}

func main() {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connecct database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	db.Create(&User{
		Name: "",
	})
}