package repository

import (
	"context"

	"github.com/baizhigit/go-ms-examples/di/ufo/internal/model"
)

type UFORepository interface {
	Create(ctx context.Context, info model.SightingInfo) (string, error)
	Get(ctx context.Context, uuid string) (model.Sighting, error)
	Update(ctx context.Context, uuid string, updateInfo model.SightingUpdateInfo) error
	Delete(ctx context.Context, uuid string) error
}
