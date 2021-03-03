package models

import "time"

type Todo struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Creator   User      `json:"creator"`
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
