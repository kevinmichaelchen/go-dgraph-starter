package service

import (
	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/db"
)

type Service struct {
	config   configuration.Config
	dbClient db.Client
}

func NewService(config configuration.Config, dbClient db.Client) Service {
	return Service{
		config:   config,
		dbClient: dbClient,
	}
}

func (s Service) GetDatabase() db.Client {
	return s.dbClient
}
