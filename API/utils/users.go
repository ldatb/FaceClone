package utils

import (
	"os"
		
	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

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
