package utils

import (
	"faceclone-api/data"
	"faceclone-api/data/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"xorm.io/xorm"
)

func GetChat(user1 int64, user2 int64) (bool, *models.Chat, *xorm.Engine, error) {
	// Connect to database
	DBengine, err := data.CreateDBEngine()
	if err != nil {
		return false, nil, DBengine, err
	}

	// Check if there's a conversation with user1 as user_1 and user2 as user_2
	chatModel := new(models.Chat)
	hasChat1, err := DBengine.Table("chat").Where("user1 = ?", user1).And("user2 = ?", user2).Get(chatModel)
	if err != nil {
		return false, chatModel, DBengine, err
	}

	// Check if there's a conversation with user1 as user_2
	hasChat2 := hasChat1 // This starts as hasChat1 so it does not send a wrong answer
	if !hasChat1 {
		hasChatInternal, err := DBengine.Table("chat").Where("user1 = ?", user2).And("user2 = ?", user1).Get(chatModel)
		if err != nil {
			return false, chatModel, DBengine, err
		}
		hasChat2 = hasChatInternal
	}

	return hasChat2, chatModel, DBengine, err
}

// hasSender, validToken, hasReceiver, haveChat, senderUserModel, receiverUserModel, chatModel, DBengine, sess, err
func FullChatCheck(store session.Store, c *fiber.Ctx, user1 string, user1token string, user2 string) (bool, bool, bool, bool, *models.User, *models.User, *models.Chat, *xorm.Engine, *session.Session, error) {
	// Check if sender exists and has a valid token
	hasSender, validToken, senderUserModel, _, sess, err := CheckUsernameAndToken(store, c, user1, user1token)
	if err != nil {
		return false, false, false, false, senderUserModel, nil, nil, nil, sess, err
	}

	// Sender does not exists
	if !hasSender {
		return false, false, false, false, senderUserModel, nil, nil, nil, sess, nil
	}

	// Token is not valid
	if !validToken {
		return true, false, false, false, senderUserModel, nil, nil, nil, sess, nil
	}

	// Check if receiver exists
	hasReceiver, receiverUserModel, _, err:= GetUserByUsername(user2)
	if err != nil {
		return true, true, false, false, senderUserModel, receiverUserModel, nil, nil, sess, err
	}

	// Receiver does not exist
	if !hasReceiver {
		return true, true, false, false, senderUserModel, receiverUserModel, nil, nil, sess, nil
	}

	// Check if they have a chat already
	haveChat, chatModel, DBengine, err := GetChat(senderUserModel.Id, receiverUserModel.Id)
	if err != nil {
		return true, true, true, false, senderUserModel, receiverUserModel, chatModel, DBengine, sess, err
	}

	// They do not have a chat
	if !haveChat {
		return true, true, true, false, senderUserModel, receiverUserModel, chatModel, DBengine, sess, nil
	}

	return true, true, true, true, senderUserModel, receiverUserModel, chatModel, DBengine, sess, nil
}