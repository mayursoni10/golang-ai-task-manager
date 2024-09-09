package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mayrusoni10/golang-ai-task-manager/internal/config"
	"github.com/mayrusoni10/golang-ai-task-manager/internal/handlers"
	"github.com/mayrusoni10/golang-ai-task-manager/internal/middleware"
	"log"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	config.InitDB()

	r := gin.Default()

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/login", handlers.LoginHandler)
		public.POST("/register", handlers.RegisterHandler)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", handlers.GetTasksHandler)
		protected.POST("/tasks", handlers.CreateTaskHandler)
		protected.PUT("/tasks/:id", handlers.UpdateTaskHandler)
		protected.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
	}

	r.Run(":8080")
}
