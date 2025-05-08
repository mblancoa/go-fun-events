package xxx_provider

import (
	"github.com/mblanco/Go-Acme-events/core"
	"github.com/mblanco/Go-Acme-events/core/ports"
	"github.com/mblanco/Go-Acme-events/tools"
	"github.com/rs/zerolog/log"
	"time"
)

type providerConfiguration struct {
	Provider struct {
		Name    string        `yaml:"name"`
		Url     string        `yaml:"url"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"provider"`
}

func SetupProviderConfiguration() {
	log.Info().Msg("Initializing xxx provider configuration")
	config := &providerConfiguration{}
	tools.LoadYamlConfiguration(core.GetConfigFile(), config)
	setupProviderContext(config)
}

func setupProviderContext(conf *providerConfiguration) {
	c := conf.Provider
	ports.ProviderContext.EventProvider = NewEventProvider(c.Name, c.Url, c.Timeout)
}
