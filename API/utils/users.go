package utils

import (
	"os"
		
	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

func CheckUser(email string) (bool, *models.User, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get user
	userRequest := new(models.User)
	has, err := DBengine.Table("user").Where("email = ?", email).Desc("id").Get(userRequest)
	if err != nil {
		return false, userRequest, DBengine, err
	}

	// User not found
	if !has {
		return false, userRequest, DBengine, nil
	}

	return true, userRequest, DBengine, nil
}

func CheckUserByUsername(username string) (bool, *models.User, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get user
	userRequest := new(models.User)
	has, err := DBengine.Table("user").Where("username = ?", username).Desc("id").Get(userRequest)
	if err != nil {
		return false, userRequest, DBengine, err
	}

	// User not found
	if !has {
		return false, userRequest, DBengine, nil
	}

	return true, userRequest, DBengine, nil
}


/* This function fetches an user's avatar */
func GetAvatar(id int64) (bool, *models.UserAvatar, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, nil, err
	}

	// Get user avatar
	userAvatarRequest := new(models.UserAvatar)
	has, err := DBengine.Table("user_avatar").Where("owner_id = ?", id).Desc("id").Get(userAvatarRequest)
	if err != nil {
		return false, userAvatarRequest, DBengine, err
	}

	// User not found
	if !has {
		return false, userAvatarRequest, DBengine, nil
	}

	return true, userAvatarRequest, DBengine, nil
}

func CreateAvatarUrl(filename string) (string, error) {
	// Load url from .env
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	// Get image url
	URL := os.Getenv("APP_URL") + "/users/avatar/" + filename

	return URL, nil
}
