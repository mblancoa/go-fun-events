package xxx_provider

import (
	"github.com/mblanco/go-fun-events/core"
	"github.com/mblanco/go-fun-events/tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func init() {
	err := os.Chdir("./../..")
	tools.ManageTestError(err)
	err = os.Setenv(core.RunMode, "test")
	tools.ManageTestError(err)
}

func TestLoadConfiguration(t *testing.T) {
	var config providerConfiguration
	tools.LoadYamlConfiguration(core.GetConfigFile(), &config)

	assert.NotEmpty(t, config)
	assert.NotEmpty(t, config.Provider)
	pr := config.Provider
	assert.NotEmpty(t, pr.Name)
	assert.Equal(t, "xxx", pr.Name)
	assert.Equal(t, "https://xxx-provider.com/events", pr.Url)
	assert.Equal(t, 2*time.Second, pr.Timeout)
}

func TestSetupProviderConfiguration(t *testing.T) {
	SetupProviderConfiguration()
	assert.NotEmpty(t, core.ProviderContext)
	assert.NotEmpty(t, core.ProviderContext.EventProvider)
}
