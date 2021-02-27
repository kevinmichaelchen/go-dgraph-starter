package db

import (
	"context"
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"

	"github.com/jackc/pgx/v4"
)

func ConnectionString(config configuration.SQLConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		config.User, config.Pass, config.Host, config.Name)
}

func Connect(config configuration.SQLConfig) (*pgx.Conn, error) {
	connectionString := ConnectionString(config)

	return pgx.Connect(context.TODO(), connectionString)
}
