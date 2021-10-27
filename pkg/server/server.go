package server

import (
	"jim/twitter/pkg/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/users", services.GetUsers)
	router.POST("/users", services.CreateUser)
	router.POST("/login", services.Login)
	router.Run("localhost:8080")

}
