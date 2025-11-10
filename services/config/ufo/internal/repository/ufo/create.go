package ufo

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/baizhigit/go-ms-examples/config/ufo/internal/model"
	repoConverter "github.com/baizhigit/go-ms-examples/config/ufo/internal/repository/converter"
	repoModel "github.com/baizhigit/go-ms-examples/config/ufo/internal/repository/model"
)

func (r *repository) Create(ctx context.Context, info model.SightingInfo) (string, error) {
	newUUID := uuid.NewString()

	sighting := repoModel.Sighting{
		Uuid:      newUUID,
		Info:      repoConverter.SightingInfoToRepoModel(info),
		CreatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, sighting)
	if err != nil {
		return "", err
	}

	return newUUID, nil
}
