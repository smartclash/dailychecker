package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env file")
	}

	ConnectDatabase()
	err := DB.AutoMigrate()
	if err != nil {
		fmt.Println("Couldn't migrate database")
	}
}

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to connect to db")
	}
}
