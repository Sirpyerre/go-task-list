package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	slug := chi.URLParam(r, "slug")
	filter := bson.M{
		"slug": slug,
	}

	deleteResult := task.DeleteTask(filter)
	if deleteResult.DeletedCount == 0 {
		fmt.Fprintf(w, "Can not delete slug given")
		return
	}

	fmt.Fprintf(w, "The task has been delete successfully")
}
