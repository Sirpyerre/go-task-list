package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	var updateTask db.Task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	json.Unmarshal(reqBody, &updateTask)

	updateResult := updateTask.UpdateTask(bson.M{"slug": slug})
	if updateResult.MatchedCount == 0 {
		fmt.Fprintf(w, "Can not updated the slug given")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "The task with slug %q has been updated successfully", slug)
}
