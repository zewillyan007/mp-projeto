package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type ChipStatusTypeUseCaseGet struct {
	repository port.ChipStatusTypeIRepository
}

func NewChipStatusTypeUseCaseGet(repository port.ChipStatusTypeIRepository) *ChipStatusTypeUseCaseGet {
	return &ChipStatusTypeUseCaseGet{repository: repository}
}

func (o *ChipStatusTypeUseCaseGet) Execute(id int64) (*entity.ChipStatusType, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
