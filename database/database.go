package database

import (
	"fmt"
	"rest-api/model"

	"github.com/jinzhu/gorm"
)

type User = model.User
type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "log/user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	// defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
