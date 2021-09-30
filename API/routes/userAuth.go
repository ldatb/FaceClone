package routes

import (
	"time"

	"faceclone-api/data"
	"faceclone-api/data/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func UserAuthRouter(app fiber.Router, store session.Store) {
	app.Post("/register", register(store))
	app.Post("/login", login(store))
	app.Post("/logout", logout(store))
}

/* This functions creates a JSON Web Token to validate the user login */
func createJWT(user entities.User) (string, int64, error) {
	expiration := time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.Id
	claims["email"] = user.Email
	claims["expiration"] = expiration
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiration, nil
}

/* This function stores a JWT session token */
func storeSession(c *fiber.Ctx, store session.Store, email string, token string) (error) {
	// Connect to store
	sess, err := store.Get(c)
    if err != nil {
        return err
    }

	// Save user session
	sess.Set(email, token)

	// Save store
    if err := sess.Save(); err != nil {
        return err
    }

	// Success
	return nil
}

/* This functions registers and connects an user to the database */
type RegisterRequest struct {
	Name string
	Email string
	Password string
}
func register(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(RegisterRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		if (request.Name == "" || request.Email == "" || request.Password == "") {
			return fiber.NewError(fiber.StatusBadRequest, "invalid register credentials")
		}

		// Connect to dabase
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Encrypt the password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Create new user
		newUser := &entities.User{
			Name: request.Name,
			Email: request.Email,
			Password: string(hashPass),
		}

		// Insert the new user in the database
		_, err = DBengine.Insert(newUser)
		if err != nil {
			return err
		}

		// Create a JWT token
		token, expiration, err := createJWT(*newUser)
		if err != nil {
			return err
		}

		// Store session
		if err := storeSession(c, store, newUser.Email, token); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "database error")
		}

		// Confirm
		return c.JSON(fiber.Map{
			"token": token,
			"expiration": expiration,
			"newUser": newUser,
		})
	}
}

/* This function connects a user to the database */
type LoginRequest struct {
	Name string
	Email string
	Password string
}
func login(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(LoginRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		if (request.Email == "" || request.Password == "") {
			return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
		}

		// Connect to dabase
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Search the user
		userRequest := new(entities.User)
		userDb, err := DBengine.Where("email = ?", request.Email).Desc("id").Get(userRequest)
		if err != nil {
			return err
		}

		// User not found
		if !userDb {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}

		// Check if the password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(userRequest.Password), []byte(request.Password)); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "password is incorrect")
		}

		// Create a JWT token
		token, expiration, err := createJWT(*userRequest)
		if err != nil {
			return err
		}

		// Store session
		if err := storeSession(c, store, userRequest.Email, token); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "database error")
		}

		return c.JSON(fiber.Map{
			"token": token,
			"expiration": expiration,
			"user": userRequest,
		})
	}
}

/* This function connects a user to the database */
type LogoutRequest struct {
	Email string
	Token string
}
func logout(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(LogoutRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		if (request.Email == "" || request.Token == "") {
			return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
		}

		// Get store
		sess, err := store.Get(c)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			panic(err)
		}

		// Check if given token is equal to store token
		check_for_token := sess.Get(request.Email)
		if check_for_token == nil || check_for_token != request.Token {
			return fiber.NewError(fiber.StatusBadRequest, "wrong credentials")
		}

		// Delete token and save
		sess.Delete(request.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}