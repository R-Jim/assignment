package utils

import (
	"jim/twitter/pkg/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationError(c *gin.Context, details string) {
	Error(c, http.StatusUnauthorized, details)
}

func Error(c *gin.Context, statusCode int, details string) {
	c.JSON(statusCode, &dao.Error{Code: statusCode, Details: details})
}
