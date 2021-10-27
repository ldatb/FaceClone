package User_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"faceclone-api/utils"
)

func UserGettersRouter(app fiber.Router, store session.Store) {
	app.Get("/user", get_user(store))
	app.Get("/get-followers", get_followers())
	app.Get("/get-following", get_following())
	app.Get("/get-friends", get_friends())
}

type UserSearchRequest struct {
	Keyword string
}

/* This function gets an user in the database */
func get_user(store session.Store) fiber.Handler {
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

		has, userModel, _, err := utils.GetUserByUsername(keywordRequest.Keyword)
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
}

func get_followers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get email
		request := new(UserSearchRequest)
		if err := c.QueryParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Keyword == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get user and friends
		hasUser, userModel, userFriendsModel, err := utils.GetFriendsList(request.Keyword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": error(err),
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid user",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user":               userModel,
			"user-follower-list": userFriendsModel.Followers,
		})
	}
}

func get_following() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get email
		request := new(UserSearchRequest)
		if err := c.QueryParser(request); err != nil {
			return err
		}

		// No keywords given
		if request.Keyword == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "no keywords given",
			})
		}

		// Get user and friends
		hasUser, userModel, userFriendsModel, err := utils.GetFriendsList(request.Keyword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid user",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user":                userModel,
			"user-following-list": userFriendsModel.Following,
		})
	}
}

func get_friends() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get email
		request := new(UserSearchRequest)
		if err := c.QueryParser(request); err != nil {
			return err
		}

		// No keywords given
		if request.Keyword == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "no keywords given",
			})
		}

		// Get user and friends
		hasUser, userModel, userFriendsModel, err := utils.GetFriendsList(request.Keyword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUser {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid user",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user":              userModel,
			"user-friends-list": userFriendsModel.Friends,
		})
	}
}