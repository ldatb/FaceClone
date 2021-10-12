package Posts_router

import (
	"strconv"

	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CommentsControlRouter(app fiber.Router, store session.Store) {
	app.Post("/create-comment", create_comment(store))
	app.Put("/change-comment", change_comment(store))
	app.Delete("/delete-comment", delete_comment(store))
}

type CreateCommentRequest struct {
	PostId   string
	Email string
	Comment  string
}

/* This function creates a comment */
func create_comment(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(CreateCommentRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.PostId == "" || request.Email == "" || request.Comment == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists
		has, userRequest, _, err := utils.CheckUser(request.Email)
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

		// Check if token is correct
		checkToken, _, err := utils.CheckToken(store, c, userRequest.Email, token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
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
		hasPost, _, DBengine, err := utils.GetPost(postId)
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

		// Create post model
		newComment := &models.PostComments{
			PostId:        postId,
			OwnerUsername: userRequest.Username,
			OwnerId:       userRequest.Id,
			Comment:       request.Comment,
		}

		// Put post in database
		_, err = DBengine.Insert(newComment)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"comment": newComment,
		})
	}
}

type ChangeCommentRequest struct {
	CommentId  string
	Email   string
	NewComment string
}

/* This function changes a comment */
func change_comment(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangeCommentRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.CommentId == "" || request.Email == "" || request.NewComment == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists
		has, userRequest, _, err := utils.CheckUser(request.Email)
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

		// Check if token is correct
		checkToken, _, err := utils.CheckToken(store, c, request.Email, token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Transform comment id in int64
		commentId, err := strconv.ParseInt(request.CommentId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post id",
			})
		}

		// Check if comment exists
		hasPost, commentRequest, DBengine, err := utils.GetComment(commentId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasPost {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "comment not found",
			})
		}

		// Check if user id is equal to the owner
		if commentRequest.OwnerId != userRequest.Id {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not the owner",
			})
		}

		// Put post in database
		commentRequest.Comment = request.NewComment
		_, err = DBengine.ID(commentId).Cols("comment").Update(commentRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"comment": commentRequest,
		})
	}
}

type DeleteCommentRequest struct {
	CommentId string
	Email  string
}

/* This function deletes a comment */
func delete_comment(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(DeleteCommentRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.CommentId == "" || request.Email == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists
		has, userRequest, _, err := utils.CheckUser(request.Email)
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

		// Check if token is correct
		checkToken, _, err := utils.CheckToken(store, c, request.Email, token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !checkToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Transform comment id in int64
		commentId, err := strconv.ParseInt(request.CommentId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post id",
			})
		}

		// Check if comment exists
		hasPost, commentRequest, DBengine, err := utils.GetComment(commentId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasPost {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "comment not found",
			})
		}

		// Check if user id is equal to the owner
		if commentRequest.OwnerId != userRequest.Id {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not the owner",
			})
		}

		// Delete post
		_, err = DBengine.ID(commentId).Delete(commentRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
