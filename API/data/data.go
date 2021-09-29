package data

import (
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

type User struct {
	Id 		 int64
	Name 	 string
	Email 	 string
	Password string `json:"-"`
}

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

	// Sync the User struct and the database
	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}
	
	return engine, nil
}