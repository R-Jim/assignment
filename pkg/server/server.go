package server

import (
	"jim/twitter/pkg/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	//Users
	router.GET("/users", services.TokenAuthMiddleware(), services.GetUsers)
	router.POST("/users", services.TokenAuthMiddleware(), services.CreateUser)
	//Tweets
	router.POST("/tweets", services.TokenAuthMiddleware(), services.CreateTweet)
	router.PUT("/tweets/:tweetID/users/:userID/like", services.TokenAuthMiddleware(), services.LikeTweet)
	router.PUT("/tweets/:tweetID/users/:userID/unlike", services.TokenAuthMiddleware(), services.UnlikeTweet)
	//Authentication
	router.POST("/login", services.Login)
	router.POST("/logout", services.TokenAuthMiddleware(), services.Logout)
	router.POST("/token/refresh", services.RefreshToken)
	router.Run("localhost:8080")

}
