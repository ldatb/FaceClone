package utils

import (
	"crypto/rand"
	"math"
	"math/big"
	"os"

	"faceclone-api/data"
	"faceclone-api/data/models"

	"xorm.io/xorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

/* This function generates a random 6-digit Auth code */
func GenerateAuthKey() (string, error) {
	a, err:= rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return "", err
	}
	n := a.String()
	return n[:6], nil
}

func ValidateAuthKey(email string, token string) (bool, error) {
	// Connect to dabase
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	// Search the user token
	// The CheckUser function is not being used because the tables are different
	authRequest := new(models.AuthToken)
	authDb, err := DBengine.Table("auth_token").Where("email = ?", email).Get(authRequest)
	if err != nil {
		return false, err
	}

	// User not found
	if !authDb {
		return false, err
	}

	// Compare tokens
	if token == authRequest.AccessToken {
		// If it's true the token can be deleted from the database
		_, err = DBengine.Table("auth_token").Where("email = ?", email).Delete()
		if err != nil {
			return true, err
		}
			
		return true, err
	} else {
		return false, err
	}
}

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

func GetPost(id int64) (bool, *models.Post, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Get user
	postRequest := new(models.Post)
	has, err := DBengine.Table("post").Where("id = ?", id).Desc("id").Get(postRequest)
	if err != nil {
		return false, postRequest, DBengine, err
	}

	// Post not found
	if !has {
		return false, postRequest, DBengine, nil
	}

	return true, postRequest, DBengine, nil
}

func CheckPassword(given_email string, given_password string) (bool, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	// Search the user in the "user" table
	userTableRequest := new(models.User)
	_, err = DBengine.Table("user").Where("email = ?", given_email).Desc("id").Get(userTableRequest)
	if err != nil {
		return false, err
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(userTableRequest.Password), []byte(given_password)); err != nil {
		return false, nil
	}

	return true, nil
}

func CheckToken(store session.Store, c *fiber.Ctx, email string, token string) (bool, *session.Session, error) {
	// Get store
	sess, err := store.Get(c)
	if err != nil {
		return false, sess, err
	}

	// Check if given token is equal to store token
	check_for_token := sess.Get(email)
	if check_for_token == nil || check_for_token != token {
		return false, sess, nil
	}

	return true, sess, nil
}

func CheckAll(email string, password string, token string, store session.Store, c *fiber.Ctx) (bool, bool, bool, *models.User, *xorm.Engine, *session.Session, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, false, false, nil, DBengine, nil, err
	}

	// Get store
	sess, err := store.Get(c)
	if err != nil {
		return false, false, false, nil, DBengine, sess, err
	}

	// Get user
	userRequest := new(models.User)
	has, err := DBengine.Table("user").Where("email = ?", email).Desc("id").Get(userRequest)
	if err != nil {
		return false, false, false, userRequest, DBengine, sess, err
	}

	// User not found
	if !has {
		return false, false, false, userRequest, DBengine, sess, nil
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(userRequest.Password), []byte(password)); err != nil {
		return true, false, false, userRequest, DBengine, sess, nil
	}


	// Check if given token is equal to store token
	check_for_token := sess.Get(email)
	if check_for_token == nil || check_for_token != token {
		return true, true, false, userRequest, DBengine, sess, nil
	}

	return true, true, true, userRequest, DBengine, sess, nil
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