package xxx_provider

import (
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"time"
)

type providerConfiguration struct {
	Provider struct {
		Name    string        `yaml:"name" env:"PROVIDER_NAME"`
		Url     string        `yaml:"url" env:"PROVIDER_URL"`
		Timeout time.Duration `yaml:"timeout" env:"PROVIDER_TIMEOUT"`
	} `yaml:"provider"`
}

func SetupProviderConfiguration() {
	log.Info().Msg("Initializing xxx provider configuration")
	config := &providerConfiguration{}
	tools.LoadYamlConfiguration(core2.GetConfigFile(), config)

	setupProviderContext(config)
}

func setupProviderContext(conf *providerConfiguration) {
	log.Info().Msg("Creating the xxx provider context")
	c := conf.Provider
	core2.ProviderContext.EventProvider = NewEventProvider(c.Name, c.Url, c.Timeout)
}
