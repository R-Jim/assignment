package db

import (
	"fmt"
	"jim/twitter/pkg/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQL *gorm.DB

func InitMySQL() {
	dsn := os.Getenv("MYSQL_DSN")
	if len(dsn) == 0 {
		dsn = "user:123@tcp(0.0.0.0:3306)/twitter?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to the database")
	}

	MYSQL = db
	migrate()
}

func migrate() {
	MYSQL.AutoMigrate(&models.User{}, &models.Tweet{}, &models.Like{})
}
