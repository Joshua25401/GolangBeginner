package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	dsn := "root:@tcp(localhost:3306)/tes_go?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	/*If err found or != nil then log.err*/
	if err != nil {
		panic("DB ERR : " + err.Error())
	}
	log.Println("Connected to Database")
}
