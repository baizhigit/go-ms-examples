package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/baizhigit/go-ms-examples/layers/internal/model"
	ufoV1 "github.com/baizhigit/go-ms-examples/layers/pkg/proto/ufo/v1"
)

func (a *api) Delete(ctx context.Context, req *ufoV1.DeleteRequest) (*emptypb.Empty, error) {
	err := a.ufoService.Delete(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, model.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "sighting with UUID %s not found", req.GetUuid())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
