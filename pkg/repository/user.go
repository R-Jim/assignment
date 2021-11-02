package repository

import (
	"jim/twitter/pkg/db"
	"jim/twitter/pkg/models"
)

func GetUserByUsernameAndPassword(username string, password string) (user *models.User, err error) {
	err = db.MYSQL.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}
