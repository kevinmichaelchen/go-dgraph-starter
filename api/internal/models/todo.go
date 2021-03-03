package models

import "time"

type Todo struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Creator   User      `json:"creator"`
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}
