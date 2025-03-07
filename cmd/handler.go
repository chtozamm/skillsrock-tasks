package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/chtozamm/skillsrock-tasks/internal/database"
	"github.com/chtozamm/skillsrock-tasks/internal/tasks"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

func (app *application) handleGetTasks(c *fiber.Ctx) error {
	// Get tasks from database
	tasks, err := app.queries.GetTasks(c.Context())
	if err != nil {
		log.Println("Error getting tasks:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// If no tasks exist, return representation of an empty JSON array
	if len(tasks) == 0 {
		return c.SendString("[]")
	}
	return c.JSON(tasks)
}

func (app *application) handleDeleteTasks(c *fiber.Ctx) error {
	// Get task ID from path
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	// Delete task from database
	err = app.queries.DeleteTask(c.Context(), int32(id))
	if err != nil {
		log.Println("Error deleting task:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (app *application) handleCreateTasks(c *fiber.Ctx) error {
	// Parse request body
	task := &tasks.CreateTaskModel{}
	if err := c.BodyParser(task); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check for required fields
	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing task title")
	}

	// Create task
	createdTask, err := app.queries.CreateTask(c.Context(), database.CreateTaskParams{
		Title:       task.Title,
		Description: pgtype.Text{String: task.Description, Valid: true},
	})
	if err != nil {
		log.Println("Error creating task:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Respond with created task
	return c.JSON(createdTask)
}

func (app *application) handleUpdateTasks(c *fiber.Ctx) error {
	// Get task ID from path
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	// Parse request body
	wantedTask := &tasks.UpdateTaskModel{}
	if err := c.BodyParser(wantedTask); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if all fields are empty
	if wantedTask.Title == "" && wantedTask.Description == "" && wantedTask.Status == "" {
		return c.Status(fiber.StatusBadRequest).SendString("At least one field is required to be present: title, description, status")
	}

	// Get current state of the task
	currentTask, err := app.queries.GetTaskByID(c.Context(), int32(taskID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		log.Println("Error getting task by ID:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Validate task status
	if wantedTask.Status != "" && wantedTask.Status != tasks.StatusNew && wantedTask.Status != tasks.StatusInProgress && wantedTask.Status != tasks.StatusDone {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task status: expected new, in_progress or done.")
	}

	// Set up update task parameters
	updateTaskParams := database.UpdateTaskParams{
		ID:          currentTask.ID,
		Title:       currentTask.Title,
		Description: currentTask.Description,
		Status:      currentTask.Status,
	}
	if wantedTask.Title != "" {
		updateTaskParams.Title = wantedTask.Title
	}
	if wantedTask.Description != "" {
		updateTaskParams.Description = pgtype.Text{String: wantedTask.Description, Valid: true}
	}
	if wantedTask.Status != "" {
		updateTaskParams.Status = pgtype.Text{String: wantedTask.Status, Valid: true}
	}

	// Update task in database
	updatedTask, err := app.queries.UpdateTask(c.Context(), updateTaskParams)

	if err != nil {
		log.Println("Error updating task:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Respond with updated task
	return c.JSON(updatedTask)
}
