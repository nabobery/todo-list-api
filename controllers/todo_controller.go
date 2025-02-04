package controllers

import (
	"net/http"
	"strconv"
	"todo-list-api/models"
	"todo-list-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoController handles endpoints for managing to-do items.
type TodoController struct {
	todoService services.TodoService
}

// NewTodoController creates a new TodoController instance.
func NewTodoController(todoService services.TodoService) *TodoController {
	return &TodoController{todoService}
}

// CreateTodo handles creating a new to-do item.
//
// @Summary Create a new to-do item
// @Description Create a new to-do item for the authenticated user
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo item"
// @Success 200 {object} models.Todo
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /todos [post]
func (tc *TodoController) CreateTodo(c *gin.Context) {
	userIDStr := c.GetString("userID")
	userObjID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.UserID = userObjID
	if err := tc.todoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo handles updating an existing to-do item.
//
// @Summary Update an existing to-do item
// @Description Update a to-do item if the user is authorized (must be the creator)
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.Todo true "Updated Todo item"
// @Success 200 {object} models.Todo
// @Failure 403 {object} map[string]string "Forbidden"
// @Router /todos/{id} [put]
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	userIDStr := c.GetString("userID")
	id := c.Param("id")

	// Fetch the existing todo to verify the creator.
	existing, err := tc.todoService.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if existing.UserID.Hex() != userIDStr {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = existing.ID
	todo.UserID = existing.UserID
	if err := tc.todoService.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles deleting an existing to-do item.
//
// @Summary Delete a to-do item
// @Description Delete a to-do item if the user is authorized (must be the creator)
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204 "No Content"
// @Failure 403 {object} map[string]string "Forbidden"
// @Router /todos/{id} [delete]
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	userIDStr := c.GetString("userID")
	id := c.Param("id")

	existing, err := tc.todoService.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if existing.UserID.Hex() != userIDStr {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}
	if err := tc.todoService.DeleteTodo(id, userIDStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetTodos handles retrieving a paginated list of the authenticated user’s to-do items.
//
// @Summary Get list of to-do items
// @Description Get paginated to-do items for the authenticated user
// @Tags todos
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page limit" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /todos [get]
func (tc *TodoController) GetTodos(c *gin.Context) {
	userIDStr := c.GetString("userID")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.ParseInt(pageStr, 10, 64)
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	todos, total, err := tc.todoService.GetTodos(userIDStr, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}
