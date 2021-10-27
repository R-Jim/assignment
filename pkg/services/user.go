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

func Login(c *gin.Context) {
	var u dao.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user := dao.User{}
	db.MYSQL.Find(&user)
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := utils.CreateToken(uint64(user.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := utils.CreateAuth(uint64(user.ID), token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
