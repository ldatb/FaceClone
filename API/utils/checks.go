package utils

import (
	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"
)

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

func CheckUserAndToken(store session.Store, c *fiber.Ctx, email string, token string) (bool, bool, *models.User, *xorm.Engine, *session.Session, error) {
	// Check User
	hasUser, userModel, DBengine, err := GetUser(email)
	if err != nil {
		return false, false, userModel, DBengine, nil, err
	}

	// Check token
	validToken, sess, err := CheckToken(store, c, email, token)
	if err != nil {
		return false, false, userModel, DBengine, sess, err
	}

	return hasUser, validToken, userModel, DBengine, sess, nil
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