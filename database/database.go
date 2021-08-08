package database

import (
	"fmt"
	"os"

	"github.com/ajipandean/todo_api/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (err error) {
	ds := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(ds), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("ConnectDB: SUCCESS: Connected to database")

	DB.AutoMigrate(&entities.Todo{})

	return nil
}
