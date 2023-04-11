package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type MicrocontrollerUseCaseGetAll struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseGetAll(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseGetAll {
	return &MicrocontrollerUseCaseGetAll{repository: repository}
}

func (o *MicrocontrollerUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Microcontroller {

	entities := o.repository.GetAll(conditions...)
	return entities
}
