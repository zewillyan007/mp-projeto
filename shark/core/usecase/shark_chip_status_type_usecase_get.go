package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkChipStatusTypeUseCaseGet struct {
	repository port.SharkChipStatusTypeIRepository
}

func NewSharkChipStatusTypeUseCaseGet(repository port.SharkChipStatusTypeIRepository) *SharkChipStatusTypeUseCaseGet {
	return &SharkChipStatusTypeUseCaseGet{repository: repository}
}

func (o *SharkChipStatusTypeUseCaseGet) Execute(id int64) (*entity.SharkChipStatusType, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
