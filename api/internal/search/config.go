package search

import (
	"github.com/meilisearch/meilisearch-go"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForMeilisearchURI       = "meilisearch_uri"
	flagForMeilisearchMasterKey = "meilisearch_master_key"
)

type Config struct {
	URI       string
	MasterKey string
}

func LoadConfig() Config {
	c := Config{
		URI:       "http://127.0.0.1:7700",
		MasterKey: "masterKey",
	}

	flag.String(flagForMeilisearchURI, c.URI, "Meilisearch URI")
	flag.String(flagForMeilisearchMasterKey, c.MasterKey, "Meilisearch master key")

	flag.Parse()

	viper.BindPFlag(flagForMeilisearchURI, flag.Lookup(flagForMeilisearchURI))
	viper.BindPFlag(flagForMeilisearchMasterKey, flag.Lookup(flagForMeilisearchMasterKey))

	c.URI = viper.GetString(flagForMeilisearchURI)
	c.MasterKey = viper.GetString(flagForMeilisearchMasterKey)

	return c
}

func (c Config) NewClient() meilisearch.ClientInterface {
	return meilisearch.NewClient(meilisearch.Config{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})
}
