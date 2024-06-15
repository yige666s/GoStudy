package main

import (
	goj "github.com/jinzhu/gorm"   // 这是老版1.x
	gormsql "gorm.io/driver/mysql" // 这是新版2.x
	"gorm.io/gorm"
	goa "gorm.io/gorm"
)

func GormCase() {
	db, err := goa.Open(gormsql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql")
	}

	db2, err2 := goj.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err2 != nil {
		panic("failed to connect to mysql")
	}
	defer db2.Close() // goj这个包中有close

}
