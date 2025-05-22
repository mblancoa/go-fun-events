package core

import (
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func init() {
	err := os.Chdir("./..")
	tools.ManageTestError(err)
	err = os.Setenv(RunMode, "test")
	tools.ManageTestError(err)
}

func setupPortContexts(t *testing.T) {
	RepositoryContext.EventRepository = NewMockEventRepository(t)
	ProviderContext.EventProvider = NewMockEventProvider(t)
}
func cleanPortContexts() {
	var eventRepository EventRepository
	var eventProvider EventProvider
	RepositoryContext.EventRepository = eventRepository
	ProviderContext.EventProvider = eventProvider
}
func TestLoadConfiguration(t *testing.T) {
	var config coreConfiguration
	tools.LoadYamlConfiguration(GetConfigFile(), &config)

	assert.NotEmpty(t, config)
	supply := config.Supply
	assert.NotEmpty(t, supply)
	assert.Equal(t, 5*time.Minute, supply.FeedInterval)
}

func TestSSetupCoreConfiguration(t *testing.T) {
	setupPortContexts(t)
	defer cleanPortContexts()

	SetupCoreConfiguration()

	assert.NotEmpty(t, DomainContext)
	assert.NotEmpty(t, DomainContext.EventService)
	assert.NotEmpty(t, DomainContext.SupplyService)
}
