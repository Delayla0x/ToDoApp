package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var taskList []Task

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(taskList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, task := range taskList {
		if task.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(task); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func createNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.ID = "Some auto-generated ID" // Replace with actual ID generation
	taskList = append(taskList, newTask)
	json.NewEncoder(w).Encode(newTask)
}

func updateExistingTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, task := range taskList {
		if task.ID == params["id"] {
			taskList = append(taskList[:index], taskList[index+1:]...)
			var updatedTask Task
			if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedTask.ID = params["id"]
			taskList = append(taskList, updatedTask)
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
}

func deleteExistingTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, task := range taskList {
		if task.ID == params["id"] {
			taskList = append(taskList[:index], taskList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(taskList)
}

func main() {
	// DB connection string from .env - Unused in example but might be used for real DB operations
	dbConnectionString := os.Getenv("DB_CONNECTION")
	_ = dbConnectionString

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", getSingleTask).Methods("GET")
	router.HandleFunc("/api/tasks", createNewTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", updateExistingTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", deleteExistingTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}