package services

import (
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/utils"
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
	if _, ok := utils.VerifyRequest(c); !ok {
		return
	}

	users := []dao.User{}
	db.MYSQL.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
