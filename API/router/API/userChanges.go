package API_router

import (
	"fmt"
	"os"

	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

func UserChangesRouter(app fiber.Router, store session.Store) {
	app.Post("/forgot-password", forgot_password())
	app.Post("/change-forgot-password", change_forgot_password(store))
	app.Post("/change-password", change_password(store))
	app.Post("/change-name", change_name(store))
}

type ForgotPasswordRequest struct {
	Email string
}

/* This function sends a token when the user has forgotten it's password */
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

		// Check if user exists
		userExist, userRequest, DBengine, err := utils.CheckUser(request.Email)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "database error")
		}
		if !userExist {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user")
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

type ChangeLostPasswordRequest struct {
	Email    string
	Password string
}

/* This function changes the user password when it was lost */
func change_forgot_password(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangeLostPasswordRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
		}

		// Check if user exists
		userExist, userRequest, DBengine, err := utils.CheckUser(request.Email)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "database error")
		}
		if !userExist {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user")
		}

		// Encrypt the password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Get id
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
		sess.Delete(request.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

type ChangePasswordRequest struct {
	Email        string
	Token        string
	Old_Password string
	New_Password string
}

/* This function changes the user password */
func change_password(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangePasswordRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Token == "" || request.Old_Password == "" || request.New_Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid validation credentials")
		}

		// Check if user exists, password and token are correct
		checkUser, checkPass, checkToken, userRequest, DBengine, _, err := utils.CheckAll(request.Email, request.Old_Password, request.Token, store, c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "database error")
		}
		if !checkUser {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user")
		}
		if !checkPass {
			return fiber.NewError(fiber.StatusBadRequest, "invalid pass")
		}
		if !checkToken {
			return fiber.NewError(fiber.StatusBadRequest, "invalid token")
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

type ChangeNameRequest struct {
	Email        string
	Password     string
	Token        string
	New_Name     string
	New_LastName string
}

/* This function changes a User's name */
func change_name(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangeNameRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Password == "" || request.Token == "" || request.New_Name == "" || request.New_LastName == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
		}

		// Check if user exists, password and token are correct
		checkUser, checkPass, checkToken, userRequest, DBengine, _, err := utils.CheckAll(request.Email, request.Password, request.Token, store, c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "database error")
		}
		if !checkUser {
			return fiber.NewError(fiber.StatusBadRequest, "invalid user")
		}
		if !checkPass {
			return fiber.NewError(fiber.StatusBadRequest, "invalid pass")
		}
		if !checkToken {
			return fiber.NewError(fiber.StatusBadRequest, "invalid token")
		}

		// Change user's name
		id := userRequest.Id
		userRequest.Name = request.New_Name
		userRequest.Lastname = request.New_LastName
		_, err = DBengine.ID(id).Cols("name", "lastname").Update(userRequest)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
