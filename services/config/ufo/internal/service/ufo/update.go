package ufo

import (
	"context"

	"github.com/baizhigit/go-ms-examples/config/ufo/internal/model"
)

func (s *service) Update(ctx context.Context, uuid string, updateInfo model.SightingUpdateInfo) error {
	err := s.ufoRepository.Update(ctx, uuid, updateInfo)
	if err != nil {
		return err
	}

	return nil
}
