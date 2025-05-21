package main

import (
	"fmt"
	repo "github.com/mblancoa/go-fun-events/adapters/mongodb-repository"
	prov "github.com/mblancoa/go-fun-events/adapters/xxx-provider"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	prov.SetupProviderConfiguration()
	repo.SetupMongodbRepositoryConfiguration()
	core.SetupCoreConfiguration()

	startSupplyCron()
}

func startSupplyCron() {
	//Todo
	feedInterval := core.DomainContext.CoreConfiguration.Supply.FeedInterval
	ticker := time.NewTicker(feedInterval)
	defer log.Info().Msgf("SupplyService has been stopped")
	defer ticker.Stop()

	service := core.DomainContext.SupplyService
	fmt.Println("Starting supply service cron")
	for {
		select {
		case <-ticker.C:
			service.FetchEventsFromProvider()
		}
	}

}
