package utils

import (
	"crypto/rand"
	"math"
	"math/big"

	"faceclone-api/data"
	"faceclone-api/data/models"
)

func GenerateAuthKey() (string, error) {
	a, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
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
	authRequest := new(models.UserAuthToken)
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