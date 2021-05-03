package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/florian-nguyen/training/fiber-todo/config"
	"gitlab.com/florian-nguyen/training/fiber-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Controller function to get the todo list
func GetTodos(c *fiber.Ctx) error {

	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	// Query to filter
	query := bson.D{{}}

	cursor, err := todoCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error upon reading todos in GetTodos()",
			"error":   err,
		})
	}

	var todos []models.Todo = make([]models.Todo, 0)

	// Iterate the cursor and decode each item into a todo item
	err = cursor.All(c.Context(), &todos)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error upon decoding cursor data into a todo item in GetTodos()",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

// Controller function to add a todo item
func CreateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	data := new(models.Todo)

	err := c.BodyParser(&data)

	// Error handling
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON in CreateTodo()",
			"error":   err,
		})
	}

	// Initial values applied to the created todo item
	data.ID = nil
	f := false
	data.Completed = &f
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := todoCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert new todo item in CreateTodo()",
			"error":   err,
		})
	}

	// Get the inserted data
	todo := &models.Todo{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	todoCollection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// Controller function to get a single todo item by id
func GetTodo(c *fiber.Ctx) error {

	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	// Get Id parameter
	paramID := c.Params("id")

	// Convert ID value from type string to ObjectId
	// Mongo expects the ID to be of time ObjectId
	// https://docs.mongodb.com/manual/core/data-modeling-introduction/
	id, err := primitive.ObjectIDFromHex(paramID)

	// Error handling
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id in GetTodo()",
		})
	}

	// Find the todo item matching the specified ID and return
	todo := &models.Todo{}
	query := bson.D{{Key: "_id", Value: id}}

	err = todoCollection.FindOne(c.Context(), query).Decode(todo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found in GetTodo()",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// Controller function to update a todo item
func UpdateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	// Find ID parameter
	paramID := c.Params("id")

	// Convert from type string to type int
	id, err := primitive.ObjectIDFromHex(paramID)

	// Error handling
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id in UpdateTodo()",
			"error":   err,
		})
	}

	// var Data Request
	data := new(models.Todo)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON upon data request in UpdateTodo()",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	// Updata data
	var dataToUpdate bson.D

	if data.Title != nil {
		// todo.Title = *data.Title
		dataToUpdate = append(dataToUpdate, bson.E{Key: "title", Value: data.Title})
	}

	if data.Completed != nil {
		// todo.Completed = *data.Completed
		dataToUpdate = append(dataToUpdate, bson.E{Key: "completed", Value: data.Completed})
	}

	dataToUpdate = append(dataToUpdate, bson.E{Key: "updatedAt", Value: time.Now()})

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	// Update the item with matching ID using the data above
	err = todoCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo not found in updateTodo()",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update todo in updateTodo()",
			"error":   err,
		})
	}

	// Get updated data
	todo := &models.Todo{}

	todoCollection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// Controller function to delete a todo item
func DeleteTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	// Getting Id parameter
	paramID := c.Params("id")

	// Convert from type string to type int
	id, err := primitive.ObjectIDFromHex(paramID)

	// Error handling if parameter cannot be parsed
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id in DeleteTodo()",
		})
	}

	// Find and delete the specified todo item
	query := bson.D{{Key: "_id", Value: id}}

	err = todoCollection.FindOneAndDelete(c.Context(), query).Err()

	// Error handling
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not found in DeleteTodo()",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete todo in DeleteTodo()",
			"error":   err,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
