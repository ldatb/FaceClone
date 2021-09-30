package router

import (
	"fmt"
	"os"
	"time"

	"faceclone-api/data"
	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

func UserAuthRouter(app fiber.Router, store session.Store) {
	app.Post("/register", register(store))
	app.Post("/login", login(store))
	app.Post("/logout", logout(store))
	app.Post("/validate", validate())
	app.Post("/forgot-password", forgot_password())
	app.Post("/change-forgot-password", change_forgot_password(store))
	app.Post("/change-password", change_password(store))
}

/* This functions creates a JSON Web Token to validate the user login */
func createJWT(user models.User) (string, int64, error) {
	expiration := time.Now().Add(time.Hour).Unix()
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
func storeSession(c *fiber.Ctx, store session.Store, email string, token string) error {
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
	Name     string
	Lastname string
	Email    string
	Password string
}

func register(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(RegisterRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Name == "" || request.Lastname == "" || request.Email == "" || request.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid register credentials")
		}

		// Connect to dabase
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Check if user already exists
		has, err := DBengine.Table("user").Where("email = ?", request.Email).Exist()
		if err != nil {
			return err
		}
		// User already exists
		if has {
			return fiber.NewError(fiber.StatusBadRequest, "user already exists")
		}

		// Encrypt the password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Create and insert the new user in the database
		newUser := &models.User{
			Name:      request.Name,
			Lastname:  request.Lastname,
			Email:     request.Email,
			Password:  string(hashPass),
			Validated: false,
		}
		_, err = DBengine.Insert(newUser)
		if err != nil {
			return err
		}

		// Generate and insert the auth token in the database
		token, _ := utils.GenerateAuthKey()
		newToken := &models.AuthAccess{
			Email:       request.Email,
			AccessToken: token,
			TokenType:   "register",
		}
		_, err = DBengine.Insert(newToken)
		if err != nil {
			return err
		}

		// Get variables for email
		err = godotenv.Load()
		if err != nil {
			return err
		}
		email_host := os.Getenv("EMAIL_HOST")
		email_username := os.Getenv("EMAIL_USERNAME")
		email_password := os.Getenv("EMAIL_PASSWORD")

		// Send confirmation email with the token
		m := gomail.NewMessage()
		m.SetHeader("From", "faceclone@api.com")
		m.SetHeader("To", request.Email)
		m.SetHeader("Subject", "Hello! Please confirm your FaceClone account")
		m.SetBody("text/html", fmt.Sprintf(`
		<div style="height: max-content; font-family: 'Helvetica'; font-size: 16px; word-spacing: 1px;">
			<div style="width: fit-content; background-color: #fff; justify-content: center; max-width: 600px; height: max-content; margin: auto; padding-top: 3rem;">
				<img src="https://i.ibb.co/Qvs67Gd/facebook.png" style="width: 150px; margin-bottom: 3rem;"/>
				<p style="font-size: 22px; font-weight: bold;">Welcome to FaceClone!</p>
				<p>Use the code below to verify your account</p>
				<p style="margin: 1rem 0; display: block; color: #3b5998 !important;"><b>%s</b></p>
			</div>
		</div>
		`, token))
		d := gomail.NewDialer(email_host, 2525, email_username, email_password)
		if err := d.DialAndSend(m); err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "token error")
		}

		// Create a JWT token
		token, expiration, err := createJWT(*newUser)
		if err != nil {
			return err
		}

		// Store session
		if err := storeSession(c, store, newUser.Email, token); err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "database error")
		}

		// Confirm
		return c.JSON(fiber.Map{
			"token":      token,
			"expiration": expiration,
			"newUser":    newUser,
		})
	}
}

/* This function validates the register of an user */
type ValidadeRequest struct {
	Email string
	Token string
}

func validate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ValidadeRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Token == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid validation credentials")
		}

		// Validate
		validation, err := utils.ValidateAuthKey(request.Email, request.Token)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "database error")
		}
		if !validation {
			return fiber.NewError(fiber.StatusBadRequest, "invalid token")
		}

		// Validate account
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Get id
		userRequest := new(models.User)
		_, err = DBengine.Table("user").Where("email = ?", request.Email).Desc("id").Get(userRequest)
		if err != nil {
			return err
		}
		id := userRequest.Id

		// Change validated from false to true
		userRequest.Validated = true
		_, err = DBengine.ID(id).Cols("validated").Update(userRequest)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

/* This function connects a user to the database */
type LoginRequest struct {
	Name     string
	Email    string
	Password string
}

