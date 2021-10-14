package User_router

import (
	"regexp"

	"github.com/gofiber/fiber/v2"

	"faceclone-api/utils"
)

func UserGettersRouter(app fiber.Router) {
	app.Get("/user", search_user())
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
			has, userModel, _, err := utils.CheckUser(keywordRequest.Keyword)
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

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"user": userModel,
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