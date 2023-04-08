package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkUseCaseRemove struct {
	repository port.SharkIRepository
}

func NewSharkUseCaseRemove(repository port.SharkIRepository) *SharkUseCaseRemove {
	return &SharkUseCaseRemove{repository: repository}
}

func (o *SharkUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *SharkUseCaseRemove {
	return NewSharkUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *SharkUseCaseRemove) Execute(Shark *entity.Shark) error {
	return o.repository.Remove(Shark)
}
