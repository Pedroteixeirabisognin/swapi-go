package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := "root:MySql2019!@tcp(127.0.0.1:3306)/testedb?charset=utf8mb4&parseTime=True&loc=Local"
	//"host=localhost port=3306 user=root dbname=testedb sslmode=disable password=MySql2019!"

	database, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
