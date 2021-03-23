package main

import (
	"sync"

	"github.com/MyOrg/user-api/internal/app"
	"github.com/MyOrg/user-api/internal/configuration"
	"github.com/MyOrg/user-api/internal/obs"
	"github.com/rs/zerolog/log"
)

func main() {
	c := configuration.LoadConfig()
	log.Info().Msg("Loaded environment config...")

	fn := obs.InitTracer(c.TraceConfig, c.AppName, c.AppID)
	defer fn()

	a := app.NewApp(c)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.Run()

	wg.Wait()
}
