package models

type Task struct {
	Id     int    `json:"Id"`
	Text   string `json:"Text"`
	IsDone bool   `json:"IsDone"`
}

var TaskStorage []Task
