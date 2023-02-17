package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)


//1. 定义models
type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

func main()  {
	//_, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN: "user=postgres password=mysecretpassword dbname=gorm port=5437 sslmode=disable TimeZone=Asia/Shanghai",
	//	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//}), &gorm.Config{})
	//if err != nil {
	//	fmt.Print(err)
	//}

	dbURL := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connecct database")
	}

	// migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Print(err)
	}

	// ==================== 3.写入数据库的n种姿势==============================

	//  insert
	user := User{Name: "Jinzhu", Age: 19}
	//db.Create(&user)

	//
	user.Age = 1
	user.Name = "user1"
	db.Select("Name", "Age", &user)

}
