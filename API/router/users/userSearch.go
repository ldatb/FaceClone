package API_router

import (
	"regexp"

	"github.com/gofiber/fiber/v2"

	"faceclone-api/utils"
)

func UserSearchRouter(app fiber.Router) {
	app.Get("/search", search_user())
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
