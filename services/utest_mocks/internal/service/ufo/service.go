package ufo

import (
	"github.com/baizhigit/go-ms-examples/utest_mocks/internal/repository"
	def "github.com/baizhigit/go-ms-examples/utest_mocks/internal/service"
)

var _ def.UFOService = (*service)(nil)

type service struct {
	ufoRepository repository.UFORepository
}

func NewService(ufoRepository repository.UFORepository) *service {
	return &service{
		ufoRepository: ufoRepository,
	}
}
