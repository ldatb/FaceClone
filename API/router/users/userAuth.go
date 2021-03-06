package User_router

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
	app.Put("/validate", validate())
	app.Post("/login", login(store))
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

type RegisterRequest struct {
	Name     string
	Lastname string
	Email    string
	Password string
}

/* This functions registers and connects an user to the database */
func register(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(RegisterRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Name == "" || request.Lastname == "" || request.Email == "" || request.Password == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Check if user exists
		has, _, DBengine, err := utils.GetUser(request.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if has {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "user already exist",
			})
		}

		// Encrypt the password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Create username
		username := request.Name + request.Lastname
		username = strings.Replace(username, " ", "", -1)
		checkUsername, _, _, err := utils.GetUserByUsername(username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		
		if checkUsername {
			originalUsername := username
			for i := 1; i != 0 ; i++ {
				username = originalUsername + strconv.Itoa(i)

				checkUsername, _, _, err = utils.GetUserByUsername(username)
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

		// Create and insert the new user in the database
		avatarURL, _ := utils.CreateAvatarUrl("base_avatar.png")
		newUser := &models.User{
			//Id:         0,
			Name:       request.Name,
			Lastname:   request.Lastname,
			Fullname:   request.Name + " " + request.Lastname,
			Username:   username,
			Email:      request.Email,
			Password:   string(hashPass),
			AvatarFile: "base_avatar.png",
			AvatarUrl:  avatarURL,
			Validated:  false,
			Followers:  0,
			Following:  0,
		}
		_, err = DBengine.Insert(newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Create Friends table
		newFriends := &models.UserFriends{
			OwnerId:   newUser.Id,
			Followers: make([]string, 0),
			Following: make([]string, 0),
			Friends:   make([]string, 0),
		}
		_, err = DBengine.Insert(newFriends)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Generate and insert an auth token in the database
		token, _ := utils.GenerateAuthKey()
		newToken := &models.UserAuthToken{
			Email:       request.Email,
			AuthToken: token,
			TokenType:   "register",
		}
		_, err = DBengine.Insert(newToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Create a JWT token
		token, expiration, err := createJWT(*newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Store session
		if err := storeSession(c, store, newUser.Email, token); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Update new user
		newUser.AccessToken = token
		_, err = DBengine.ID(newUser.Id).Cols("access_token").Update(newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
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
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "service error",
			})
		}

		// Confirm
		return c.JSON(fiber.Map{
			"token":      token,
			"expiration": expiration,
			"newUser":    newUser,
		})
	}
}

type ValidadeRequest struct {
	Email   string
	AuthKey string
}

/* This function validates the register of an user */
func validate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ValidadeRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.AuthKey == "" {
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

		// Validate
		validation, err := utils.ValidateAuthKey(request.Email, request.AuthKey)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !validation {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid key",
			})
		}

		// Get id
		id := userRequest.Id

		// Change validated from false to true
		userRequest.Validated = true
		_, err = DBengine.ID(id).Cols("validated").Update(userRequest)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userRequest,
		})
	}
}

type LoginRequest struct {
	Email    string
	Password string
}

/* This function connects a user to the database */
func login(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(LoginRequest)

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

		// Check if the password is correct
		checkPass, err := utils.CheckPassword(request.Email, request.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkPass {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid password",
			})
		}

		// Create a JWT token
		token, expiration, err := createJWT(*userRequest)
		if err != nil {
			return err
		}

		// Store session
		if err := storeSession(c, store, userRequest.Email, token); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Store session in user model
		userRequest.AccessToken = token
		_, err = DBengine.ID(userRequest.Id).Cols("access_token").Update(userRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.JSON(fiber.Map{
			"token":      token,
			"expiration": expiration,
			"user":       userRequest,
		})
	}
}
