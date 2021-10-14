package Posts_router

import (
	"strconv"

	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ReactionControlRouter(app fiber.Router, store session.Store) {
	app.Post("/react-post", react(store))
	app.Delete("/remove-reaction", remove_reaction(store))
}

type AddReactionRequest struct {
	Email    string
	PostId   string
	Reaction string
}

/* This function add or change a reaction in a post */
func react(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(AddReactionRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.PostId == "" || request.Reaction == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists and token is correct
		hasUser, validToken, userModel, _, _, err := utils.CheckUserAndToken(store, c, request.Email, token)
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

		// Transform post id in int64
		postId, err := strconv.ParseInt(request.PostId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post id",
			})
		}

		// Check if post exists
		hasPost, postRequest, DBengine, err := utils.GetPost(postId)
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

		// Check if reaction has a right name
		if !(request.Reaction == "like" || request.Reaction == "heart" || request.Reaction == "laugh" || request.Reaction == "sad" || request.Reaction == "angry") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid reaction",
			})
		}

		// Get post reactions
		postReactions := new(models.PostReactions)
		_, err = DBengine.Table("post_reactions").Where("post_id = ?", postRequest.Id).Get(postReactions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Check if user already has reacted this post
		userReactedPosts := new(models.UserReactedPosts)
		hasReacted, err := DBengine.Table("user_reacted_posts").Where("owner_id = ?", userModel.Id).And("post_id = ?", postRequest.Id).Get(userReactedPosts)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Has reacted, so we have to change the reaction of the user and the post
		if hasReacted {
			// Subtract reaction in post
			switch userReactedPosts.Reaction {
			case "like":
				postReactions.Likes--
			case "heart":
				postReactions.Hearts--
			case "laugh":
				postReactions.Laughs--
			case "sad":
				postReactions.Sads--
			case "angry":
				postReactions.Angries--
			}

			// Add reaction in post
			switch request.Reaction {
			case "like":
				postReactions.Likes++
			case "heart":
				postReactions.Hearts++
			case "laugh":
				postReactions.Laughs++
			case "sad":
				postReactions.Sads++
			case "angry":
				postReactions.Angries++
			}

			// Update on user reacted posts
			userReactedPosts.Reaction = request.Reaction
			_, err = DBengine.Table("user_reacted_posts").Where("owner_id = ?", userReactedPosts.OwnerId).And("post_id = ?", postRequest.Id).Cols("reaction").Update(userReactedPosts)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}

			// Update on post reactions
			_, err = DBengine.Table("post_reactions").Where("post_id = ?", postReactions.PostId).AllCols().Update(postReactions)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}

			// Return
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"user_reaction": userReactedPosts,
				"post_reactions": postReactions,
			})

		} else { // Has not reacted, so we have to create the reaction of the user and the post
			// Add reaction in post
			switch request.Reaction {
			case "like":
				postReactions.Likes++
			case "heart":
				postReactions.Hearts++
			case "laugh":
				postReactions.Laughs++
			case "sad":
				postReactions.Sads++
			case "angry":
				postReactions.Angries++
			}

			// Create user reaction
			newUserReaction := &models.UserReactedPosts{
				OwnerId: userModel.Id,
				PostId: postRequest.Id,
				Reaction: request.Reaction,
			}

			// Post user reaction and update post reactions
			_, err = DBengine.Insert(newUserReaction)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}
			_, err = DBengine.Table("post_reactions").Where("post_id = ?", postReactions.PostId).AllCols().Update(postReactions)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": error(err),
				})
			}

			// Return
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"user_reaction": newUserReaction,
				"post_reactions": postReactions,
			})
		}
	}
}


type RemoveReactionRequest struct {
	Email    string
	PostId   string
}

/* This function deletes a reaction in a post */
func remove_reaction(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(RemoveReactionRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Email == "" || request.PostId == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists and token is correct
		hasUser, validToken, userModel, _, _, err := utils.CheckUserAndToken(store, c, request.Email, token)
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

		// Transform post id in int64
		postId, err := strconv.ParseInt(request.PostId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post id",
			})
		}

		// Check if post exists
		hasPost, postRequest, DBengine, err := utils.GetPost(postId)
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

		// Get post reactions
		postReactions := new(models.PostReactions)
		_, err = DBengine.Table("post_reactions").Where("post_id = ?", postRequest.Id).Get(postReactions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Check if user already reacted this post
		userReactedPosts := new(models.UserReactedPosts)
		hasReacted, err := DBengine.Table("user_reacted_posts").Where("owner_id = ?", userModel.Id).And("post_id = ?", postRequest.Id).Get(userReactedPosts)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasReacted {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "reaction not found",
			})
		}

		// Subtract reaction in post
		switch userReactedPosts.Reaction {
		case "like":
			postReactions.Likes--
		case "heart":
			postReactions.Hearts--
		case "laugh":
			postReactions.Laughs--
		case "sad":
			postReactions.Sads--
		case "angry":
			postReactions.Angries--
		}

		// Update on post reactions
		_, err = DBengine.Table("post_reactions").Where("post_id = ?", postReactions.PostId).AllCols().Update(postReactions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Delete user reacted posts
		_, err = DBengine.Table("user_reacted_posts").Where("owner_id = ?", userModel.Id).And("post_id = ?", postRequest.Id).Delete(userReactedPosts)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
