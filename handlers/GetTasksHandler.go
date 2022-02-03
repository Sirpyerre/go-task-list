package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sirpyerre/taskList/db"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks db.Task
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks.GetTasks(10))
}
