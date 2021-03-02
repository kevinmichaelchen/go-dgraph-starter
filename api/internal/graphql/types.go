package graphql

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
}

func buildTypeForTodo() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Description: "The Todo's ID",
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
				Description: "The Todo's creation time",
			},
			"title": &graphql.Field{
				Type: graphql.String,
				Description: "The Todo's title",
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
				Description: "Whether or not the Todo has been marked as completed",
			},
		},
	})
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}