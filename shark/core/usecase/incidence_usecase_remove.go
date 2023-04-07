package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type IncidenceUseCaseRemove struct {
	repository port.IncidenceIRepository
}

func NewIncidenceUseCaseRemove(repository port.IncidenceIRepository) *IncidenceUseCaseRemove {
	return &IncidenceUseCaseRemove{repository: repository}
}

func (o *IncidenceUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *IncidenceUseCaseRemove {
	return NewIncidenceUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *IncidenceUseCaseRemove) Execute(Incidence *entity.Incidence) error {
	return o.repository.Remove(Incidence)
}
