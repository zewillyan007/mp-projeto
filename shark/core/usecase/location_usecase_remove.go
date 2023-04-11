package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type LocationUseCaseRemove struct {
	repository port.LocationIRepository
}

func NewLocationUseCaseRemove(repository port.LocationIRepository) *LocationUseCaseRemove {
	return &LocationUseCaseRemove{repository: repository}
}

func (o *LocationUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *LocationUseCaseRemove {
	return NewLocationUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *LocationUseCaseRemove) Execute(Location *entity.Location) error {
	return o.repository.Remove(Location)
}
