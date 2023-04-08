package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipUseCaseGet struct {
	repository port.SharkChipIRepository
}

func NewSharkChipUseCaseGet(repository port.SharkChipIRepository) *SharkChipUseCaseGet {
	return &SharkChipUseCaseGet{repository: repository}
}

func (o *SharkChipUseCaseGet) Execute(id int64) (*entity.SharkChip, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
