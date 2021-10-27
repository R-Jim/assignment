package services

import (
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := new(dao.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	db.MYSQL.Create(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	users := []dao.User{}
	db.MYSQL.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
