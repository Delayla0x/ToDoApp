package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // For unique ID generation
	"os"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks = make(map[string]Task) // Changing from slice to map for efficient ID-based operations

func getTasks(c *gin.Context) {
	c.IndentedJSON(200, tasks)
}

func postTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil || newTask.Description == "" { // Validate task description
		c.IndentedJSON(400, gin.H{"message": "invalid request"})
		return
	}
	newTask.ID = uuid.NewString() // Automatically generate unique ID
	tasks[newTask.ID] = newTask
	c.IndentedJSON(201, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	if task, found := tasks[id]; found {
		c.IndentedJSON(200, task)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not found"})
	}
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil || updatedTask.Description == "" { // Validate task description
		c.IndentedJSON(400, gin.H{"message": "invalid request"})
		return
	}
	updatedTask.ID = id // Ensure updated task retains the original ID
	if _, found := tasks[id]; found {
		tasks[id] = updatedTask
		c.IndentedJSON(200, updatedTask)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not found"})
	}
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	if _, found := tasks[id]; found {
		delete(tasks, id)
		c.IndentedJSON(204, nil)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not found"})
	}
}

func searchTasks(c *gin.Context) {
	query := c.Query("query")
	var foundTasks []Task
	for _, task := range tasks {
		if query != "" && (contains(task.Description, query) || (query == "completed" && task.Completed)) {
			foundTasks = append(foundTasks, task)
		}
	}
	c.IndentedJSON(200, foundTasks)
}

func contains(s, substr string) bool {
	return s != "" && substr != "" && len(s) >= len(substr) && len(substr) != 0
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.GET("/search", searchTasks) // New search endpoint

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}