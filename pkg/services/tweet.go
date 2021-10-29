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
	tweet, ok := getTweet(c)
	if !ok {
		return
	}
	accessDetails, _ := utils.ExtractTokenMetadata(c.Request)
	userID, _ := utils.FetchAuth(accessDetails)

	like := &dao.Like{TweetID: tweet.ID, UserID: uint(userID)}
	if errors.Is(db.MYSQL.Where(like).First(&like).Error, gorm.ErrRecordNotFound) {
		if db.MYSQL.Create(like).RowsAffected == 1 {
			db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count + 1"))
			tweet.LikeCount += 1
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}

func UnlikeTweet(c *gin.Context) {
	tweet, ok := getTweet(c)
	if !ok {
		return
	}
	accessDetails, _ := utils.ExtractTokenMetadata(c.Request)
	userID, _ := utils.FetchAuth(accessDetails)

	like := &dao.Like{TweetID: tweet.ID, UserID: uint(userID)}
	if db.MYSQL.Delete(like).RowsAffected == 1 {
		db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count - 1"))
		tweet.LikeCount -= 1
	}
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}

func getTweet(c *gin.Context) (tweet *dao.Tweet, ok bool) {
	tweetID, ok := utils.GetTweetID(c)
	if !ok {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid Tweet ID provided")
		return
	}
	err := db.MYSQL.Where("id = ?", tweetID).First(&tweet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Cannot find tweet")
	} else if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
	} else {
		return tweet, true
	}
	return nil, false
}
