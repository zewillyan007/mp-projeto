package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipUseCaseRemove struct {
	repository port.SharkChipIRepository
}

func NewSharkChipUseCaseRemove(repository port.SharkChipIRepository) *SharkChipUseCaseRemove {
	return &SharkChipUseCaseRemove{repository: repository}
}

func (o *SharkChipUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *SharkChipUseCaseRemove {
	return NewSharkChipUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *SharkChipUseCaseRemove) Execute(SharkChip *entity.SharkChip) error {
	return o.repository.Remove(SharkChip)
}
