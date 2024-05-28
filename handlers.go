package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var taskStore = make(map[int]Task)
var idCounter = 1
var mutex sync.Mutex

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = idCounter
	idCounter++
	taskStore[task.ID] = task
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	tasks := make([]Task, 0, len(taskStore))
	for _, task := range taskStore {
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := taskStore[updatedTask.ID]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	taskStore[updatedTask.ID] = updatedTask
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if _, ok := taskStore[id]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	delete(taskStore, id)
	w.WriteHeader(http.StatusOK)
}

func setupRoutes() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasksHandler(w, r)
		case http.MethodPost:
			createTaskHandler(w, r)
		}
	})

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateTaskHandler(w, r)
		case http.MethodDelete:
			deleteTaskHandler(w, r)
		}
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}