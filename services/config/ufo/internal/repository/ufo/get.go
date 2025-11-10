package ufo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/baizhigit/go-ms-examples/config/ufo/internal/model"
	repoConverter "github.com/baizhigit/go-ms-examples/config/ufo/internal/repository/converter"
	repoModel "github.com/baizhigit/go-ms-examples/config/ufo/internal/repository/model"
)

func (r *repository) Get(ctx context.Context, uuid string) (model.Sighting, error) {
	var repoSighting repoModel.Sighting

	err := r.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&repoSighting)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Sighting{}, model.ErrSightingNotFound
		}
		return model.Sighting{}, err
	}

	return repoConverter.SightingToModel(repoSighting), nil
}
