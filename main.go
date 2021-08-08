package main

import (
	"fmt"
	"os"

	"github.com/ajipandean/todo_api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("godotenv.Load: ERROR: Failed loading environment")
	}

	if err := database.ConnectDB(); err != nil {
		panic("database.ConnectDB: ERROR: Failed to connect to database")
	}
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	app.Listen(addr)
}
