package Chat_router

import (
	"faceclone-api/data/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/websocket/v2"
)

/* Websocket configuration */
var clients = make(map[*websocket.Conn]models.User)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)
func runWebsocketHub() {
	for {
		select {
			case connection := <-register:
				clients[connection] = models.User{}

			case message := <-broadcast:
				for connection := range clients {
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						delete(clients, connection)
					}
				}
			case connection := <-unregister:
				delete(clients, connection)
		}
	}
}

func WebsocketRouter(app fiber.Router, store session.Store) {
	app.Get("/ws", websocket_function(store))
}

func websocket_function(store session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}