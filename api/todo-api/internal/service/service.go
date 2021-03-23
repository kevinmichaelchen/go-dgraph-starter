package service

import (
	"github.com/MyOrg/todo-api/internal/configuration"
	"github.com/MyOrg/todo-api/internal/db"
	"github.com/MyOrg/todo-api/internal/search"
)

type Service struct {
	config       configuration.Config
	dbClient     db.Client
	searchClient search.Client
}

func NewService(config configuration.Config, dbClient db.Client, searchClient search.Client) Service {
	return Service{
		config:       config,
		dbClient:     dbClient,
		searchClient: searchClient,
	}
}

func (s Service) GetDatabase() db.Client {
	return s.dbClient
}