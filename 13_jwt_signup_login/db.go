package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("authdemo1.db"))
	if err != nil {
		log.Fatal("Faild to connect with Database")
	}
	db.AutoMigrate(&User{}, &Book{}) //here we migrating the struct
	return db
}
