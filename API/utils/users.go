package utils

import (
	"os"

	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

func GetUser(email string) (bool, *models.User, *xorm.Engine, error) {
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

func GetUserByUsername(username string) (bool, *models.User, *xorm.Engine, error) {
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

func GetUserByToken(token string) (bool, *models.User, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Search the user in the "user" table
	userModel := new(models.User)
	hasUser, err := DBengine.Table("user").Where("access_token = ?", token).Desc("id").Get(userModel)
	if err != nil {
		return false, userModel, DBengine, err
	}

	return hasUser, userModel, DBengine, nil
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

func GetFriendsList(username string) (bool, *models.User, *models.UserFriends, error) {
	hasUser, userModel, DBengine, err := GetUserByUsername(username)
	if err != nil {
		return hasUser, userModel, nil, err
	}

	// Search User Friends
	userFriendsRequest := new(models.UserFriends)
	_, err = DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Get(userFriendsRequest)
	if err != nil {
		return hasUser, userModel, userFriendsRequest, err
	}

	return hasUser, userModel, userFriendsRequest, nil
}