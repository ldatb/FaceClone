package Posts_router

import (
	"strconv"

	"faceclone-api/utils"
	"faceclone-api/data/models"

	"github.com/gofiber/fiber/v2"
)

func PostsGettersRouter(app fiber.Router) {
	app.Get("/post/:username/:post_id", get_post())
	app.Get("/user-posts/:username", get_user_posts())
}

/* This function gets a specific post */
func get_post() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get params
		getUsername := c.Params("username")
		getPostId := c.Params("post_id")

		// Not enough values given
		if getPostId == "" || getUsername == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Transform post id in int64
		postId, err := strconv.ParseInt(getPostId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post id",
			})
		}

		// Check if post exists
		hasPost, postRequest, _, err := utils.GetPost(postId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasPost {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post not found",
			})
		}

		// Get username id
		hasUser, userRequest, _, err := utils.CheckUserByUsername(getUsername)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// User not found
		if !hasUser {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user",
			})
		}

		// Compare user id and post owner id
		if postRequest.OwnerId != userRequest.Id {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user and post combination",
			})
		}


		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"post": postRequest,
		})
	}
}

/* This function gets the last 10 posts of an user */
func get_user_posts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get params
		username := c.Params("username")

		// Not enough values given
		if username == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Check if user exists
		has, userRequest, DBengine, err := utils.CheckUserByUsername(username)
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

		// Get latest 10 posts
		var posts []models.Post
		err = DBengine.Table("post").Where("owner_id = ?", userRequest.Id).Desc("id").Find(&posts)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"posts": posts,
		})
	}
}