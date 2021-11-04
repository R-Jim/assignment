package controllers

import (
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	db.MYSQL.Create(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	users := []models.User{}
	db.MYSQL.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
