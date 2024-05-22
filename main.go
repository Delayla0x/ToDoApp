package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: "1", Name: "Learn Go"},
	{ID: "2", Name: "Build a web server"},
}

func main() {
	r := gin.Default()

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

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tasks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
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
			c.IndentedJSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, a := range tasks {
		if a.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}