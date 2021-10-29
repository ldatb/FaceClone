package Posts_router

import (
	"strconv"

	"faceclone-api/data"
	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CommentsControlRouter(app fiber.Router, store session.Store) {
	app.Post("/comments/create", create_comment(store))
	app.Put("/comments/change", change_comment(store))
	app.Delete("/comments/delete", delete_comment(store))
	app.Get("/comments/get/:post_id", get_comments())
}

type CreateCommentRequest struct {
	PostId   string
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
		if request.PostId == "" || request.Comment == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Connect to database
		DBengine, err := data.CreateDBEngine()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Search the user in the "user" table
		userModel := new(models.User)
		hasUser, err := DBengine.Table("user").Where("access_token = ?", token).Desc("id").Get(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// User not found (probably token is not valid)
		if !hasUser {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user or token",
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
			OwnerUsername: userModel.Username,
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
		if request.CommentId == "" || request.NewComment == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "no token",
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
		if commentRequest.OwnerId != userModel.Id {
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
		if request.CommentId == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// If there's no token return
		if token == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "no token",
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
		if commentRequest.OwnerId != userModel.Id {
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

/* This function gets the comments of a post */
func get_comments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get params
		getPostId := c.Params("post_id")

		// Not enough values given
		if getPostId == "" {
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
		hasPost, _, _, err := utils.GetPost(postId)
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

		// Get post comments
		hasComments, postCommentsRequest, _, err := utils.GetPostComments(postId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasComments {
			postCommentsRequest = nil
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"comments": postCommentsRequest,
		})
	}
}
