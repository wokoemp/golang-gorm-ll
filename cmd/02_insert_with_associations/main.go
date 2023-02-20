package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 外键 是user的id
}

type User struct {
	gorm.Model // id 是主键
	Name       string
	CreditCard CreditCard
}

func main() {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connecct database")
	}

	err = db.AutoMigrate(&CreditCard{})
	if err != nil {
		fmt.Printf("create database:: CreditCard:: %+v", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("create database:: User:: %+v", err)
	}

	// 自定进行表关联
	//db.Create(&User{
	//	Name: "jinzhu",
	//	CreditCard: CreditCard{Number: "411111111111"},
	//})

	// 也可以跳过表关联,只插入user 数据
	//user := User{Name: "apple"}
	//db.Omit("CreditCard").Create(&user)

	//
	//db.Omit(clause.Associations).Create(&User{Name: "cucumber"})

	// batch insert
	//var users = []User{{
	//	Name: "judy1",
	//	CreditCard: CreditCard{Number: "100000001"},
	//	},
	//	{
	//		Name: "judy2",
	//		CreditCard: CreditCard{Number: "100000002"},
	//	},
	//	{
	//		Name: "judy3",
	//		CreditCard: CreditCard{Number: "100000003"},
	//	},
	//}
	//
	//db.Create(&users)
}
