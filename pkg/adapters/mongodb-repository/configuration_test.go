package mongodb_repository

import (
	"context"
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	tools2 "github.com/mblancoa/go-fun-events/pkg/tools"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mongodbServer *mim.Server

func init() {
	err := os.Chdir("./../../..")
	tools2.ManageTestError(err)
	err = os.Setenv(core2.RunMode, "test")
	tools2.ManageTestError(err)
}

func setupDB() {
	server, err := mim.StartWithOptions(context.TODO(), "6.0.23", mim.WithPort(27018))
	tools2.ManageTestError(err)
	mongodbServer = server
}

func TearDownDB() {
	mongodbServer.Stop(context.TODO())
}

func TestLoadConfiguration(t *testing.T) {
	var config mongoDbConfiguration
	tools2.LoadYamlConfiguration(core2.GetConfigFile(), &config)

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

	assert.NotEmpty(t, core2.RepositoryContext)
	assert.NotEmpty(t, core2.RepositoryContext.EventRepository)
}
