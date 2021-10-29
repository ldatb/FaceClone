package User_router

import (
	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserFriendsRouter(app fiber.Router, store session.Store) {
	app.Post("/follow", follow(store))
	app.Delete("/unfollow", unfollow(store))

}

type FollowRequest struct {
	Username string
}

func follow(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(FollowRequest)
		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Username == "" {
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

		// Check if username to follow exists
		hasUserToFollow, followedUserModel, _, err := utils.GetUserByUsername(request.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUserToFollow {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user to follow",
			})
		}

		// Check if target is not the same as the user
		if userModel.Id == followedUserModel.Id {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "can't follow self",
			})
		}

		// Get user friends
		userFriendsModel := new(models.UserFriends)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Get(userFriendsModel)
		if err != nil {
			return err
		}

		// Check if it's already following
		isFollowingAlready := utils.Find(userFriendsModel.Following, followedUserModel.Username)
		if isFollowingAlready {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "user is already being followed",
			})
		}

		// Update following array
		userFriendsModel.Following = append(userFriendsModel.Following, followedUserModel.Username)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Cols("following").Update(userFriendsModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Update following number
		userModel.Following += 1
		_, err = DBengine.Table("user").Where("id = ?", userModel.Id).Cols("following").Update(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Get followed user friends
		FollowedUserFriendsModel := new(models.UserFriends)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", followedUserModel.Id).Get(FollowedUserFriendsModel)
		if err != nil {
			return err
		}

		// Update followed user array
		FollowedUserFriendsModel.Followers = append(FollowedUserFriendsModel.Followers, userModel.Username)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", followedUserModel.Id).Cols("followers").Update(FollowedUserFriendsModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Update followed user number
		followedUserModel.Followers += 1
		_, err = DBengine.Table("user").Where("id = ?", followedUserModel.Id).Cols("followers").Update(followedUserModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Check if they follow each other, if so, they become friends
		mutuals := utils.Find(FollowedUserFriendsModel.Following, userModel.Username)
		if mutuals {
			// Add each other to the friends list
			userFriendsModel.Friends = append(userFriendsModel.Friends, followedUserModel.Username)
			FollowedUserFriendsModel.Friends = append(FollowedUserFriendsModel.Friends, userModel.Username)

			// Increase friends number
			userModel.Friends++
			followedUserModel.Friends++

			// Append everything
			_, err1 := DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Cols("friends").Update(userFriendsModel)
			_, err2 := DBengine.Table("user_friends").Where("owner_id = ?", followedUserModel.Id).Cols("friends").Update(FollowedUserFriendsModel)
			_, err3 := DBengine.Table("user").Where("id = ?", userModel.Id).Cols("friends").Update(userModel)
			_, err4 := DBengine.Table("user").Where("id = ?", followedUserModel.Id).Cols("friends").Update(followedUserModel)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user-following":             userModel.Following,
			"user-following-usernames":   userFriendsModel.Following,
			"target-followers":           followedUserModel.Followers,
			"target-followers-usernames": FollowedUserFriendsModel.Followers,
			"friends":                    mutuals,
		})
	}
}

func unfollow(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(FollowRequest)
		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Username == "" {
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

		// Check if username to unfollow exists
		hasUserToUnfollow, unfollowedUserModel, _, err := utils.GetUserByUsername(request.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasUserToUnfollow {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user to follow",
			})
		}

		// Get user friends
		userFriendsModel := new(models.UserFriends)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Get(userFriendsModel)
		if err != nil {
			return err
		}

		// Check if it's not following
		isNotFollowing := utils.Find(userFriendsModel.Following, request.Username)
		if !isNotFollowing {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "user is not being followed",
			})
		}

		// Update following array
		userFriendsModel.Following = utils.FindAndDelete(userFriendsModel.Following, request.Username)
		_, err = DBengine.Where("owner_id = ?", userModel.Id).Cols("following").Update(userFriendsModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Update following number
		userModel.Following -= 1
		_, err = DBengine.Where("id = ?", userModel.Id).Cols("following").Update(userModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Get unfollowed user friends
		UnfollowedUserFriendsModel := new(models.UserFriends)
		_, err = DBengine.Table("user_friends").Where("owner_id = ?", unfollowedUserModel.Id).Get(UnfollowedUserFriendsModel)
		if err != nil {
			return err
		}

		// Update followed user array
		UnfollowedUserFriendsModel.Followers = utils.FindAndDelete(UnfollowedUserFriendsModel.Followers, userModel.Username)
		_, err = DBengine.Where("owner_id = ?", unfollowedUserModel.Id).Cols("followers").Update(UnfollowedUserFriendsModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Update followed user number
		unfollowedUserModel.Followers -= 1
		_, err = DBengine.Where("id = ?", unfollowedUserModel.Id).Cols("followers").Update(unfollowedUserModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		// Check if they follow each other, if so, they become friends
		mutuals := utils.Find(UnfollowedUserFriendsModel.Following, userModel.Username)
		if mutuals {
			// Add each other to the friends list
			userFriendsModel.Friends = utils.FindAndDelete(userFriendsModel.Friends, unfollowedUserModel.Username)
			UnfollowedUserFriendsModel.Friends = utils.FindAndDelete(UnfollowedUserFriendsModel.Friends, userModel.Username)

			// Increase friends number
			userModel.Friends--
			unfollowedUserModel.Friends--

			// Append everything
			_, err1 := DBengine.Table("user_friends").Where("owner_id = ?", userModel.Id).Cols("friends").Update(userFriendsModel)
			_, err2 := DBengine.Table("user_friends").Where("owner_id = ?", unfollowedUserModel.Id).Cols("friends").Update(UnfollowedUserFriendsModel)
			_, err3 := DBengine.Table("user").Where("id = ?", userModel.Id).Cols("friends").Update(userModel)
			_, err4 := DBengine.Table("user").Where("id = ?", unfollowedUserModel.Id).Cols("friends").Update(unfollowedUserModel)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user-following":             userModel.Following,
			"user-following-usernames":   userFriendsModel.Following,
			"target-followers":           unfollowedUserModel.Followers,
			"target-followers-usernames": UnfollowedUserFriendsModel.Followers,
			"where-friends":              mutuals,
		})
	}
}