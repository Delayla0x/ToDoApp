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
	newTask := bindTaskFromRequest(c)
	if newTask == nil { // Error handling is done inside bindTaskFromRequest
		return
	}

	tasks = append(tasks, *newTask)
	c.IndentedJSON(201, newTask)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, found := findTaskByID(id)
	if found {
		c.IndentedJSON(200, task)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not found"})
	}
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	updatedTask := bindTaskFromRequest(c)
	if updatedTask == nil { // Error handling inside bindTaskFromRequest
		return
	}

	if updateTaskByID(id, updatedTask) {
		c.IndentedJSON(200, updatedTask)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not found"})
	}
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	if deleteTaskByID(id) {
		c.IndentedJSON(204, nil)
	} else {
		c.IndentedJSON(404, gin.H{"message": "task not yet deleted"})
	}
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func bindTaskFromRequest(c *gin.Context) *Task {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request"})
		return nil
	}
	return &newTask
}

func findTaskByID(id string) (*Task, bool) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, true
		}
	}
	return nil, false
}

func updateTaskByID(id string, updatedTask *Task) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = *updatedTask
			return true
		}
	}
	return false
}

func deleteTaskByID(id string) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}