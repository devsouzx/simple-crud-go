package models

type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

const (
	TableName = "tasks"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE
	)`
)