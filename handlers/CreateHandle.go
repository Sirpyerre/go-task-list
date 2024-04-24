package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/gosimple/slug"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask db.Task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid task")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.Slug = slug.Make(newTask.Name)
	newTask.CreateTask()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
