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
	Age        int
	CreditCard CreditCard
}

func main() {
	dbUrl := "postgres://postgres:mysecretpassword@localhost:5437/gorm"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connecct database")
	}

	//var user User

	// select * from users order by `primary key` limit 1
	//db.First(&user)
	//fmt.Printf("user:: %+v", user)

	// select * from users order by `primary key`  desc limit 1
	//db.Last(&user)
	//fmt.Printf("user:: %+v", user)

	// select * from users limit 1
	//db.Take(&user)
	//fmt.Printf("user:: %+v", user)

	// ⚠️ first 和last 的表如果没有指定primary key ，会用第一个field进行排序

	// the Find method accepts both struct and slice data
	// ⚠️ 不要直接find，要因为find 原理是查一遍表后返回一条，得加上limit，这样不会去查一遍整张表
	//db.Limit(1).Find(&user)
	//fmt.Printf("user:: %+v", user)

	// query from map
	//var users []User
	// select * from users order by users.id limit 1
	//result := map[string]interface{}{}
	//db.Model(&User{}).First(&result)
	//fmt.Printf("result:: %+v", result)

	// select * from users where id = 2
	//var user User
	//db.First(&user, 2) // 2 或者 "2" 都可以
	//fmt.Printf("user:: %+v", user)

	// select * from users where ID in (2,3)
	//var users []User
	//db.Find(&users, []int{2,3})
	//fmt.Printf("users::%+v", users)

	// ⚠️ 如果id 是string ，查询语句
	// select * from users where id = "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	// b.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")

	// retrieving all objects
	//var users []User
	//db.Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select * from users where name = 'judy1' order by id limit 1
	//var user User
	//db.Where("name = ?", "judy1").First(&user)
	//fmt.Printf("user:: %+v", user)

	// select * from users where name <> 'judy'
	// 不等于
	//var users []User
	//db.Where("name <> ?", "judy").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select * from users where name in ('judy2', 'judy3')
	//var users []User
	//db.Where("name in ?", []string{"judy2", "judy3"}).Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select * from users where name like '%dy%'
	//var users []User
	//db.Where("name like ?", "%dy1").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select * from user where name = 'judy1' and age  >= 2
	//var users []User
	//db.Where("name = ? and age >= ?", "judy1", 2).Find(&users)
	//fmt.Printf("users:: %+v", users)

	// not
	// select * from users where not name = 'judy1'
	//var users []User
	//db.Not("name = ?", "judy1").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// or
	// select * from users where name = 'judy1' or name = 'judy2'
	//var users []User
	//db.Where("name = ?", "judy1").Or("name = ?", "judy2").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select specific fields
	// select age from users
	//var users []User
	//db.Select("age").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// order
	// select * from users order by age desc
	//var users []User
	//db.Order("age desc").Find(&users)
	//fmt.Printf("users:: %+v", users)

	// limit
	// select * from users limit 2
	//var users []User
	//db.Limit(2).Find(&users)
	//fmt.Printf("users:: %+v", users)

	// cancel limit with -1
	//var users1 []User
	//var users2 []User
	//db.Limit(2).Find(&users1).Limit(-1).Find(&users2)
	//fmt.Printf("users1:: %+v", users1)
	//fmt.Print("\n=======================\n")
	//fmt.Printf("users2:: %+v", users2)

	// select * from users offset 3
	//var users []User
	//db.Offset(3).Find(&users)
	//fmt.Printf("users:: %+v", users)

	// select * from users offset 3 limit 1
	//var users []User
	//db.Offset(3).Limit(1).Find(&users)
	//fmt.Printf("users:: %+v", users)

	// group by
	// SELECT name, SUM(age) as total from users group by name
	//type result struct {
	//	Name string
	//	Total int
	//}
	//var results []result
	//db.Model(&User{}).Select("name, sum(age) as total").Group("name").Find(&results)
	//fmt.Printf("results:: %+v", results)

	// having 属于group的where
	// select name, sum(age) from users group by name having name = 'judy1'
	//type Result struct {
	//	Name string
	//	Total int
	//}
	//var results []Result
	//db.Model(&User{}).Group("name").Select("name, sum(age) as total").Having("name = ?", "judy1").Find(&results)
	//fmt.Printf("results:: %+v", results)

	// distinct 用于去重
	// select distinct  name from users
	var users []User
	db.Distinct("name").Find(&users)
	fmt.Printf("users:: %+v", users)

}
