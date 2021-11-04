package utils

import (
	"jim/twitter/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationError(c *gin.Context, details string) {
	Error(c, http.StatusUnauthorized, details)
}

func ValidationError(c *gin.Context) {
	Error(c, http.StatusUnprocessableEntity, "Failed to validate request")
}

func Error(c *gin.Context, statusCode int, details string) {
	c.JSON(statusCode, &dto.Error{Code: statusCode, Details: details})
}
