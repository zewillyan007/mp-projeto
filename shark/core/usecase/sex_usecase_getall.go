package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SexUseCaseGetAll struct {
	repository port.SexIRepository
}

func NewSexUseCaseGetAll(repository port.SexIRepository) *SexUseCaseGetAll {
	return &SexUseCaseGetAll{repository: repository}
}

func (o *SexUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Sex {

	entities := o.repository.GetAll(conditions...)
	return entities
}
