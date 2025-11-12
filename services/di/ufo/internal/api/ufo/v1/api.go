package v1

import (
	ufoV1 "github.com/baizhigit/go-ms-examples/di/shared/pkg/proto/ufo/v1"
	"github.com/baizhigit/go-ms-examples/di/ufo/internal/service"
)

type api struct {
	ufoV1.UnimplementedUFOServiceServer

	ufoService service.UFOService
}

func NewAPI(ufoService service.UFOService) *api {
	return &api{
		ufoService: ufoService,
	}
}
