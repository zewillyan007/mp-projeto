package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipUseCaseGetAll struct {
	repository port.SharkChipIRepository
}

func NewSharkChipUseCaseGetAll(repository port.SharkChipIRepository) *SharkChipUseCaseGetAll {
	return &SharkChipUseCaseGetAll{repository: repository}
}

func (o *SharkChipUseCaseGetAll) Execute(conditions ...interface{}) []*entity.SharkChip {

	entities := o.repository.GetAll(conditions...)
	return entities
}
