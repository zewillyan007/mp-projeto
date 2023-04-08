package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipUseCaseGetAll struct {
	repository port.ChipIRepository
}

func NewChipUseCaseGetAll(repository port.ChipIRepository) *ChipUseCaseGetAll {
	return &ChipUseCaseGetAll{repository: repository}
}

func (o *ChipUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Chip {

	entities := o.repository.GetAll(conditions...)
	return entities
}
