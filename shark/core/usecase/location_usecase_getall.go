package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type LocationUseCaseGetAll struct {
	repository port.LocationIRepository
}

func NewLocationUseCaseGetAll(repository port.LocationIRepository) *LocationUseCaseGetAll {
	return &LocationUseCaseGetAll{repository: repository}
}

func (o *LocationUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Location {

	entities := o.repository.GetAll(conditions...)
	return entities
}
