package graphql

import (
	"time"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
}

func buildTodo(in *todoV1.Todo) (Todo, error) {
	createdAt, err := ptypes.Timestamp(in.CreatedAt)
	if err != nil {
		return Todo{}, err
	}
	return Todo{
		ID:        in.Id,
		CreatedAt: createdAt,
		Title:     in.Title,
		Done:      in.Done,
	}, nil
}

func buildTypeForTodo() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "The Todo's ID",
			},
			"createdAt": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The Todo's creation time",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "The Todo's title",
			},
			"done": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Whether or not the Todo has been marked as completed",
			},
		},
	})
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

func buildTypePageInfo() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "TodosPageInfo",
		Fields: graphql.Fields{
			"endCursor": &graphql.Field{
				Type:        graphql.String,
				Description: "The last cursor in the page",
			},
			"hasNextPage": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Whether there's a next page",
			},
		},
	})
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
