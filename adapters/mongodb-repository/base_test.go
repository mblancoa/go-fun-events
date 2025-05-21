package mongodb_repository

import (
	"context"
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBSuite struct {
	suite.Suite
	server           *mim.Server
	client           *mongo.Client
	database         *mongo.Database
	eventsCollection *mongo.Collection
}

func (suite *mongoDBSuite) SetupSuite() {
	testCtx := context.Background()

	server, err := mim.Start(testCtx, "5.0.2")
	tools.ManageTestError(err)
	suite.server = server

	client, err := mongo.Connect(testCtx, options.Client().ApplyURI(server.URI()))
	tools.ManageTestError(err)
	//Use client as needed
	err = client.Ping(testCtx, nil)
	tools.ManageTestError(err)
	suite.client = client
	suite.database = client.Database("FunDatabase")
	suite.Assert()
}

func (suite *mongoDBSuite) TearDownSuite() {
	ctx := context.TODO()
	defer suite.server.Stop(ctx)
	err := suite.client.Disconnect(ctx)
	tools.ManageTestError(err)
}

func (suite *mongoDBSuite) setupEventsCollection() {
	db := suite.database
	log.Debug().Msgf("Creating collection '%s'", EventsCollection)
	err := db.CreateCollection(context.TODO(), EventsCollection)
	tools.ManageTestError(err)

	collection := db.Collection(EventsCollection)

	idIdx := []mongo.IndexModel{
		{
			Keys: bson.M{
				"_id": 1, // index in ascending order
			},
		},
		{
			Keys: bson.M{
				"prov_id": 1, // index in ascending order
			},
		},
		{
			Keys: bson.M{
				"starts_at": 1, // index in ascending order
			},
		},
		{
			Keys: bson.M{
				"ends_at": 1, // index in ascending order
			},
		},
	}
	s, err := collection.Indexes().CreateMany(context.TODO(), idIdx)
	tools.ManageTestError(err)
	for _, str := range s {
		log.Debug().Msg(str)
	}

	suite.eventsCollection = collection
}

func insertOne(coll *mongo.Collection, ctx context.Context, obj interface{}) {
	log.Debug().Msgf("Inserting %v", obj)
	_, err := coll.InsertOne(ctx, obj)
	tools.ManageTestError(err)
}

func insertMany(coll *mongo.Collection, ctx context.Context, list []interface{}) {
	log.Debug().Msgf("Inserting %v", list)
	_, err := coll.InsertMany(ctx, list)
	tools.ManageTestError(err)
}

func findOne(coll *mongo.Collection, ctx context.Context, property string, value, entity interface{}) {
	log.Debug().Msgf("Finding object from collection '%s'", coll.Name())
	err := coll.FindOne(ctx, bson.M{
		property: value,
	}, options.FindOne().SetSort(bson.M{})).Decode(entity)
	tools.ManageTestError(err)
}

func deleteAll(coll *mongo.Collection, ctx context.Context) {
	log.Debug().Msgf("Deleting all documents in collection '%s'", coll.Name())
	_, err := coll.DeleteMany(ctx, bson.D{})
	tools.ManageTestError(err)
}

func count(coll *mongo.Collection, ctx context.Context) int64 {
	c, err := coll.CountDocuments(ctx, bson.D{})
	tools.ManageTestError(err)
	return c
}
