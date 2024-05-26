package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var tasks = make(map[string]Task)

func main() {
	r := gin.Default()

	// Middleware for logging the type of HTTP request and path
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	// Initialize tasks with environment variable or default values
	initializeTasks()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":" + port)
}

func initializeTasks() {
	// Example: Load tasks from environment or use defaults
	defaultTasks := []Task{
		{ID: uuid.New().String(), Name: "Learn Go"},
		{ID: uuid.New().String(), Name: "Build a web server"},
	}

	for _, task := range defaultTasks {
		tasks[task.ID] = task
	}
}

func getTasks(c *gin.Context) {
	var tasksList []Task
	for _, task := range tasks {
		tasksList = append(tasksList, task)
	}
	c.IndentedJSON(http.StatusOK, tasksList)
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.ID = uuid.New().String() // Generate a unique ID
	tasks[newTask.ID] = newTask
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	if task, exists := tasks[id]; exists {
		c.IndentedJSON(http.StatusOK, task)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Task

	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	if _, exists := tasks[id]; exists {
		// Preserve the task ID and only update task's name
		updatedTask.ID = id
		tasks[id] = updatedTask
		c.IndentedJSON(http.StatusOK, updatedTask)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		c.Status(http.StatusNoContent)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}