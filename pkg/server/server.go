package server

import (
	"jim/twitter/pkg/controllers"

	"github.com/gin-gonic/gin"
)

const (
	authentication = "/auth"
	users          = "/users"
	tweets         = "/tweets"
	likes          = tweets + "/:tweetID/like"
)

func Run() {
	router := gin.Default()
	//Users
	router.GET(users, controllers.TokenAuthMiddleware(), controllers.GetUsers)
	router.POST(users, controllers.TokenAuthMiddleware(), controllers.CreateUser)
	//Tweets
	router.POST(tweets, controllers.TokenAuthMiddleware(), controllers.CreateTweet)
	router.PUT(likes, controllers.TokenAuthMiddleware(), controllers.LikeTweet)
	router.DELETE(likes, controllers.TokenAuthMiddleware(), controllers.UnlikeTweet)
	//Authentication
	router.POST(authentication+"/login", controllers.Login)
	router.POST(authentication+"/logout", controllers.TokenAuthMiddleware(), controllers.Logout)
	router.POST(authentication+"/token/refresh", controllers.RefreshToken)
	router.Run("localhost:8080")

}
