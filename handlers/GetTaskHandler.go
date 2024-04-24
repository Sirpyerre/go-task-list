package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks db.Task
	slug := chi.URLParam(r, "slug")

	filter := bson.M{
		"slug": slug,
	}

	results := tasks.AllTasks(filter, 10)
	json.NewEncoder(w).Encode(results)
}
