package v1

import (
	"github.com/baizhigit/go-ms-examples/utest_mocks/internal/service"
	ufoV1 "github.com/baizhigit/go-ms-examples/utest_mocks/pkg/proto/ufo/v1"
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
