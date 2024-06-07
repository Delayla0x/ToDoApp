package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks = []Task{
	{ID: "1", Description: "Learn Go", Completed: false},
	{ID: "2", Description: "Read Gin documentation", Completed: false},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(200, tasks)
}

func postTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(201, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(200, a)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "task not found"})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Task

	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	for i, a := range tasks {
		if a.ID == id {
			tasks[i] = updatedTask
			c.IndentedJSON(200, updatedTask)
			return
		}
	}

	c.IndentedJSON(404, gin.H{"message": "task not found"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, a := range tasks {
		if a.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(204, nil)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "task not yet deleted"})
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", post 
Task)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}