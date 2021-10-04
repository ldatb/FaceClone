package main

import (
	"faceclone-api/data"
	Users_router "faceclone-api/router/users"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	//jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/* NOT USING YET
/* Private request to when a user is logged in, requires access token to enter
func private(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"path":    "private",
	})
}

func public(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"path":    "public",
	})
}
NOT USING YET */

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
	user_group := app.Group("/users")
	user_group.Static("/avatar", "./media/avatar")
	Users_router.UserAuthRouter(user_group, *store)
	Users_router.UserChangesRouter(user_group, *store)
	Users_router.UserGettersRouter(user_group)

	/* NOT USING YET
	privateAPI := app.Group("/private")
	privateAPI.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	privateAPI.Get("/", private)

	publicApp := app.Group("/public")
	publicApp.Get("/", public)
	NOT USING YET */

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
