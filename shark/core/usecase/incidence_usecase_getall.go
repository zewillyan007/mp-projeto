package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type IncidenceUseCaseGetAll struct {
	repository port.IncidenceIRepository
}

func NewIncidenceUseCaseGetAll(repository port.IncidenceIRepository) *IncidenceUseCaseGetAll {
	return &IncidenceUseCaseGetAll{repository: repository}
}

func (o *IncidenceUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Incidence {

	entities := o.repository.GetAll(conditions...)
	return entities
}
