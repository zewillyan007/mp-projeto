package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SexUseCaseGet struct {
	repository port.SexIRepository
}

func NewSexUseCaseGet(repository port.SexIRepository) *SexUseCaseGet {
	return &SexUseCaseGet{repository: repository}
}

func (o *SexUseCaseGet) Execute(id int64) (*entity.Sex, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
