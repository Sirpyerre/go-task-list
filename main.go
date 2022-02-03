package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/Sirpyerre/taskList/handlers"
	"github.com/gorilla/mux"
)

type allTask []db.Task

var tasks = allTask{
	{
		Name:    "Task 1",
		Content: "Some content",
	},
}


func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to my API")
}

// func updateTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}

// 	var updateTask db.Task
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Please enter valid data")
// 	}

// 	json.Unmarshal(reqBody, &updateTask)

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			updateTask.ID = taskID
// 			tasks = append(tasks, updateTask)

// 			fmt.Fprintf(w, "The task with ID %q has been updated successfully", taskID)
// 		}
// 	}
// }

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{slug}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", hand).Methods("DELETE")
	// router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
