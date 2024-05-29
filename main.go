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

var taskMap = make(map[string]Task)

func main() {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		context.Next()
	})

	loadInitialTasks()

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTaskByID)
	router.DELETE("/tasks/:id", deleteTaskByID)

	router.Run(":" + serverPort)
}

func loadInitialTasks() {
	defaultTasks := []Task{
		{ID: uuid.New().String(), Name: "Learn Go"},
		{ID: uuid.New().String(), Name: "Build a web server"},
	}

	for _, task := range defaultTasks {
		taskMap[task.ID] = task
	}
}

func getAllTasks(context *gin.Context) {
	var taskList []Task
	for _, task := range taskMap {
		taskList = append(taskList, task)
	}
	context.IndentedJSON(http.StatusOK, taskList)
}

func createTask(context *gin.Context) {
	var newTask Task

	if err := context.BindJSON(&newTask); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task format"})
		return
	}

	newTask.ID = uuid.New().String()
	taskMap[newTask.ID] = newTask
	context.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if task, exists := taskMap[taskID]; exists {
		context.IndentedJSON(http.StatusOK, task)
		return
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTaskByID(context *gin.Context) {
	taskID := context.Param("id")
	var updatedTask Task

	if err := context.BindJSON(&updatedTask); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task format"})
		return
	}

	if _, exists := taskMap[taskID]; exists {
		updatedTask.ID = taskID
		taskMap[taskID] = updatedTask
		context.IndentedJSON(http.StatusOK, updatedTask)
		return
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func deleteTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if _, exists := taskMap[taskID]; exists {
		delete(taskMap, taskID)
		context.Status(http.StatusNoContent)
		return
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}