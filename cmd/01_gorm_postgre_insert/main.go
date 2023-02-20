package main

import (
	"database/sql"
	"errors"
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

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Printf("executing before save")
	return errors.New("")
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error)  {
	fmt.Printf("executing before create")
	return errors.New("")
}


func (u *User) AfterSave(tx *gorm.DB) (err error) {
	fmt.Printf("executing after save")
	return errors.New("")
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("executing after save")
	return errors.New("")
}

func main()  {
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

	//  insert 1 == insert into users("name","age") VALUES('Jinzhu', 19)
	//user := User{Name: "hook1", Age: 19}
	//db.Create(&user)


	// insert 2 == insert into users("name", "age) values('jinzhu3', 3)
	//user := User{Name: "jinzhu3", Age: 3}
	//db.Select("Name", "Age").Create(&user)


	// batch insert
	// configuration: CreateBatchSize: 1000 the number of batching insert
	//var users = []User{{Name: "h5", Age: 1},{Name: "h6", Age: 2}}
	//db.Create(&users)
	//for _, user := range users {
	//	fmt.Printf("%+v",user)
	//}


	// create hooks
	// model 实例实现hooks 接口：BeforeSave、BeforeCreate、AfterSave、AfterCreate
	// execution order: BeforeSave -> before create -> after save -> after create

	// create from map
	// 通过这种方式插入的，gorm.model的时间字段都不会自动插入
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "map1",
		"age": 1, // 注意逗号一定要有
	})



}
