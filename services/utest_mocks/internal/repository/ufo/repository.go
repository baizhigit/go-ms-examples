package ufo

import (
	"sync"

	def "github.com/baizhigit/go-ms-examples/utest_mocks/internal/repository"
	repoModel "github.com/baizhigit/go-ms-examples/utest_mocks/internal/repository/model"
)

var _ def.UFORepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.Sighting
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repoModel.Sighting),
	}
}
