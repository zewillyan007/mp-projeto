package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkUseCaseGet struct {
	repository port.SharkIRepository
}

func NewSharkUseCaseGet(repository port.SharkIRepository) *SharkUseCaseGet {
	return &SharkUseCaseGet{repository: repository}
}

func (o *SharkUseCaseGet) Execute(id int64) (*entity.Shark, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
