package xxx_provider

import (
	"github.com/mblancoa/go-fun-events/core"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"os"
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
	config = updateConfigFromEnvironment(config)

	setupProviderContext(config)
}

func updateConfigFromEnvironment(config *providerConfiguration) *providerConfiguration {
	name := os.Getenv("PROVIDER_NAME")
	if name != "" {
		config.Provider.Name = name
	}
	url := os.Getenv("PROVIDER_URL")
	if url != "" {
		config.Provider.Url = url
	}
	return config
}

func setupProviderContext(conf *providerConfiguration) {
	log.Info().Msg("Creating the xxx provider context")
	c := conf.Provider
	core.ProviderContext.EventProvider = NewEventProvider(c.Name, c.Url, c.Timeout)
}
