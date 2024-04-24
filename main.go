package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sirpyerre/taskList/handlers"
	"github.com/go-chi/chi/v5"
)

const portNumber = ":8080"

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("index route")
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	router := chi.NewRouter()

	router.Get("/", indexRoute)
	router.Get("/tasks", handlers.GetTasks)
	router.Post("/tasks", handlers.CreateTask)
	router.Get("/tasks/{slug}", handlers.GetTask)
	router.Delete("/tasks/{slug}", handlers.DeleteTask)
	router.Put("/tasks/{slug}", handlers.UpdateTask)

	fmt.Println(fmt.Sprintf("Starting application in port %q", portNumber))
	log.Fatal(http.ListenAndServe(portNumber, router))
}
