package User_router

import (
	"regexp"

	"github.com/gofiber/fiber/v2"

	"faceclone-api/utils"
)

func UserGettersRouter(app fiber.Router) {
	app.Get("/user", search_user())
	app.Get("/avatar/:user_id", get_avatar())
}


type UserSearchRequest struct {
	Keyword string
}

/* This function searches an user in the database */
func search_user() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get keyword
		keywordRequest := new(UserSearchRequest)
        if err := c.QueryParser(keywordRequest); err != nil {
			return err
        }

		// No keywords given
		if keywordRequest.Keyword == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "no keywords given",
			})
		}

		// Check if query is an email or a name
		match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, keywordRequest.Keyword)
		if err != nil {
			return err
		}

		// Search user by email
		if match {
			has, userRequest, _, err := utils.CheckUser(keywordRequest.Keyword)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}

			// User not found
			if !has {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "user not found",
				})
			}
			
			// Get user avatar
			hasAvatar, userAvatarRequest, _, err := utils.GetAvatar(userRequest.Id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}

			// Avatar not found
			if !hasAvatar {
				return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
					"error": "user not found",
				})
			}
			
			// Get user avatar url
			avatarURL, err := utils.CreateAvatarUrl(userAvatarRequest.FileName)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"user": userRequest,
				"avatar_url": avatarURL,
			})
		}
	
		// Search user by name
		if !match {
			/* STILL WORKING ON HOW TO DO IT */
			return nil
		}

		return nil
	}
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