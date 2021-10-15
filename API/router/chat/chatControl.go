package Chat_router

import (
	"time"

	"faceclone-api/data/models"
	"faceclone-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ChatControlRouter(app fiber.Router, store session.Store) {
	app.Get("/get-chat", get_chat(store))
	app.Post("/send-message", send_message(store))
}

type GetChatRequest struct {
	Sender   string
	Receiver string
}

/* This function fetches a chat */
func get_chat(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(GetChatRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Sender == "" || request.Receiver == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// User is trying to send a message to itself
		if request.Sender == request.Receiver {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "can't send message to self",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists, token is correct, receiver exists and if they have a chat
		hasSender, validToken, hasReceiver, haveChat, _, receiverUserModel, chatModel, DBengine, _, err := utils.FullChatCheck(store, c, request.Sender, token, request.Receiver)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasSender {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid sender user",
			})
		}
		if !validToken {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		if !hasReceiver {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid receiver user",
			})
		}
		if !haveChat {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "chat not found",
			})
		}

		// Search messages of this chat
		var messages []models.ChatMessages
		err = DBengine.Table("chat_messages ").Where("chat_id = ?", chatModel.Id).Desc("time").Find(&messages)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"messages": messages,
			"receiver-name": receiverUserModel.Fullname,
			"receiver-avatar": receiverUserModel.AvatarUrl,
		})
	}
}

type SendMessageRequest struct {
	Content  string
	Sender   string
	Receiver string
}

/* This function sends a message to a chat */
func send_message(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request
		request := new(SendMessageRequest)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		// Not enough values given
		if request.Content == "" || request.Sender == "" || request.Receiver == "" {
			return c.Status(fiber.StatusPartialContent).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}

		// User is trying to send a message to itself
		if request.Sender == request.Receiver {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "can't send message to self",
			})
		}

		// Get token in header
		token := c.Get("access_token")

		// Check if user exists, token is correct, receiver exists and if they have a chat
		hasSender, validToken, hasReceiver, haveChat, senderUserModel, receiverUserModel, chatModel, DBengine, _, err := utils.FullChatCheck(store, c, request.Sender, token, request.Receiver)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}
		if !hasSender {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid sender user",
			})
		}
		if !validToken {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		if !hasReceiver {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid receiver user",
			})
		}

		// If they do not have a chat, create one
		if !haveChat {
			newChat := &models.Chat{
				//Id: ,
				User1: senderUserModel.Id,
				User2: receiverUserModel.Id,
			}
			_, err = DBengine.Insert(newChat)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "database error",
				})
			}
			chatModel = newChat
		}

		// Send message
		newMessage := &models.ChatMessages{
			//Id: ,
			ChatId:  chatModel.Id,
			Time:    time.Now(),
			UserId:  senderUserModel.Id,
			Content: request.Content,
		}
		_, err = DBengine.Insert(newMessage)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "database error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": newMessage,
		})
	}
}
