package main

import (
	"sync"

	"github.com/MyOrg/go-dgraph-starter/internal/app"
	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/rs/zerolog/log"
)

func main() {
	c := configuration.LoadConfig()
	log.Info().Msg("Loaded environment config...")

	fn := obs.InitTracer(c)
	defer fn()

	a := app.NewApp(c)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.Run()

	wg.Wait()
}
