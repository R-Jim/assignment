package services

import (
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTweet(c *gin.Context) {
	tweet := new(dao.Tweet)
	if err := c.ShouldBindJSON(&tweet); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	if err := db.MYSQL.Create(tweet).Error; err != nil {
		utils.Error(c, http.StatusBadRequest, "Unable to create tweet")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": tweet})
}

func LikeTweet(c *gin.Context) {
	var tweetid, userid uint
	if id, err := strconv.Atoi(c.Param("tweetid")); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid Tweet Id provided")
		return
	} else {
		tweetid = uint(id)
	}
	if id, err := strconv.Atoi(c.Param("userid")); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid User Id provided")
		return
	} else {
		userid = uint(id)
	}
	tweet := new(dao.Tweet)
	if err := db.MYSQL.Find(&tweet, "id = ?", tweetid).Error; err == nil {
		if err := db.MYSQL.Create(&dao.Like{TweetId: tweetid, UserId: uint(userid)}).Error; err == nil {
			db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count + 1"))
			tweet.LikeCount += 1
		}
		c.JSON(http.StatusOK, gin.H{"data": tweet})
	} else {
		utils.Error(c, http.StatusNotFound, "Cannot find tweet")
	}
}
