package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyRequest(c *gin.Context) (userid uint64, ok bool) {
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return uint64(0), false
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return uint64(0), false
	}
	return uint64(userId), true
}
