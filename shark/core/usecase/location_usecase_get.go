package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type LocationUseCaseGet struct {
	repository port.LocationIRepository
}

func NewLocationUseCaseGet(repository port.LocationIRepository) *LocationUseCaseGet {
	return &LocationUseCaseGet{repository: repository}
}

func (o *LocationUseCaseGet) Execute(id int64) (*entity.Location, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
