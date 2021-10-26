package main

import (
	"fmt"
	"jim/twitter/pkg/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// _, err := gorm.Open(sqlite.Open("twitter.db"), &gorm.Config{})
	dsn := "user:123@tcp(0.0.0.0:3306)/twitter?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to the database")
	}
	server.Run()
}
