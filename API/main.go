package main

import (
	"faceclone-api/data"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name string
	Email string
	Password string
}

type LoginRequest struct {
	Name string
	Email string
	Password string
}

/* This functions creates a JSON Web Token to validate the user login */
func createJWT(user data.User) (string, int64, error) {
	expiration := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["expiration"] = expiration
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, expiration, nil
}

/* This functions registers and connects an user to the database */
func register(c *fiber.Ctx) error {
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
	newUser := &data.User{
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

	return c.JSON(fiber.Map{
		"token": token,
		"expiration": expiration,
		"newUser": newUser,
	})
}

/* This function connects a user to the database */
func login(c *fiber.Ctx) error {
	// Get request
	request := new(RegisterRequest)

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
	userRequest := new(data.User)
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

	return c.JSON(fiber.Map{
		"token": token,
		"expiration": expiration,
		"user": userRequest,
	})
}

/* Private request to when a user is logged in, requires access token to enter */
func private(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"path": "private",
	})
}

func public(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"path": "public",
	})
}

/* Main */
func main() {
	app := fiber.New()

	app.Post("/register", register)
	app.Post("/login", login)

	privateApp := app.Group("/private")
	privateApp.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	privateApp.Get("/", private)

	publicApp := app.Group("/public")
	publicApp.Get("/", public)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}