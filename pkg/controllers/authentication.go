package controllers

import (
	"errors"
	"fmt"
	"jim/twitter/pkg/dto"
	"jim/twitter/pkg/repository"
	"jim/twitter/pkg/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
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
	var requestBody dto.LoginRequestBody
	if err := utils.GetAndValidateRequestBody(c, &requestBody); err != nil {
		utils.ValidationError(c)
		return
	}
	user, err := repository.GetUserByUsernameAndPassword(requestBody.Username, requestBody.Password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.AuthenticationError(c, "Please provide valid login details")
		return
	}
	token, err := utils.CreateToken(uint64(user.ID))
	if err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = utils.CreateAuth(uint64(user.ID), token)
	if err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, err.Error())
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
	var requestBody dto.RefreshTokenRequestBody
	if err := utils.GetAndValidateRequestBody(c, &requestBody); err != nil {
		utils.ValidationError(c)
		return
	}
	refreshToken := requestBody.RefreshToken

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
