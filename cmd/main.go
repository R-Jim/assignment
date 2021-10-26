package main

import (
	"fmt"
	"jim/twitter/pkg/server"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("twitter.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to the database")
	}
	server.Run()
}
