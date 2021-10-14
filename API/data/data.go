package data

import (
	"fmt"
	"log"
	"os"

	"faceclone-api/data/models"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

/* User Database */
func CreateDBEngine() (*xorm.Engine, error) {
	// Load .env file with secrets
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_User := os.Getenv("DB_USER")
	DB_Pass := os.Getenv("DB_PASSWORD")
	DB_Name := os.Getenv("DB_NAME")

	// Create XORM engine
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, DB_User, DB_Pass, DB_Name)
	engine, err := xorm.NewEngine("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	// Check if there is a connection
	if err := engine.Ping(); err != nil {
		return nil, err
	}

	// Sync all User related stuff
	if err := engine.Sync(new(models.User)); err != nil { // User general
		return nil, err
	}
	if err := engine.Sync(new(models.AuthToken)); err != nil { // User auth token
		return nil, err
	}
	if err := engine.Sync(new(models.UserAvatar)); err != nil { // User avatar
		return nil, err
	}
	if err := engine.Sync(new(models.UserReactedPosts)); err != nil { // User avatar
		return nil, err
	}

	// Sync all Post related stuff
	if err := engine.Sync(new(models.Post)); err != nil { // Posts
		return nil, err
	}
	if err := engine.Sync(new(models.PostMedia)); err != nil { // Post media
		return nil, err
	}
	if err := engine.Sync(new(models.PostComments)); err != nil { // Post comments
		return nil, err
	}
	if err := engine.Sync(new(models.PostReactions)); err != nil { // Post comments
		return nil, err
	}

	return engine, nil
}

/* Session storage stores user session JWT */
func CreateStore() *session.Store {
	// Load .env file with secrets
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_User := os.Getenv("DB_USER")
	DB_Pass := os.Getenv("DB_PASSWORD")
	DB_Name := os.Getenv("DB_NAME")

	// Create storage
	storage := postgres.New(postgres.Config{
		Host:     "localhost",
		Username: DB_User,
		Password: DB_Pass,
		Port:     5432,
		Database: DB_Name,
		Table:    "session_storage",
	})

	store := session.New(session.Config{Storage: storage})

	return store
}
