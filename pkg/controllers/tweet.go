package controllers

import (
	"errors"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/dto"
	"jim/twitter/pkg/models"
	"jim/twitter/pkg/repository"
	"jim/twitter/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTweet(c *gin.Context) {
	var requestBody dto.CreateTweetRequestBody
	if err := utils.GetAndValidateRequestBody(c, &requestBody); err != nil {
		utils.ValidationError(c)
		return
	}
	accessDetails, _ := utils.ExtractTokenMetadata(c.Request)
	userID, _ := utils.FetchAuth(accessDetails)

	tweet := new(models.Tweet)
	tweet.Content = requestBody.Content
	tweet.AuthorID = uint(userID)

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

	repository.LikeTweet(tweet, uint(userID))
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}

func UnlikeTweet(c *gin.Context) {
	tweet, ok := getTweet(c)
	if !ok {
		return
	}
	accessDetails, _ := utils.ExtractTokenMetadata(c.Request)
	userID, _ := utils.FetchAuth(accessDetails)
	repository.UnlikeTweet(tweet, uint(userID))
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}

func getTweet(c *gin.Context) (tweet *models.Tweet, ok bool) {
	tweetID, ok := utils.GetTweetID(c)
	if !ok {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid Tweet ID provided")
		return
	}
	tweet, err := repository.GetTweetByID(tweetID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Cannot find tweet")
	} else if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
	} else {
		return tweet, true
	}
	return nil, false
}
