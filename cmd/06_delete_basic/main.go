package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserDeletePractice struct {
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

	err = db.AutoMigrate(&UserDeletePractice{})
	if err != nil {
		log.Fatalln("failed to create table")
	}

	//insert value into table
	//var users []UserDeletePractice = []UserDeletePractice{
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

	// ⚠️ 删除一条记录，如果没有指定primary key 则会触发批量删除

}
