package db

import (
	"fmt"
	"jim/twitter/pkg/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "user:123@tcp(0.0.0.0:3306)/twitter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to the database")
	}

	DB = db
	migrate()
}

func migrate() {
	DB.AutoMigrate(&dao.User{})
}
