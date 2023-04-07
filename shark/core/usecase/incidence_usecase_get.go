package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type IncidenceUseCaseGet struct {
	repository port.IncidenceIRepository
}

func NewIncidenceUseCaseGet(repository port.IncidenceIRepository) *IncidenceUseCaseGet {
	return &IncidenceUseCaseGet{repository: repository}
}

func (o *IncidenceUseCaseGet) Execute(id int64) (*entity.Incidence, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
