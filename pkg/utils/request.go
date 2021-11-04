package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAndValidateRequestBody(c *gin.Context, requestBody interface{}) (err error) {
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		return err
	}
	err = validator.New().Struct(requestBody)
	if err != nil {
		return err
	}
	return nil
}
