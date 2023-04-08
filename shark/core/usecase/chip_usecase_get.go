package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipUseCaseGet struct {
	repository port.ChipIRepository
}

func NewChipUseCaseGet(repository port.ChipIRepository) *ChipUseCaseGet {
	return &ChipUseCaseGet{repository: repository}
}

func (o *ChipUseCaseGet) Execute(id int64) (*entity.Chip, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
