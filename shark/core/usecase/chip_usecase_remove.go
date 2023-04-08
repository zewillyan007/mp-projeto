package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipUseCaseRemove struct {
	repository port.ChipIRepository
}

func NewChipUseCaseRemove(repository port.ChipIRepository) *ChipUseCaseRemove {
	return &ChipUseCaseRemove{repository: repository}
}

func (o *ChipUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *ChipUseCaseRemove {
	return NewChipUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *ChipUseCaseRemove) Execute(Chip *entity.Chip) error {
	return o.repository.Remove(Chip)
}
