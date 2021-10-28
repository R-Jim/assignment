package services

import (
	"errors"
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/utils"
	"net/http"

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
	tweetID, userID, ok := utils.GetUserIDAndTweetID(c)
	if !ok {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid Tweet ID or User ID provided")
		return
	}

	tweet := new(dao.Tweet)
	if tweet = utils.GetTweet(tweetID); tweet != nil {
		if err := db.MYSQL.Create(&dao.Like{TweetID: tweetID, UserID: userID}).Error; err == nil {
			db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count + 1"))
			tweet.LikeCount += 1
		}
		c.JSON(http.StatusOK, gin.H{"data": tweet})
	} else {
		utils.Error(c, http.StatusNotFound, "Cannot find tweet")
	}
}

func UnlikeTweet(c *gin.Context) {
	tweetID, userID, ok := utils.GetUserIDAndTweetID(c)
	if !ok {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid Tweet ID or User ID provided")
		return
	}

	tweet := new(dao.Tweet)
	if tweet = utils.GetTweet(tweetID); tweet != nil {
		like := &dao.Like{TweetID: tweetID, UserID: userID}
		if err := db.MYSQL.Where(like).First(&like).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.MYSQL.Delete(like).Error; err == nil {
				db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count - 1"))
				tweet.LikeCount -= 1
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": tweet})
	} else {
		utils.Error(c, http.StatusNotFound, "Cannot find tweet")
	}
}
