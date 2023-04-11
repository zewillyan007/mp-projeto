package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type LocationUseCaseSave struct {
	repository port.LocationIRepository
}

func NewLocationUseCaseSave(repository port.LocationIRepository) *LocationUseCaseSave {
	return &LocationUseCaseSave{repository: repository}
}

func (o *LocationUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *LocationUseCaseSave {
	return NewLocationUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *LocationUseCaseSave) Execute(Location *entity.Location) (*entity.Location, error) {

	err := Location.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(Location)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
