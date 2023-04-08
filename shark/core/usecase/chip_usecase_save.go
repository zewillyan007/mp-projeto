package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipUseCaseSave struct {
	repository port.ChipIRepository
}

func NewChipUseCaseSave(repository port.ChipIRepository) *ChipUseCaseSave {
	return &ChipUseCaseSave{repository: repository}
}

func (o *ChipUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *ChipUseCaseSave {
	return NewChipUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *ChipUseCaseSave) Execute(Chip *entity.Chip) (*entity.Chip, error) {

	err := Chip.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(Chip)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
