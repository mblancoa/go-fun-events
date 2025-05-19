package core

import (
	"fmt"
	"github.com/mblanco/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

const (
	RunMode = "RUN_MODE"
)

type coreConfiguration struct {
	App *struct {
	} `yaml:"app"`
	Supply *struct {
		FeedInterval time.Duration `yaml:"feed-interval"`
	} `yaml:"supply"`
}

var configFile string
var DomainContext *domainContext = &domainContext{}

type domainContext struct {
	EventService      EventService
	SupplyService     SupplyService
	CoreConfiguration *coreConfiguration
}

func SetupCoreConfiguration() {
	log.Info().Msg("Initializing core configuration")
	config := &coreConfiguration{}
	tools.LoadYamlConfiguration(GetConfigFile(), config)
	setupCoreContext(config)
}

func setupCoreContext(conf *coreConfiguration) {
	log.Info().Msg("Creating the domain context")
	DomainContext.CoreConfiguration = conf
	//c := conf.App
	p := ProviderContext
	r := RepositoryContext
	eventService := NewEventService(r.EventRepository)
	DomainContext.EventService = eventService
	DomainContext.SupplyService = NewSupplyService(eventService, p.EventProvider)
}

func GetConfigFile() string {
	if configFile == "" {
		mode := os.Getenv(RunMode)
		if mode != "" {
			configFile = fmt.Sprintf("conf/%s.application.yml", mode)
		} else {
			configFile = "conf/application.yml"
		}
		log.Info().Msgf("configfile: %s", configFile)
	}
	return configFile
}
