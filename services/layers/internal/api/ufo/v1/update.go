package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/baizhigit/go-ms-examples/layers/internal/converter"
	"github.com/baizhigit/go-ms-examples/layers/internal/model"
	ufoV1 "github.com/baizhigit/go-ms-examples/layers/pkg/proto/ufo/v1"
)

func (a *api) Update(ctx context.Context, req *ufoV1.UpdateRequest) (*emptypb.Empty, error) {
	if req.UpdateInfo == nil {
		return nil, status.Error(codes.InvalidArgument, "update_info cannot be nil")
	}

	err := a.ufoService.Update(ctx, req.GetUuid(), converter.UpdateInfoToModel(req.GetUpdateInfo()))
	if err != nil {
		if errors.Is(err, model.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "sighting with UUID %s not found", req.GetUuid())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
