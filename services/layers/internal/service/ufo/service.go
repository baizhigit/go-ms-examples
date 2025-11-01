package ufo

import (
	"github.com/baizhigit/go-ms-examples/layers/internal/repository"
	def "github.com/baizhigit/go-ms-examples/layers/internal/service"
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
