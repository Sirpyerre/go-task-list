package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	vars := mux.Vars(r)
	slug := vars["slug"]

	filter := bson.M{
		"slug": slug,
	}

	deleteResult := task.DeleteTask(filter)
	if deleteResult.DeletedCount == 0 {
		fmt.Fprintf(w, "Can not delete slug given")
		return
	}

	fmt.Fprintf(w, "The task has been delete successfully")
	return
}
