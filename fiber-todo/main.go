package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gitlab.com/florian-nguyen/training/fiber-todo/config"
	"gitlab.com/florian-nguyen/training/fiber-todo/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// dotenv - initial setup
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// config database
	config.ConnectDB()

	// Setup routes
	setupRoutes(app)

	// Listen port 8000 and catch error if any
	err = app.Listen(":8000")

	// Error handling
	if err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// connect todo routes
	routes.TodoRoute(api.Group("/todos"))
}
