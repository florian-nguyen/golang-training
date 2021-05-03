package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/florian-nguyen/training/fiber-todo/controllers"
)

// Define API routes
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}
