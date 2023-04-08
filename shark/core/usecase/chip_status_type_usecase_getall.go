package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipStatusTypeUseCaseGetAll struct {
	repository port.ChipStatusTypeIRepository
}

func NewChipStatusTypeUseCaseGetAll(repository port.ChipStatusTypeIRepository) *ChipStatusTypeUseCaseGetAll {
	return &ChipStatusTypeUseCaseGetAll{repository: repository}
}

func (o *ChipStatusTypeUseCaseGetAll) Execute(conditions ...interface{}) []*entity.ChipStatusType {

	entities := o.repository.GetAll(conditions...)
	return entities
}
