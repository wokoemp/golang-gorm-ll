package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// connect db driver
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		//panic("failed to connecct database")
	}

	// migrate the schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		fmt.Print(err)
	}

	//create
	db.Create(&Product{Code: "D32", Price: 100})

	// read
	// 这里的read 的结果，是赋值给了product对象
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "F21" )

	// update - update product's price to 200
	// 把find 的值找出来，然后修改它的价格字段，这里是只修改一条数据
	//db.Model(&product).Update("price", 200)
	// update - update multiple fields
	//db.Model(&product).Updates(Product{Code: "D32", Price: 200}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "D32"})
	//
	//// delete - delete product
	//db.Delete(&product, 1)
}