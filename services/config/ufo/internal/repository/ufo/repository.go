package ufo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"

	def "github.com/baizhigit/go-ms-examples/config/ufo/internal/repository"
)

var _ def.UFORepository = (*repository)(nil)

const (
	databaseName   = "ufo_db"
	collectionName = "sightings"
)

type repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client) *repository {
	return &repository{
		client:     client,
		collection: client.Database(databaseName).Collection(collectionName),
	}
}
