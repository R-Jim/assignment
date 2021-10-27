package services

import (
	"fmt"
	"jim/twitter/pkg/dao"
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Logout(c *gin.Context) {
	accessDetails, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		utils.AuthenticationError(c, "Unable to extract token metadata")
		return
	}
	deleted, delErr := utils.DeleteAuth(accessDetails.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		utils.AuthenticationError(c, "Unable to remove Authentication info")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
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

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessDetails, err := utils.ExtractTokenMetadata(c.Request)
		if err != nil {
			utils.AuthenticationError(c, "Unable to extract token metadata")
			c.Abort()
			return
		}
		if _, err := utils.FetchAuth(accessDetails); err != nil {
			utils.AuthenticationError(c, "Error when Fetching Authentication info")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RefreshToken(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		utils.AuthenticationError(c, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		utils.AuthenticationError(c, err.Error())
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := utils.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			utils.AuthenticationError(c, "Unable to delete Authentication info")
			return
		}
		//Create new pairs of refresh and access tokens
		token, createErr := utils.CreateToken(userId)
		if createErr != nil {
			utils.Error(c, http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := utils.CreateAuth(userId, token)
		if saveErr != nil {
			utils.Error(c, http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		utils.AuthenticationError(c, "Refresh token expired")
	}
}
