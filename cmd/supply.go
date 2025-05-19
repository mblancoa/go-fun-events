package main

import (
	"fmt"
	mongodb_repository "github.com/mblanco/go-fun-events/adapters/mongodb-repository"
	xxx_provider "github.com/mblanco/go-fun-events/adapters/xxx-provider"
	"github.com/mblanco/go-fun-events/core"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	xxx_provider.SetupProviderConfiguration()
	mongodb_repository.SetupRepositoryConfiguration()
	core.SetupCoreConfiguration()

	startSupplyCron()
}

func startSupplyCron() {
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
