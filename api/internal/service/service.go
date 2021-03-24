package service

import (
	"github.com/MyOrg/todo-api/internal/configuration"
	"github.com/MyOrg/todo-api/internal/db"
	"github.com/MyOrg/todo-api/internal/search"
	"google.golang.org/grpc"
)

type Service struct {
	config       configuration.Config
	dbClient     db.Client
	searchClient search.Client
	usersConn    *grpc.ClientConn
}

func NewService(
	config configuration.Config,
	dbClient db.Client,
	searchClient search.Client,
	usersConn *grpc.ClientConn) Service {
	return Service{
		config:       config,
		dbClient:     dbClient,
		searchClient: searchClient,
		usersConn:    usersConn,
	}
}

func (s Service) GetDatabase() db.Client {
	return s.dbClient
}
