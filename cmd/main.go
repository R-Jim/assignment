package main

import (
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/server"
)

func main() {
	db.InitMySQL()
	db.InitRedis()
	server.Run()
}
