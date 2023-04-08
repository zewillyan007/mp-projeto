package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipStatusTypeUseCaseGetAll struct {
	repository port.SharkChipStatusTypeIRepository
}

func NewSharkChipStatusTypeUseCaseGetAll(repository port.SharkChipStatusTypeIRepository) *SharkChipStatusTypeUseCaseGetAll {
	return &SharkChipStatusTypeUseCaseGetAll{repository: repository}
}

func (o *SharkChipStatusTypeUseCaseGetAll) Execute(conditions ...interface{}) []*entity.SharkChipStatusType {

	entities := o.repository.GetAll(conditions...)
	return entities
}
