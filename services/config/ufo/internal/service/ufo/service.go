package ufo

import (
	"github.com/baizhigit/go-ms-examples/config/ufo/internal/repository"
	def "github.com/baizhigit/go-ms-examples/config/ufo/internal/service"
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
