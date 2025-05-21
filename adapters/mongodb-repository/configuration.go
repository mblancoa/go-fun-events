package mongodb_repository

import (
	"context"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/mblancoa/go-fun-events/errors"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDbConfiguration struct {
	Mongodb struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
		Uri  string `yaml:"uri"`
	} `yaml:"mongodb"`
}

func SetupMongodbRepositoryConfiguration() {
	log.Info().Msg("Initializing mongodb repository configuration")
	var config mongoDbConfiguration
	tools.LoadYamlConfiguration(core.GetConfigFile(), &config)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongodb.Uri))
	errors.ManageErrorPanic(err)
	err = client.Ping(ctx, nil)
	errors.ManageErrorPanic(err)

	database := client.Database(config.Mongodb.Name)
	setupRepositoryContext(database)
}

func setupRepositoryContext(database *mongo.Database) {
	log.Info().Msg("Creating the mongodb repository context")

	eventRepository := NewEventRepository(database.Collection(EventsCollection))
	core.RepositoryContext.EventRepository = eventRepository
}
