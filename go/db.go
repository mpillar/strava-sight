package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DatabaseInit() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "tristrava@/tristrava?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return db, nil

	/*
		// Create
		db.Create(&Product{Code: "L1212", Price: 1000})

		// Read
		var product Product
		db.First(&product, 1)                   // find product with id 1
		db.First(&product, "code = ?", "L1212") // find product with code l1212

		// Update - update product's price to 2000
		db.Model(&product).Update("Price", 2000)

		// Delete - delete product
		db.Delete(&product)
	*/
}