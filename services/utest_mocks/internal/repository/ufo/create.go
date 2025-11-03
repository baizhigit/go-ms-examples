package ufo

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/baizhigit/go-ms-examples/utest_mocks/internal/model"
	repoConverter "github.com/baizhigit/go-ms-examples/utest_mocks/internal/repository/converter"
	repoModel "github.com/baizhigit/go-ms-examples/utest_mocks/internal/repository/model"
)

func (r *repository) Create(_ context.Context, info model.SightingInfo) (string, error) {
	newUUID := uuid.NewString()

	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[newUUID] = repoModel.Sighting{
		Uuid:      newUUID,
		Info:      repoConverter.SightingInfoToRepoModel(info),
		CreatedAt: time.Now(),
	}

	return newUUID, nil
}
