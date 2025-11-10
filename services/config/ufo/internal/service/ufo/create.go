package ufo

import (
	"context"

	"github.com/baizhigit/go-ms-examples/config/ufo/internal/model"
)

func (s *service) Create(ctx context.Context, info model.SightingInfo) (string, error) {
	uuid, err := s.ufoRepository.Create(ctx, info)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
