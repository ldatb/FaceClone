package main

import (
	"faceclone-api/data"
	Users_router "faceclone-api/router/users"
	Posts_router "faceclone-api/router/posts"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Try a connection to the database
	_, err := data.CreateDBEngine()
	if err != nil {
		log.Fatal("Database Connection Error: $s", err)
	}
	fmt.Println("Database connection success!")

	// Create fiber route
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to FaceClone!"))
	})

	// Store user session
	store := data.CreateStore()

	// User router
	users_group := app.Group("/users")
	users_group.Static("/avatar", "./media/avatar")
	Users_router.UserAuthRouter(users_group, *store)
	Users_router.UserChangesRouter(users_group, *store)
	Users_router.UserGettersRouter(users_group)
	Users_router.UserFriendsRouter(users_group, *store)

	// Posts router
	posts_group := app.Group("/posts")
	posts_group.Static("/post/media", "./media/post_media")
	Posts_router.PostsControlRouter(posts_group, *store)
	Posts_router.PostsGettersRouter(posts_group)
	Posts_router.CommentsControlRouter(posts_group, *store)
	Posts_router.ReactionControlRouter(posts_group, *store)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Get port to listen
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := ":" + os.Getenv("PORT")

	// Fiber listen
	log.Fatal(app.Listen(PORT))
}
