package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/Sirpyerre/taskList/handlers"
	"github.com/go-chi/chi/v5"
)

type allTask []db.Task

var tasks = allTask{
	{
		Name:    "Task 1",
		Content: "Some content",
	},
}

const portNumber = ":8080"

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	router := chi.NewRouter()

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{slug}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{slug}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{slug}", handlers.UpdateTask).Methods("PUT")

	fmt.Println(fmt.Sprintf("Starting application in port %q", portNumber))
	log.Fatal(http.ListenAndServe(portNumber, router))
}
