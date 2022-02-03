package db

type Task struct {
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
