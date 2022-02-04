package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var updateTask db.Task
	reqBody, err := ioutil.ReadAll(r.Body)
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
