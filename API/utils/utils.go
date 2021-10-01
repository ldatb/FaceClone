package utils

import (
	"crypto/rand"
	"math"
	"math/big"

	"faceclone-api/data"
	"faceclone-api/data/models"

	"xorm.io/xorm"
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
		authRequest := new(models.AuthAccess)
		authDb, err := DBengine.Table("auth_access").Where("email = ?", email).Get(authRequest)
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
			_, err = DBengine.Table("auth_access").Where("email = ?", email).Delete()
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
	userDb, err := DBengine.Table("user").Where("email = ?", email).Desc("id").Get(userRequest)
	if err != nil {
		return false, userRequest, DBengine, err
	}

	// User not found
	if !userDb {
		return false, userRequest, DBengine, err
	}

	return true, userRequest, DBengine, err
}

func CheckPassword(given_email string, given_password string) (bool, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	// Search the user in the "user" table
	userTableRequest := new(models.User)
	userSearch, err := DBengine.Table("user").Where("email = ?", given_email).Desc("id").Get(userTableRequest)
	if err != nil {
		return false, err
	}

	// User not found
	if !userSearch {
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
	userDb, err := DBengine.Table("user").Where("email = ?", email).Desc("id").Get(userRequest)
	if err != nil {
		return false, false, false, userRequest, DBengine, sess, err
	}

	// User not found
	if !userDb {
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