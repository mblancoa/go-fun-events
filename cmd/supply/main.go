package main

import (
	repo "github.com/mblancoa/go-fun-events/adapters/mongodb-repository"
	prov "github.com/mblancoa/go-fun-events/adapters/xxx-provider"
	"github.com/mblancoa/go-fun-events/pkg/core"
	"time"
)

func main() {
	prov.SetupProviderConfiguration()
	repo.SetupMongodbRepositoryConfiguration()
	core.SetupCoreConfiguration()

	startSupplyCron()
}

func startSupplyCron() {
	feedInterval := core.DomainContext.CoreConfiguration.Supply.FeedInterval
	service := core.DomainContext.SupplyService
	for {
		service.FetchEventsFromProvider()
		time.Sleep(feedInterval)
	}
}
