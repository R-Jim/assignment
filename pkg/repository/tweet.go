package repository

import (
	"errors"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/models"

	"gorm.io/gorm"
)

func GetTweetByID(tweetID uint) (tweet *models.Tweet, err error) {
	err = db.MYSQL.Where("id = ?", tweetID).First(&tweet).Error
	return tweet, err
}

func LikeTweet(tweet *models.Tweet, userID uint) {
	like := &models.Like{TweetID: tweet.ID, UserID: userID}
	if errors.Is(db.MYSQL.Where(like).First(&like).Error, gorm.ErrRecordNotFound) {
		if db.MYSQL.Create(like).RowsAffected == 1 {
			db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count + 1"))
			tweet.LikeCount += 1
		}
	}
}

func UnlikeTweet(tweet *models.Tweet, userID uint) {
	like := &models.Like{TweetID: tweet.ID, UserID: uint(userID)}
	if db.MYSQL.Delete(like).RowsAffected == 1 {
		db.MYSQL.Model(&tweet).UpdateColumn("like_count", gorm.Expr("like_count - 1"))
		tweet.LikeCount -= 1
	}
}
