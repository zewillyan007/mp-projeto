package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkUseCaseGetAll struct {
	repository port.SharkIRepository
}

func NewSharkUseCaseGetAll(repository port.SharkIRepository) *SharkUseCaseGetAll {
	return &SharkUseCaseGetAll{repository: repository}
}

func (o *SharkUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Shark {

	entities := o.repository.GetAll(conditions...)
	return entities
}
