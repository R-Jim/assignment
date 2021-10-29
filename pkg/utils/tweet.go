package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTweetID(c *gin.Context) (tweetID uint, ok bool) {
	if id, err := strconv.Atoi(c.Param("tweetID")); err != nil {
		ok = false
		return
	} else {
		return uint(id), true
	}
}
