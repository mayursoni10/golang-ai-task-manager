package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mayursoni10/golang-ai-task-manager/internal/models"
	"net/http"
)

func GetTasksHandler(c *gin.Context) {
	userID := c.GetString("user_id")

	// Fetch tasks from database (pseudo-code)
	tasks, err := models.GetTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	task.UserID = userID

	// Save task to database (pseudo-code)
	if err := models.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	taskID := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	task.ID = taskID
	task.UserID = userID

	// Update task in database (pseudo-code)
	if err := models.UpdateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	taskID := c.Param("id")

	// Delete task from database (pseudo-code)
	if err := models.DeleteTask(taskID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
