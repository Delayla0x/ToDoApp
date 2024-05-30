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
	router := setupRouter()
	loadInitialTasks()

	serverPort := getServerPort()
	router.Run(":" + serverPort)
}

// setupRouter initializes and configures the router
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTaskByID)
	router.DELETE("/tasks/:id", deleteTaskByID)

	return router
}

// getServerPort retrieves the port number from the environment variables,
// defaulting to 8080 if none is found
func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

// loadInitialTasks pre-loads the task map with default tasks
func loadInitialTasks() {
	defaultTasks := []Task{
		{ID: uuid.New().String(), Name: "Learn Go"},
		{ID: uuid.New().String(), Name: "Build a web server"},
	}

	for _, task := range defaultTasks {
		taskMap[task.ID] = task
	}
}

// getAllTasks handles the GET request to retrieve all tasks
func getAllTasks(context *gin.Context) {
	var taskList []Task
	for _, task := range taskMap {
		taskList = append(taskList, task)
	}

	context.IndentedJSON(http.StatusOK, taskList)
}

// createTask handles the POST request to create a new task
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

// getTaskByID handles the GET request to retrieve a task by its ID
func getTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if task, exists := taskMap[taskID]; exists {
		context.IndentedJSON(http.StatusOK, task)
		return
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// updateTaskByID handles the PUT request to update a task by its ID
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

// deleteTaskByID handles the DELETE request to remove a task by its ID
func deleteTaskByID(context *gin.Context) {
	taskID := context.Param("id")

	if _, exists := taskMap[taskID]; exists {
		delete(taskMap, taskID)
		context.Status(http.StatusNoContent)
		return
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}