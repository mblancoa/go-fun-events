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
	App struct {
		TimeToFeed time.Duration `yaml:"time_to_feed"`
	} `yaml:"app"`
}

var configFile string
var CoreContext *coreContext = &coreContext{}

type coreContext struct {
	EventService EventService
}

func SetupCoreConfiguration() {
	log.Info().Msg("Initializing core configuration")
	config := &coreConfiguration{}
	tools.LoadYamlConfiguration(GetConfigFile(), config)
	setupCoreContext(config)
}

func setupCoreContext(conf *coreConfiguration) {
	c := conf.App
	p := ProviderContext
	r := RepositoryContext
	CoreContext.EventService = NewEventService(c.TimeToFeed, r.EventRepository, p.EventProvider)
}

func GetConfigFile() string {
	if configFile == "" {
		mode := os.Getenv(RunMode)
		if mode != "" {
			configFile = fmt.Sprintf("conf/%s.application.yml", mode)
		} else {
			configFile = "conf/application.yml"
		}
	}
	return configFile
}
