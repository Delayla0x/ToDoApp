package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var taskMap = make(map[string]Task)

func main() {
	router := setupRouter()
	loadInitialTasks()

	serverPort := getServerPort()
	log.Printf("Server starting on port %s\n", serverPort)
	router.Run(":" + serverPort)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTaskByID)
	router.DELETE("/tasks/:id", deleteTaskByID)

	return router
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func loadInitialTasks() {
	defaultTasks := []Task{
		{ID: uuid.New().String(), Name: "Learn Go"},
		{ID: uuid.New().String(), Name: "Build a web server"},
	}

	for _, task := range defaultTasks {
		taskMap[task.ID] = task
	}
	log.Println("Initial tasks loaded")
}

func getAllTasks(context *gin.Context) {
	var taskList []Task
	for _, task := range taskMap {
		taskList = append(taskList, task)
	}

	log.Printf("Retrieving all tasks, total count: %d\n", len(taskList))
	context.IndentedJSON(http.StatusOK, taskList)
}

func createTask(context *gin.Context) {
	var newTask Task

	if err := context.BindJSON(&newTask); err != nil {
		log.Println("Error creating task: Invalid task format")
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task format"})
		return
	}

	newTask.ID = uuid.New().String()
	taskMap[newTask.ID] = newTask
	log.Printf("New task created: ID %s, Name %s\n", newTask.ID, newTask.Name)
	context.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if task, exists := taskMap[taskID]; exists {
		log.Printf("Retrieving task by ID: %s\n", taskID)
		context.IndentedJSON(http.StatusOK, task)
		return
	}

	log.Printf("Task not found by ID: %s\n", taskID)
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTaskByID(context *gin.Context) {
	taskID := context.Param("id")
	var updatedTask Task

	if err := context.BindJSON(&updatedTask); err != nil {
		log.Println("Error updating task: Invalid task format")
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task format"})
		return
	}

	if _, exists := taskMap[taskID]; exists {
		updatedTask.ID = taskID
		taskMap[taskID] = updatedTask
		log.Printf("Task updated: ID %s\n", taskID)
		context.IndentedJSON(http.StatusOK, updatedTask)
		return
	}

	log.Printf("Attempted to update task, but not found by ID: %s\n", taskID)
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func deleteTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if _, exists := taskMap[taskID]; exists {
		delete(taskMap, taskID)
		log.Printf("Task deleted: ID %s\n", taskID)
		context.Status(http.StatusNoContent)
		return
	}

	log.Printf("Attempted to delete task, but not found by ID: %s\n", taskID)
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}