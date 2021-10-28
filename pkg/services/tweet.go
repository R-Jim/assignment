package services

import (
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTweet(c *gin.Context) {
	tweet := new(dao.Tweet)
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	db.MYSQL.Create(tweet)
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}
