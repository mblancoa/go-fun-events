package mongodb_repository

import (
	"context"
	"github.com/mblancoa/go-fun-events/errors"
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDbConfiguration struct {
	Mongodb struct {
		Name string `yaml:"name" env:"MONGODB_NAME"`
		Port int    `yaml:"port" env:"MONGODB_NAME"`
		Uri  string `yaml:"uri" env:"MONGODB_URI"`
	} `yaml:"mongodb"`
}

func SetupMongodbRepositoryConfiguration() {
	log.Info().Msg("Initializing mongodb repository configuration")
	var config = &mongoDbConfiguration{}
	tools.LoadYamlConfiguration(core2.GetConfigFile(), config)

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
	core2.RepositoryContext.EventRepository = eventRepository
}
