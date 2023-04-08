package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type IncidenceUseCaseSave struct {
	repository port.IncidenceIRepository
}

func NewIncidenceUseCaseSave(repository port.IncidenceIRepository) *IncidenceUseCaseSave {
	return &IncidenceUseCaseSave{repository: repository}
}

func (o *IncidenceUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *IncidenceUseCaseSave {
	return NewIncidenceUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *IncidenceUseCaseSave) Execute(Incidence *entity.Incidence) (*entity.Incidence, error) {

	err := Incidence.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(Incidence)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
