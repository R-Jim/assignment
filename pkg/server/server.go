package server

import (
	"jim/twitter/pkg/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/users", services.TokenAuthMiddleware(), services.GetUsers)
	router.POST("/users", services.TokenAuthMiddleware(), services.CreateUser)
	router.POST("/login", services.Login)
	router.POST("/logout", services.TokenAuthMiddleware(), services.Logout)
	router.Run("localhost:8080")

}
