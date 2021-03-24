package configuration

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	flagForUsersHost = "users_host"
	flagForUsersPort = "users_port"
)

type UsersConfig struct {
	Host string
	Port int
}

func LoadUsersConfig() UsersConfig {
	c := UsersConfig{
		Host: "localhost",
		Port: 8086,
	}

	flag.String(flagForUsersHost, c.Host, "Users host")
	flag.Int(flagForUsersPort, c.Port, "Users port")

	flag.Parse()

	viper.BindPFlag(flagForUsersHost, flag.Lookup(flagForUsersHost))
	viper.BindPFlag(flagForUsersPort, flag.Lookup(flagForUsersPort))

	c.Host = viper.GetString(flagForUsersHost)
	c.Port = viper.GetInt(flagForUsersPort)

	return c
}

func (c UsersConfig) Dial() (*grpc.ClientConn, error) {
	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%d", c.Host, c.Port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect via gRPC: %w", err)
	}
	return grpcConn, nil
}
