package tools

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

const (
	testFileName          = "TEST_FILE_NAME"
	testMetadataSize      = "TEST_METADATA_SIZE"
	testMetadataKeyNumber = "TEST_METADATA_KEY_NUMBER"
	testMetadataFrequency = "TEST_METADATA_FREQUENCY"
)

type fileConfiguration struct {
	Root struct {
		File struct {
			Name        string `yaml:"name" json:"name" json:"name" env:"TEST_FILE_NAME"`
			Description string `yaml:"description" json:"description"`
		} `yaml:"file" json:"file"`
		Metadata struct {
			Size      int64         `yaml:"size" json:"size" env:"TEST_METADATA_SIZE"`
			KeyNumber int           `yaml:"key-number" json:"key-number" env:"TEST_METADATA_KEY_NUMBER"`
			Frequency time.Duration `yaml:"frequency" json:"frequency" env:"TEST_METADATA_FREQUENCY"`
		} `yaml:"metadata" json:"metadata"`
	} `yaml:"test" json:"test"`
}

func init() {
	err := os.Chdir("./../..")
	ManageTestError(err)
}

func removeEnvironmentVariables() {
	_ = os.Setenv(testFileName, "")
	_ = os.Setenv(testMetadataSize, "")
	_ = os.Setenv(testMetadataKeyNumber, "")
	_ = os.Setenv(testMetadataFrequency, "")
}

func TestLoadEnvironmentConfiguration(t *testing.T) {
	defer removeEnvironmentVariables()
	_ = os.Setenv(testFileName, "app.conf")
	_ = os.Setenv(testMetadataSize, "2048")
	_ = os.Setenv(testMetadataKeyNumber, "10")
	_ = os.Setenv(testMetadataFrequency, "45s")

	var config fileConfiguration
	LoadEnvironmentConfiguration(&config)

	assert.NotEmpty(t, config)
	root := config.Root
	assert.NotEmpty(t, root)
	file := root.File
	assert.NotEmpty(t, file)
	assert.Equal(t, "app.conf", file.Name)
	assert.Empty(t, file.Description)
	metadata := root.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, int64(2048), metadata.Size)
	assert.Equal(t, int(10), metadata.KeyNumber)
	assert.Equal(t, 45*time.Second, metadata.Frequency)
}

func TestLoadJsonConfiguration(t *testing.T) {
	var config fileConfiguration
	LoadYamlConfiguration("testdata/configuration.json", &config)

	assert.NotEmpty(t, config)
	root := config.Root
	assert.NotEmpty(t, root)
	file := root.File
	assert.NotEmpty(t, file)
	assert.Equal(t, "configuration.json", file.Name)
	assert.Equal(t, "Configuration file to test", file.Description)
	metadata := root.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, int64(1024), metadata.Size)
	assert.Equal(t, int(5), metadata.KeyNumber)
	assert.Equal(t, 3*time.Minute, metadata.Frequency)
}
func TestLoadConfigurationFromJsonAndEnv(t *testing.T) {
	defer removeEnvironmentVariables()
	_ = os.Setenv(testMetadataSize, "2048")
	_ = os.Setenv(testMetadataFrequency, "45s")
	var config fileConfiguration
	LoadYamlConfiguration("testdata/configuration.json", &config)

	assert.NotEmpty(t, config)
	root := config.Root
	assert.NotEmpty(t, root)
	file := root.File
	assert.NotEmpty(t, file)
	assert.Equal(t, "configuration.json", file.Name)
	assert.Equal(t, "Configuration file to test", file.Description)
	metadata := root.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, int64(2048), metadata.Size)
	assert.Equal(t, int(5), metadata.KeyNumber)
	assert.Equal(t, 45*time.Second, metadata.Frequency)
}

func TestLoadYamlConfiguration(t *testing.T) {
	var config fileConfiguration
	LoadYamlConfiguration("testdata/configuration.yml", &config)

	assert.NotEmpty(t, config)
	root := config.Root
	assert.NotEmpty(t, root)
	file := root.File
	assert.NotEmpty(t, file)
	assert.Equal(t, "configuration.yml", file.Name)
	assert.Equal(t, "Configuration file to test", file.Description)
	metadata := root.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, int64(1024), metadata.Size)
	assert.Equal(t, int(5), metadata.KeyNumber)
	assert.Equal(t, 3*time.Minute, metadata.Frequency)
}

func TestLoadConfigurationFromYamlAndEnv(t *testing.T) {
	defer removeEnvironmentVariables()
	_ = os.Setenv(testMetadataSize, "2048")
	_ = os.Setenv(testMetadataFrequency, "45s")
	var config fileConfiguration
	LoadYamlConfiguration("testdata/configuration.yml", &config)

	assert.NotEmpty(t, config)
	root := config.Root
	assert.NotEmpty(t, root)
	file := root.File
	assert.NotEmpty(t, file)
	assert.Equal(t, "configuration.yml", file.Name)
	assert.Equal(t, "Configuration file to test", file.Description)
	metadata := root.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, int64(2048), metadata.Size)
	assert.Equal(t, int(5), metadata.KeyNumber)
	assert.Equal(t, 45*time.Second, metadata.Frequency)
}
