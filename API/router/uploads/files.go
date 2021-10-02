package Uploads

import (
	"github.com/gofiber/fiber/v2"

	"faceclone-api/utils"
)

func FilesRouter(app fiber.Router) {
	app.Get("/avatar/:user_id", get_avatar())
}

/* This function fetches and displays an user's avatar */
func get_avatar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user id
		user_id, err := c.ParamsInt("user_id")
		
		// Invalid user id
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid id",
			})	
		}
		
		// Get user avatar
		has, userAvatarRequest, _, err := utils.GetAvatar(int64(user_id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Avatar not found
		if !has {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}
		
		// Get image url
		URL, err := utils.CreateAvatarUrl(userAvatarRequest.FileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Show user avatar
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"url": URL,
		})
	}
}