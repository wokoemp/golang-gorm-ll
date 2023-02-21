package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserUpdatePractice struct {
	ID        uint   `gorm:"primarykey"`
	CreatedAt uint64 `gorm:"autoCreateTime"` // autoCreateTime 修改存入的格式 utc->timestamps
	UpdatedAt uint64 `gorm:"autoUpdateTime"` // autoCreateTime 修改存入的格式 utc->timestamps
	DeletedAt uint64 `gorm:"index"`          // autoCreateTime 修改存入的格式 utc->timestamps
	Name      string
	Age       uint8
	Birthday  time.Time
}

func main() {
	dbUrl := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	err = db.AutoMigrate(&UserUpdatePractice{})
	if err != nil {
		log.Fatalln("failed to create table")
	}

	// insert value into table
	//var users []UserUpdatePractice = []UserUpdatePractice{
	//	{
	//		Name: "name1",
	//		Age: 18,
	//		Birthday: time.Now().UTC(),
	//	},
	//	{
	//		Name: "name2",
	//		Age: 19,
	//		Birthday: time.Now().UTC(),
	//	},
	//	{
	//		Name: "name3",
	//		Age: 20,
	//		Birthday: time.Now().UTC(),
	//	},
	//	{
	//		Name: "name4",
	//		Age: 21,
	//		Birthday: time.Now().UTC(),
	//	},
	//}
	//db.Create(&users)

	/*
		UPDATE user_update_practices
		SET name = 'hello'
		WHERE id = 1
	*/
	//db.Model(&UserUpdatePractice{}).Where("name = ?", "hello").Update("name", "name1")

	/*
		update user_update_practices
		set birthday=xx
	*/
	//var users []UserUpdatePractice
	//db.Order("id desc").Find(&users)
	//db.Model(&users).Updates(UserUpdatePractice{Birthday: time.Date(2021,2,2,2,2,2,0,time.UTC)})

	// update user_update_practices set name = "hello" where id = 1
	//result := db.Exec("update user_update_practices set name = ? where id = 1", "hello")
	// get the number of affected rows
	//fmt.Printf("rows affected:: %d", result.RowsAffected)

	// update user_update_practices
	// set age = age -10

	/*
		update user_update_practices
		set age = age - 10
		where id = 1
	*/
	db.Model(&UserUpdatePractice{}).Where("id = 1").UpdateColumn("age", gorm.Expr("age + ?", 10))
	// ⚠️ UpdateColumn、UpdateColumns vs Update、Updates:
	// 1. UpdateColumn、UpdateColumns 关闭了hook方法和追踪时间自动更新

}
