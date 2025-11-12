package ufo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"

	def "github.com/baizhigit/go-ms-examples/di/ufo/internal/repository"
)

var _ def.UFORepository = (*repository)(nil)

const (
	collectionName = "sightings"
)

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *repository {
	repo := &repository{
		collection: db.Collection(collectionName),
	}

	return repo
}