func login(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(LoginRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
		}

		// Connect to dabase
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Search the user
		userRequest := new(models.User)
		userDb, err := DBengine.Table("user").Where("email = ?", request.Email).Desc("id").Get(userRequest)
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
			return fiber.NewError(fiber.StatusServiceUnavailable, "database error")
		}

		return c.JSON(fiber.Map{
			"token":      token,
			"expiration": expiration,
			"user":       userRequest,
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

		// Not enough values given
		if request.Email == "" || request.Token == "" {
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

/* This function sends a token when the user has forgotten it's password */
type ForgotPasswordRequest struct {
	Email string
}

func forgot_password() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(LoginRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
		}

		// Connect to dabase
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Search the user
		userRequest := new(models.User)
		userDb, err := DBengine.Table("user").Where("email = ?", request.Email).Desc("id").Get(userRequest)
		if err != nil {
			return err
		}

		// User not found
		if !userDb {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}

		// Generate and insert the auth token in the database
		token, _ := utils.GenerateAuthKey()
		newToken := &models.AuthAccess{
			Email:       request.Email,
			AccessToken: token,
			TokenType:   "forgot password",
		}
		_, err = DBengine.Insert(newToken)
		if err != nil {
			return err
		}

		// Get variables for email
		err = godotenv.Load()
		if err != nil {
			return err
		}
		email_host := os.Getenv("EMAIL_HOST")
		email_username := os.Getenv("EMAIL_USERNAME")
		email_password := os.Getenv("EMAIL_PASSWORD")

		// Send confirmation email with the token
		m := gomail.NewMessage()
		m.SetHeader("From", "faceclone@api.com")
		m.SetHeader("To", request.Email)
		m.SetHeader("Subject", "Forgot Password")
		m.SetBody("text/html", fmt.Sprintf(`
		<div style="height: max-content; font-family: 'Helvetica'; font-size: 16px; word-spacing: 1px;">
			<div style="width: fit-content; background-color: #fff; justify-content: center; max-width: 600px; height: max-content; margin: auto; padding-top: 3rem;">
				<img src="https://i.ibb.co/Qvs67Gd/facebook.png" style="width: 150px; margin-bottom: 3rem;"/>
				<p style="font-size: 22px; font-weight: bold;">This is your token to change your password</p>
				<p>Use the code below to reset your password</p>
				<p style="margin: 1rem 0; display: block; color: #3b5998 !important;"><b>%s</b></p>
			</div>
		</div>
		`, token))
		d := gomail.NewDialer(email_host, 2525, email_username, email_password)
		if err := d.DialAndSend(m); err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "token error")
		}

		return c.JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

/* This function changes the user password when it was lost */
type ChangeLostPasswordRequest struct {
	Email    string
	Token    string
	Password string
}

func change_forgot_password(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangeLostPasswordRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Token == "" || request.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid validation credentials")
		}

		// Validate
		validation, err := utils.ValidateAuthKey(request.Email, request.Token)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "database error")
		}
		if !validation {
			return fiber.NewError(fiber.StatusBadRequest, "invalid token")
		}

		// Encrypt the password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Get id
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}
		userRequest := new(models.User)
		_, err = DBengine.Table("user").Where("email = ?", request.Email).Desc("id").Get(userRequest)
		if err != nil {
			return err
		}
		id := userRequest.Id

		// Change password in the database
		userRequest.Password = string(hashPass)
		_, err = DBengine.ID(id).Cols("password").Update(userRequest)
		if err != nil {
			return err
		}

		// Delete previous access tokens
		sess, err := store.Get(c)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			panic(err)
		}

		// Delete token and save
		sess.Delete(request.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

/* This function changes the user password */
type ChangePasswordRequest struct {
	Email        string
	Old_Password string
	New_Password string
}

func change_password(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangePasswordRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Old_Password == "" || request.New_Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid validation credentials")
		}

		// Connect to database
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			panic(err)
		}

		// Get user
		userRequest := new(models.User)
		userDb, err := DBengine.Table("user").Where("email = ?", request.Email).Desc("id").Get(userRequest)
		if err != nil {
			return err
		}

		// User not found
		if !userDb {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}

		// Check if the password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(userRequest.Password), []byte(request.Old_Password)); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "password is incorrect")
		}

		// Encrypt the new password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.New_Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Change password in the database
		id := userRequest.Id
		userRequest.Password = string(hashPass)
		_, err = DBengine.ID(id).Cols("password").Update(userRequest)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
