package graphql

import (
	"github.com/rs/zerolog/log"

	"github.com/graphql-go/graphql"
)

func (s Server) buildSchema() *graphql.Schema {
	todoType := buildTypeForTodo()

	queryFields := graphql.Fields{
		"todo":   s.buildFieldForGetTodo(todoType),
		"todos":  s.buildFieldForGetTodos(todoType),
		"search": s.buildFieldForSearchTodos(todoType),
	}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	})

	mutationFields := graphql.Fields{
		"createTodo": s.buildFieldForCreateTodo(todoType),
		"updateTodo": s.buildFieldForUpdateTodo(todoType),
		"deleteTodo": s.buildFieldForDeleteTodo(todoType),
		"nuke":       s.buildFieldForNuke(),
	}

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create GraphQL schema")
	}

	return &schema
}
