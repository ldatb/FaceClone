package main

import (
	"faceclone-api/data"
	"faceclone-api/router/API"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	//jwtware "github.com/gofiber/jwt/v3"
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

	// Routes and groups
	api := app.Group("/api")
	API_router.UserAuthRouter(api, *store)
	API_router.UserChangesRouter(api, *store)

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

	// Fiber listen
	log.Fatal(app.Listen(":3000"))
}
