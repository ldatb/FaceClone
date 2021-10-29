package Private_router

import (
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func PrivateRouter(app fiber.Router, store session.Store) {
	app.Post("/login", private_login())
	app.Get("/user", get_user())
	app.Delete("/logout", logout(store))
}

/* This function loggs in a user by it's JWT token */
func private_login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Check if user exists and token is correct
		hasUser, userModel, _, err := utils.GetUserByToken(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user or token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userModel,
		})
	}
}

/* This function gets a user by it's JWT token */
func get_user() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Check if user exists and token is correct
		hasUser, userModel, _, err := utils.GetUserByToken(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user or token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": userModel,
		})
	}
}



/* This function connects a user to the database */
func logout(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Check if user exists and token is correct
		hasUser, userModel, DBengine, err := utils.GetUserByToken(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user or token",
			})
		}

		// Get store
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Delete token and save
		sess.Delete(userModel.Email)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Delete session in user model
		userModel.AccessToken = ""
		_, err = DBengine.ID(userModel.Id).Cols("access_token").Update(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
