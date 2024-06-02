package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Task struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
}

var taskMap = make(map[int]Task)

var nextTaskID = 1

var mutex sync.Mutex

func createTask(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()

    var newTask Task
    if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newTask.ID = nextTaskID
    nextTaskID++
    taskMap[newTask.ID] = newTask

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(newTask); err != nil {
        log.Printf("Error responding with created task: %v", err)
    }
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()

    tasks := make([]Task, 0, len(taskMap))
    for _, task := range taskMap {
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(tasks); err != nil {
        log.Printf("Error responding with all tasks: %v", err)
    }
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()

    var modifiedTask Task
    if err := json.NewDecoder(r.Body).Decode(&modifiedTask); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if _, exists := taskMap[modifiedTask.ID]; !exists {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    taskMap[modifiedTask.ID] = modifiedTask

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(modifiedTask); err != nil {
        log.Printf("Error responding with the modified task: %v", err)
    }
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
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

    if _, exists := taskMap[id]; !exists {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    delete(taskMap, id)

    w.WriteHeader(http.StatusOK)
}

func configureRoutes() {
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getAllTasks(w, r)
        case http.MethodPost:
            createTask(w, r)
        }
    })

    http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            updateTask(w, r)
        case http.MethodDelete:
            deleteTask(w, r)
        }
    })
}

func main() {
    configureRoutes()
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}