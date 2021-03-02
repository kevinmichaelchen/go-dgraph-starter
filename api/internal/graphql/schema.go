package graphql

import (
	"github.com/rs/zerolog/log"

	"github.com/graphql-go/graphql"
)

func (s Server) buildSchema() *graphql.Schema {
	todoType := buildTypeForTodo()

	// Schema
	fields := graphql.Fields{
		"todo":  s.buildFieldForGetTodo(todoType),
		"todos": s.buildFieldForGetTodos(todoType),
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
