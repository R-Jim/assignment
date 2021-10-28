package utils

import (
	"errors"
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserIDAndTweetID(c *gin.Context) (tweetID uint, userID uint, ok bool) {
	ok = true
	if id, err := strconv.Atoi(c.Param("tweetID")); err != nil {
		ok = false
		return
	} else {
		tweetID = uint(id)
	}
	if id, err := strconv.Atoi(c.Param("userID")); err != nil {
		ok = false
		return
	} else {
		userID = uint(id)
	}
	return
}

func GetTweet(tweetID uint) *dao.Tweet {
	tweet := &dao.Tweet{}
	if err := db.MYSQL.Where("id = ?", tweetID).First(&tweet).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return tweet
	}
	return nil
}
