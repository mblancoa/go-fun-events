package mongodb_repository

import (
	"context"
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mongodbServer *mim.Server

func init() {
	err := os.Chdir("./../..")
	tools.ManageTestError(err)
	err = os.Setenv(core.RunMode, "test")
	tools.ManageTestError(err)
}

func setupDB() {
	server, err := mim.StartWithOptions(context.TODO(), "5.0.2", mim.WithPort(27018))
	tools.ManageTestError(err)
	mongodbServer = server
}

func TearDownDB() {
	mongodbServer.Stop(context.TODO())
}

func TestLoadConfiguration(t *testing.T) {
	var config mongoDbConfiguration
	tools.LoadYamlConfiguration(core.GetConfigFile(), &config)

	assert.NotEmpty(t, config)
	mongodb := config.Mongodb
	assert.NotEmpty(t, mongodb)
	assert.Equal(t, "FunDatabase", mongodb.Name)
	assert.Equal(t, "mongodb://localhost:27018", mongodb.Uri)
}

func TestSetupMongodbRepositoryConfiguration(t *testing.T) {
	setupDB()
	defer TearDownDB()

	SetupMongodbRepositoryConfiguration()

	assert.NotEmpty(t, core.RepositoryContext)
	assert.NotEmpty(t, core.RepositoryContext.EventRepository)
}
