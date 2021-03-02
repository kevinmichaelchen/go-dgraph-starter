package graphql

import (
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/rs/zerolog/log"

	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
}

func buildSchema() *graphql.Schema {
	todoType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

	// Schema
	fields := graphql.Fields{
		"todo": &graphql.Field{
			Type: todoType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ctx := p.Context
				logger := obs.ToLogger(ctx)
				logger.Info().Msg("Resolving GraphQL field: todo")
				return Todo{"1", time.Now(), "Title", false}, nil
			},
			Description: "Retrieve a Todo object",
		},
		"todos": &graphql.Field{
			Type: graphql.NewList(todoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ctx := p.Context
				logger := obs.ToLogger(ctx)
				logger.Info().Msg("Resolving GraphQL field: hello")
				return "world", nil
			},
			Description: "hello",
		},
	}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create GraphQL schema")
	}

	return &schema
}
