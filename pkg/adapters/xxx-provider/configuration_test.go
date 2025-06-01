package xxx_provider

import (
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	tools2 "github.com/mblancoa/go-fun-events/pkg/tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func init() {
	err := os.Chdir("./../../..")
	tools2.ManageTestError(err)
	err = os.Setenv(core2.RunMode, "test")
	tools2.ManageTestError(err)
}

func TestLoadConfiguration(t *testing.T) {
	var config providerConfiguration
	tools2.LoadYamlConfiguration(core2.GetConfigFile(), &config)

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
	assert.NotEmpty(t, core2.ProviderContext)
	assert.NotEmpty(t, core2.ProviderContext.EventProvider)
}
