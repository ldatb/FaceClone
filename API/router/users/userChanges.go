package User_router

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	app.Put("/change-forgot-password", change_forgot_password(store))
	app.Put("/change-password", change_password(store))
	app.Put("/change-name", change_name(store))
	app.Put("/change-username", change_username(store))
	app.Put("/change-avatar", change_avatar(store))
	app.Delete("/delete-avatar", delete_avatar(store))
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
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Check if user exists
		has, userRequest, DBengine, err := utils.GetUser(request.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !has {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}

		// Generate and insert the auth token in the database
		token, _ := utils.GenerateAuthKey()
		newToken := &models.UserAuthToken{
			Email:       request.Email,
			AuthToken: token,
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
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "service error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
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
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Check if user exists
		has, userRequest, DBengine, err := utils.GetUser(request.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !has {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
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

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

type ChangePasswordRequest struct {
	Email        string
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
		if request.Email == "" || request.Old_Password == "" || request.New_Password == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Check token in header
		token := c.Get("access_token")

		// Check if user exists, password and token are correct
		checkUser, checkPass, checkToken, userRequest, DBengine, _, err := utils.CheckAll(request.Email, request.Old_Password, token, store, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		if !checkPass {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid password",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
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

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

type ChangeNameRequest struct {
	Email        string
	Password     string
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
		if request.Email == "" || request.Password == "" || request.New_Name == "" || request.New_LastName == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists, password and token are correct
		checkUser, checkPass, checkToken, userRequest, DBengine, _, err := utils.CheckAll(request.Email, request.Password, token, store, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		if !checkPass {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid password",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Change user's name
		id := userRequest.Id
		userRequest.Name = request.New_Name
		userRequest.Lastname = request.New_LastName
		userRequest.Fullname = request.New_Name + " " + request.New_LastName
		_, err = DBengine.ID(id).Cols("name", "lastname", "fullname").Update(userRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

type ChangeUsernameRequest struct {
	Email        string
	Password     string
	New_Username string
}

/* This function changes a User's username */
func change_username(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangeUsernameRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.Password == "" || request.New_Username == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Prepare username
		new_username := strings.Replace(request.New_Username, " ", "", -1)
		checkUsername, _, _, err := utils.GetUserByUsername(new_username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		
		if checkUsername {
			originalUsername := new_username
			for i := 1; i != 0 ; i++ {
				new_username = originalUsername + strconv.Itoa(i)

				checkUsername, _, _, err = utils.GetUserByUsername(new_username)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "database error",
					})
				}

				if !checkUsername {
					break
				}
			}
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists, password and token are correct
		checkUser, checkPass, checkToken, userRequest, DBengine, _, err := utils.CheckAll(request.Email, request.Password, token, store, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		if !checkPass {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid password",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Check if there's a user with this username
		hasUsername, _, _, err := utils.GetUserByUsername(new_username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if hasUsername {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "username already in use",
			})
		}

		// Change user's name
		id := userRequest.Id
		userRequest.Username = new_username
		_, err = DBengine.ID(id).Cols("username").Update(userRequest)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

/* This function changes an user avatar */
func change_avatar(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "failed getting credentials",
			})
		}

		// Get user email
		email_form := form.Value["email"]
		if len(email_form) <= 0 {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}
		email := email_form[0]

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists and token is correct
		hasUser, validToken, userModel, DBengine, _, err := utils.CheckUserAndToken(store, c, email, token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		if !validToken {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Get avatar from request
		avatar, err := c.FormFile("avatar")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid image",
			})
		}

		// Save file
		filename := strconv.Itoa(int(userModel.Id)) + "_" + avatar.Filename
		c.SaveFile(avatar, fmt.Sprintf("./media/avatar/%s", filename))
		avatarURL, err := utils.CreateAvatarUrl(filename)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid image",
			})
		}

		// Update user_avatar in database
		userModel.AvatarFile = filename
		userModel.AvatarUrl = avatarURL
		_, err = DBengine.Where("id = ?", userModel.Id).Cols("avatar_file", "avatar_url").Update(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userModel,
		})
	}
}

type DeleteAvatarRequest struct {
	Email string
}

/* This function deletes an user avatar and sets it to the base one */
func delete_avatar(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(DeleteAvatarRequest)
		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists and token is correct
		hasUser, validToken, userModel, DBengine, _, err := utils.CheckUserAndToken(store, c, request.Email, token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		if !validToken {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Get new avatar url
		avatarURL, _ := utils.CreateAvatarUrl("base_avatar.png")

		// Update user_avatar in database
		userModel.AvatarFile = "base_avatar.png"
		userModel.AvatarUrl = avatarURL
		_, err = DBengine.Where("id = ?", userModel.Id).Cols("avatar_file", "avatar_url").Update(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userModel,
		})
	}
}
