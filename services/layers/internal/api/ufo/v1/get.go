package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/baizhigit/go-ms-examples/layers/internal/converter"
	"github.com/baizhigit/go-ms-examples/layers/internal/model"
	ufoV1 "github.com/baizhigit/go-ms-examples/layers/pkg/proto/ufo/v1"
)

func (a *api) Get(ctx context.Context, req *ufoV1.GetRequest) (*ufoV1.GetResponse, error) {
	sighting, err := a.ufoService.Get(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, model.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "sighting with UUID %s not found", req.GetUuid())
		}
		return nil, err
	}

	return &ufoV1.GetResponse{
		Sighting: converter.SightingToProto(sighting),
	}, nil
}
