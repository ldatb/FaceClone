package Posts_router

import (
	"fmt"
	"time"
	"strconv"

	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func PostsControlRouter(app fiber.Router, store session.Store) {
	app.Post("/create-post", create_post(store))
	app.Post("/create-post-media", create_post_media(store))
	app.Put("/change-post", change_post(store))
	app.Delete("/delete-post", delete_post(store))
}

type CreatePostRequest struct {
	PostDescription string
}

/* This function creates a post */
func create_post(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(CreatePostRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.PostDescription == "" {
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

		// Create post model
		newPost := &models.Post{
			//Id:          0,
			OwnerId:     userModel.Id,
			Time: time.Now(),
			MediaName:   "",
			MediaUrl:    "",
			Description: request.PostDescription,
			Likes:       0,
			Hearts:      0,
			Laughs:      0,
			Sads:        0,
			Angries:     0,
		}

		// Put post in database
		_, err = DBengine.Insert(newPost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"post": newPost,
		})
	}
}

/* This function creates a post media */
func create_post_media(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "failed getting credentials",
			})
		}

		// Get user email and post id
		post_id_form := form.Value["post_id"]
		if len(post_id_form) <= 0 {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}
		postIdString := post_id_form[0]

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

		// Transform post id in int64
		postId, err := strconv.ParseInt(postIdString, 10, 64)
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

		// Check if user id is equal to the owner
		if postRequest.OwnerId != userModel.Id {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not the owner",
			})
		}

		// Check if this post already has a media (medias can't be changed)
		if postRequest.MediaName != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post already has media",
			})
		}

		// Get media from request
		media, err := c.FormFile("media")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid image",
			})
		}

		// Save file
		filename := strconv.Itoa(int(postRequest.OwnerId)) + "_" + strconv.Itoa(int(postRequest.Id)) + "_" + media.Filename
		c.SaveFile(media, fmt.Sprintf("./media/post_media/%s", filename))
		mediaURL, _ := utils.CreatePostMediaUrl(filename)

		// Update post media id
		postRequest.MediaName = filename
		postRequest.MediaUrl = mediaURL
		_, err = DBengine.Table("post").Where("id = ?", postRequest.Id).Cols("media_name", "media_url").Update(postRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"post_media_url": mediaURL,
		})
	}
}

type ChangePostRequest struct {
	PostId              string
	New_PostDescription string
}

/* This function changes the description of a post */
func change_post(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(ChangePostRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.PostId == "" || request.New_PostDescription == "" {
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

		// Check if user id is equal to the owner
		if postRequest.OwnerId != userModel.Id {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not the owner",
			})
		}

		// Create and push new post
		postRequest.Description = request.New_PostDescription
		_, err = DBengine.Where("post_id = ?", postId).Cols("description").Update(postRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"post": postRequest,
		})
	}
}

type DeletePostRequest struct {
	PostId string
}

/* This function deletes a post */
func delete_post(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(DeletePostRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.PostId == "" {
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

		// Check if user id is equal to the owner
		if postRequest.OwnerId != userModel.Id {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "not the owner",
			})
		}

		// Delete post
		_, err = DBengine.Where("post_id = ?", postId).Delete(postRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}