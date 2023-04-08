package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipUseCaseSave struct {
	repository port.SharkChipIRepository
}

func NewSharkChipUseCaseSave(repository port.SharkChipIRepository) *SharkChipUseCaseSave {
	return &SharkChipUseCaseSave{repository: repository}
}

func (o *SharkChipUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *SharkChipUseCaseSave {
	return NewSharkChipUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *SharkChipUseCaseSave) Execute(SharkChip *entity.SharkChip) (*entity.SharkChip, error) {

	err := SharkChip.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(SharkChip)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
