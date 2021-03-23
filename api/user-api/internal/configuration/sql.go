package configuration

import (
	"database/sql"
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForDBUser = "db_user"
	flagForDBPass = "db_pass"
	flagForDBHost = "db_host"
	flagForDBName = "db_name"
)

type SQLConfig struct {
	Host string
	User string
	Pass string
	Name string
}

func LoadSQLConfig() SQLConfig {
	c := SQLConfig{
		User: "myorg",
		Pass: "myorg",
		Host: "postgres",
		Name: "myorg",
	}

	flag.String(flagForDBUser, c.User, "DB user")
	flag.String(flagForDBPass, c.Pass, "DB pass")
	flag.String(flagForDBHost, c.Host, "DB host")
	flag.String(flagForDBName, c.Name, "DB name")

	flag.Parse()

	viper.BindPFlag(flagForDBUser, flag.Lookup(flagForDBUser))
	viper.BindPFlag(flagForDBPass, flag.Lookup(flagForDBPass))
	viper.BindPFlag(flagForDBHost, flag.Lookup(flagForDBHost))
	viper.BindPFlag(flagForDBName, flag.Lookup(flagForDBName))

	c.User = viper.GetString(flagForDBUser)
	c.Pass = viper.GetString(flagForDBPass)
	c.Host = viper.GetString(flagForDBHost)
	c.Name = viper.GetString(flagForDBName)

	return c
}

func (c SQLConfig) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		c.User, c.Pass, c.Host, c.Name)
}

func (c SQLConfig) Connect() (*sql.DB, error) {
	connectionString := c.GetConnectionString()

	return sql.Open("postgres", connectionString)
}
