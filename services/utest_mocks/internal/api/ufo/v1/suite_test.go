package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/baizhigit/go-ms-examples/utest_mocks/internal/service/mocks"
)

type APISuite struct {
	suite.Suite

	ctx context.Context

	ufoService *mocks.UFOService

	api *api
}

func (s *APISuite) SetupTest() {
	s.ctx = context.Background()

	s.ufoService = mocks.NewUFOService(s.T())

	s.api = NewAPI(
		s.ufoService,
	)
}

func (s *APISuite) TearDownTest() {
}

func TestAPIIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
