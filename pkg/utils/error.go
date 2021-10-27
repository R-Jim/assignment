package utils

import (
	"jim/twitter/pkg/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationError(c *gin.Context, details string) {
	c.JSON(http.StatusUnauthorized, &dao.Error{Code: http.StatusUnauthorized, Details: details})
}
