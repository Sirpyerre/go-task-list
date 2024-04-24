package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks db.Task
	w.Header().Set("Content-Type", "application/json")

	log.Println("GetTasks route ")

	json.NewEncoder(w).Encode(tasks.AllTasks(bson.M{}, 10))
}
